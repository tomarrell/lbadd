package raft

import (
	"context"
	"fmt"
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
	Input(string)
	io.Closer
}

// ReplicationHandler is a handler setter.
type ReplicationHandler func(string)

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

	SelfID   id.ID
	LeaderID id.ID          // LeaderID is nil at init, and the id.ID of the node after the leader is elected.
	PeerIPs  []network.Conn // PeerIPs has the connection variables of all the other nodes in the cluster.
	mu       sync.Mutex
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

var _ Server = (*simpleServer)(nil)

// simpleServer implements a server in a cluster.
type simpleServer struct {
	node            *Node
	cluster         Cluster
	onReplication   ReplicationHandler
	log             zerolog.Logger
	timeoutProvider func(*Node) *time.Timer
	lock            sync.Mutex
	closeSignal     chan struct{}
}

// incomingData describes every request that the server gets.
type incomingData struct {
	conn network.Conn
	msg  message.Message
}

// NewServer enables starting a raft server/cluster.
func NewServer(log zerolog.Logger, cluster Cluster) Server {
	return newServer(log, cluster, nil)
}

func newServer(log zerolog.Logger, cluster Cluster, timeoutProvider func(*Node) *time.Timer) Server {
	if timeoutProvider == nil {
		timeoutProvider = randomTimer
	}
	// TODO: length needs to be figured out
	closingChannel := make(chan struct{}, 5)
	return &simpleServer{
		log:             log.With().Str("component", "raft").Logger(),
		cluster:         cluster,
		timeoutProvider: timeoutProvider,
		closeSignal:     closingChannel,
	}
}

// Start starts a single raft node into beginning raft operations.
// This function starts the leader election and keeps a check on whether
// regular heartbeats to the node exists. It restarts leader election on failure to do so.
// This function also continuously listens on all the connections to the nodes
// and routes the requests to appropriate functions.
func (s *simpleServer) Start() (err error) {
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
			s.lock.Lock()
			if s.node == nil {
				return
			}
			s.lock.Unlock()
		}
	}()

	// This block of code checks what kind of request has to be serviced
	// and calls the necessary function to complete it.

	// If any sort of request (heartbeat,appendEntries,requestVote)
	// isn't received by the server(node) it restarts leader election.
	select {
	case <-s.getDoneChan():
		return
	case <-s.timeoutProvider(node).C:
		s.lock.Lock()
		if s.node == nil {
			return
		}
		s.lock.Unlock()
		s.StartElection()
	case data := <-liveChan:
		err = node.processIncomingData(data)
		if err != nil {
			return
		}
	}
	return
}

func (s *simpleServer) OnReplication(handler ReplicationHandler) {
	s.onReplication = handler
}

// Input appends the input log into the leaders log, only if the current node is the leader.
// If this was not a leader, the leaders data is communicated to the client.
func (s *simpleServer) Input(input string) {
	s.node.PersistentState.mu.Lock()
	defer s.node.PersistentState.mu.Unlock()

	if s.node.State == StateLeader.String() {
		logData := message.NewLogData(s.node.PersistentState.CurrentTerm, input)
		s.node.PersistentState.Log = append(s.node.PersistentState.Log, logData)
	} else {
		// Relay data to leader.
		fmt.Println("TODO")
	}
}

// Close closes the node and returns an error on failure.
func (s *simpleServer) Close() error {
	s.lock.Lock()
	s.node.PersistentState.mu.Lock()
	// Maintaining idempotency of the close function.
	if s.node == nil {
		return network.ErrClosed
	}
	s.node.
		log.
		Debug().
		Str("self-id", s.node.PersistentState.SelfID.String()).
		Msg("closing node")

	s.node.PersistentState.mu.Unlock()
	s.node = nil
	err := s.cluster.Close()
	s.closeSignal <- struct{}{}
	s.lock.Unlock()
	return err
}

// NewRaftNode initialises a raft cluster with the given configuration.
func NewRaftNode(cluster Cluster) *Node {
	var nextIndex, matchIndex []int

	for range cluster.Nodes() {
		nextIndex = append(nextIndex, -1)
		matchIndex = append(matchIndex, -1)
	}
	node := &Node{
		State: StateCandidate.String(),
		PersistentState: &PersistentState{
			CurrentTerm: 0,
			VotedFor:    nil,
			SelfID:      cluster.OwnID(),
			PeerIPs:     cluster.Nodes(),
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
func (node *Node) processIncomingData(data *incomingData) error {

	ctx := context.TODO()

	switch data.msg.Kind() {
	case message.KindRequestVoteRequest:
		requestVoteRequest := data.msg.(*message.RequestVoteRequest)
		requestVoteResponse := node.RequestVoteResponse(requestVoteRequest)
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
		appendEntriesResponse := node.AppendEntriesResponse(appendEntriesRequest)
		payload, err := message.Marshal(appendEntriesResponse)
		if err != nil {
			return err
		}
		err = data.conn.Send(ctx, payload)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *simpleServer) getDoneChan() <-chan struct{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.closeSignal
}
