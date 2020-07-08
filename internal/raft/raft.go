package raft

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// Server is a description of a raft server.
type Server interface {
	Start() error
	OnReplication(ReplicationHandler)
	Input(*message.Command)
	io.Closer
}

// ReplicationHandler is a handler setter.
// It takes in the log entries as a string and returns the number
// of succeeded application of entries.
type ReplicationHandler func([]*message.Command) int

// Node describes the current state of a raft node.
// The raft paper describes this as a "State" but node
// seemed more intuitive.
type Node struct {
	State string
	log   zerolog.Logger

	PersistentState     *PersistentState
	VolatileState       *VolatileState
	VolatileStateLeader *VolatileStateLeader
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int32
	VotedFor    id.ID // VotedFor is nil at init, and id.ID of the node after voting is complete.
	Log         []*message.LogData

	SelfID    id.ID
	LeaderID  id.ID          // LeaderID is nil at init, and the ID of the node after the leader is elected.
	PeerIPs   []network.Conn // PeerIPs has the connection variables of all the other nodes in the cluster.
	ConnIDMap map[id.ID]int  // ConnIDMap has a mapping of the ID of the server to its connection.
	mu        sync.Mutex
}

// VolatileState describes the volatile state data on a raft node.
type VolatileState struct {
	CommitIndex int32
	LastApplied int32
}

// VolatileStateLeader describes the volatile state data that exists on a raft leader.
type VolatileStateLeader struct {
	NextIndex  []int // Holds the nextIndex value for each of the followers in the cluster.
	MatchIndex []int // Holds the matchIndex value for each of the followers in the cluster.
}

var _ Server = (*SimpleServer)(nil)

// SimpleServer implements a server in a cluster.
type SimpleServer struct {
	node            *Node
	cluster         Cluster
	onReplication   ReplicationHandler
	log             zerolog.Logger
	timeoutProvider func(*Node) *time.Timer
	lock            sync.Mutex

	onRequestVotes  func(*message.RequestVoteRequest)
	onLeaderElected func()
	onAppendEntries func(*message.AppendEntriesRequest)
}

// incomingData describes every request that the server gets.
type incomingData struct {
	conn network.Conn
	msg  message.Message
}

// NewServer enables starting a raft server/cluster.
func NewServer(log zerolog.Logger, cluster Cluster) *SimpleServer {
	return newServer(log, cluster, nil)
}

func newServer(log zerolog.Logger, cluster Cluster, timeoutProvider func(*Node) *time.Timer) *SimpleServer {
	if timeoutProvider == nil {
		timeoutProvider = randomTimer
	}
	return &SimpleServer{
		log:             log.With().Str("component", "raft").Logger(),
		cluster:         cluster,
		timeoutProvider: timeoutProvider,
	}
}

// NewRaftNode initialises a raft cluster with the given configuration.
func NewRaftNode(cluster Cluster) *Node {
	var nextIndex, matchIndex []int

	for range cluster.Nodes() {
		nextIndex = append(nextIndex, -1)
		matchIndex = append(matchIndex, -1)
	}
	connIDMap := make(map[id.ID]int)
	for i := range cluster.Nodes() {
		connIDMap[cluster.Nodes()[i].RemoteID()] = i
	}
	node := &Node{
		State: StateCandidate.String(),
		PersistentState: &PersistentState{
			CurrentTerm: 0,
			VotedFor:    nil,
			SelfID:      cluster.OwnID(),
			PeerIPs:     cluster.Nodes(),
			ConnIDMap:   connIDMap,
		},
		VolatileState: &VolatileState{
			CommitIndex: -1,
			LastApplied: -1,
		},
		VolatileStateLeader: &VolatileStateLeader{
			NextIndex:  nextIndex,
			MatchIndex: matchIndex,
		},
	}
	return node
}

// Start starts a single raft node into beginning raft operations.
// This function starts the leader election and keeps a check on whether
// regular heartbeats to the node exists. It restarts leader election on failure to do so.
// This function also continuously listens on all the connections to the nodes
// and routes the requests to appropriate functions.
func (s *SimpleServer) Start() (err error) {
	// Making the function idempotent, returns whether the server is already open.
	s.lock.Lock()
	if s.node != nil {
		s.log.Debug().
			Str("self-id", s.node.PersistentState.SelfID.String()).
			Msg("already open")
		return network.ErrOpen
	}
	// Initialise all raft variables in this node.
	node := NewRaftNode(s.cluster)
	node.PersistentState.mu.Lock()
	node.log = s.log
	s.node = node
	s.lock.Unlock()
	selfID := node.PersistentState.SelfID
	node.PersistentState.mu.Unlock()

	ctx := context.Background()
	// liveChan is a channel that passes the incomingData once received.
	liveChan := make(chan *incomingData)
	// Listen forever on all node connections.

	go func() {
		for {
			// Parallely start waiting for incoming data.
			conn, msg, err := s.cluster.Receive(ctx)
			node.log.
				Debug().
				Str("self-id", selfID.String()).
				Str("received", msg.Kind().String()).
				Msg("received request")
			liveChan <- &incomingData{
				conn,
				msg,
			}
			if err != nil {
				return
			}
		}
	}()

	// This block of code checks what kind of request has to be serviced
	// and calls the necessary function to complete it.

	for {
		// If any sort of request (heartbeat,appendEntries,requestVote)
		// isn't received by the server(node) it restarts leader election.
		select {
		case <-s.timeoutProvider(node).C:
			s.lock.Lock()
			if s.node == nil {
				return
			}
			s.lock.Unlock()
			s.StartElection()
		case data := <-liveChan:
			err = s.processIncomingData(data)
			if err != nil {
				return
			}
		}
	}
}

// OnReplication is a handler setter.
func (s *SimpleServer) OnReplication(handler ReplicationHandler) {
	s.onReplication = handler
}

// Input appends the input log into the leaders log, only if the current node is the leader.
// If this was not a leader, the leaders data is communicated to the client.
func (s *SimpleServer) Input(input *message.Command) {
	s.node.PersistentState.mu.Lock()
	defer s.node.PersistentState.mu.Unlock()

	if s.node.State == StateLeader.String() {
		logData := message.NewLogData(s.node.PersistentState.CurrentTerm, input)
		s.node.PersistentState.Log = append(s.node.PersistentState.Log, logData)
	} else {
		// Relay data to leader.
		logAppendRequest := message.NewLogAppendRequest(input)

		s.relayDataToServer(logAppendRequest)
	}
}

// Close closes the node and returns an error on failure.
func (s *SimpleServer) Close() error {
	s.lock.Lock()
	// Maintaining idempotency of the close function.
	if s.node == nil {
		return network.ErrClosed
	}
	s.node.
		log.
		Debug().
		Str("self-id", s.node.PersistentState.SelfID.String()).
		Msg("closing node")

	s.node = nil
	err := s.cluster.Close()
	s.lock.Unlock()
	return err
}

// randomTimer returns tickers ranging from 150ms to 300ms.
func randomTimer(node *Node) *time.Timer {
	randomInt := rand.Intn(150) + 150
	node.log.
		Debug().
		Str("self-id", node.PersistentState.SelfID.String()).
		Int("random timer set to", randomInt).
		Msg("heart beat timer")
	ticker := time.NewTimer(time.Duration(randomInt) * time.Millisecond)
	return ticker
}

// processIncomingData is responsible for parsing the incoming data and calling
// appropriate functions based on the request type.
func (s *SimpleServer) processIncomingData(data *incomingData) error {

	ctx := context.TODO()

	switch data.msg.Kind() {
	case message.KindRequestVoteRequest:
		requestVoteRequest := data.msg.(*message.RequestVoteRequest)
		requestVoteResponse := s.node.RequestVoteResponse(requestVoteRequest)
		payload, err := message.Marshal(requestVoteResponse)
		if err != nil {
			return err
		}
		err = data.conn.Send(ctx, payload)
		if err != nil {
			return err
		}
	case message.KindAppendEntriesRequest:

		appendEntriesRequest := data.msg.(*message.AppendEntriesRequest)
		appendEntriesResponse := s.AppendEntriesResponse(appendEntriesRequest)
		payload, err := message.Marshal(appendEntriesResponse)
		if err != nil {
			return err
		}
		err = data.conn.Send(ctx, payload)
		if err != nil {
			return err
		}
	// When the leader gets a forwarded append input message from one of it's followers.
	case message.KindLogAppendRequest:
		logAppendRequest := data.msg.(*message.LogAppendRequest)
		input := logAppendRequest.Data
		logData := message.NewLogData(s.node.PersistentState.CurrentTerm, input)
		s.node.PersistentState.Log = append(s.node.PersistentState.Log, logData)
	}
	return nil
}

// relayDataToServer sends the input log from the follower to a leader node.
// TODO: Figure out what to do with the errors generated here.
func (s *SimpleServer) relayDataToServer(req *message.LogAppendRequest) {
	ctx := context.Background()

	payload, _ := message.Marshal(req)

	leaderNodeConn := s.cluster.Nodes()[s.node.PersistentState.ConnIDMap[s.node.PersistentState.LeaderID]]
	_ = leaderNodeConn.Send(ctx, payload)
}

// OnRequestVotes is a hook setter for RequestVotesRequest.
func (s *SimpleServer) OnRequestVotes(hook func(*message.RequestVoteRequest)) {
	s.onRequestVotes = hook
}

// OnLeaderElected is a hook setter for LeadeElectedRequest.
func (s *SimpleServer) OnLeaderElected(hook func()) {
	s.onLeaderElected = hook
}

// OnAppendEntries is a hook setter for AppenEntriesRequest.
func (s *SimpleServer) OnAppendEntries(hook func(*message.AppendEntriesRequest)) {
	s.onAppendEntries = hook
}
