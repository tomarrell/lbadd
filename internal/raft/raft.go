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
	Input(string)
	io.Closer
}

// Cluster is a description of a cluster of servers.
type Cluster interface {
	OwnID() id.ID
	Nodes() []network.Conn
	Receive(context.Context) (network.Conn, message.Message, error)
	Broadcast(context.Context, message.Message) error
}

// ReplicationHandler is a handler setter.
type ReplicationHandler func(string)

// Node describes the current state of a raft node.
// The raft paper describes this as a "State" but node
// seemed more intuitive.
type Node struct {
	State      string
	LogChannel chan (message.LogData) // LogChannel is used to store the incoming logs from clients.

	PersistentState     *PersistentState
	VolatileState       *VolatileState
	VolatileStateLeader *VolatileStateLeader
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int32
	VotedFor    id.ID // VotedFor is nil at init, -1 if the node voted for itself and any number in the slice Nodes() to point at its voter.
	Log         []message.LogData

	SelfID   id.ID
	LeaderID id.ID          // LeaderID is nil at init, -1 if its the leader and any number in the slice Nodes() to point at the leader.
	PeerIPs  []network.Conn // PeerIPs has the connection variables of all the other nodes in the cluster.
	mu       sync.Mutex
}

// VolatileState describes the volatile state data on a raft node.
type VolatileState struct {
	CommitIndex int
	LastApplied int
}

// VolatileStateLeader describes the volatile state data that exists on a raft leader.
type VolatileStateLeader struct {
	NextIndex  []int // Holds the nextIndex value for each of the followers in the cluster.
	MatchIndex []int // Holds the matchIndex value for each of the followers in the cluster.
}

var _ Server = (*simpleServer)(nil)

// simpleServer implements a server in a cluster.
type simpleServer struct {
	node          *Node
	cluster       Cluster
	onReplication ReplicationHandler
	log           zerolog.Logger
}

// incomingData describes every request that the server gets.
type incomingData struct {
	conn network.Conn
	msg  message.Message
}

// NewServer enables starting a raft server/cluster.
func NewServer(log zerolog.Logger, cluster Cluster) Server {
	return &simpleServer{
		log:     log.With().Str("component", "raft").Logger(),
		cluster: cluster,
	}
}

// Start starts a single raft node into beginning raft operations.
// This function starts the leader election and keeps a check on whether
// regular heartbeats to the node exists. It restarts leader election on failure to do so.
// This function also continuously listens on all the connections to the nodes
// and routes the requests to appropriate functions.
func (s *simpleServer) Start() (err error) {
	// Making the function idempotent, returns whether the server is already open.
	if s.node != nil {
		return network.ErrOpen
	}

	// Initialise all raft variables in this node.
	node := NewRaftNode(s.cluster)
	s.node = node

	ctx := context.Background()
	// liveChan is a channel that passes the incomingData once received.
	liveChan := make(chan *incomingData)
	// Listen forever on all node connections. This block of code checks what kind of
	// request has to be serviced and calls the necessary function to complete it.
	go func() {
		for {
			// Parallely start waiting for incoming data.
			conn, msg, err := s.cluster.Receive(ctx)
			liveChan <- &incomingData{
				conn,
				msg,
			}
			if err != nil {
				return
			}
		}
	}()

	for {
		// If any sort of request (heartbeat,appendEntries,requestVote)
		// isn't received by the server(node) it restarts leader election.
		select {
		case <-randomTimer().C:
			StartElection(node)
		case data := <-liveChan:
			err = processIncomingData(data, node)
			if err != nil {
				return
			}
		}
	}
}

func (s *simpleServer) OnReplication(handler ReplicationHandler) {
	s.onReplication = handler
}

// Input pushes the input data in the for of log data into the LogChannel.
// AppendEntries, whenever it occurs will be sent by obtaining data out of this channel.
func (s *simpleServer) Input(input string) {
	logData := message.NewLogData(s.node.PersistentState.CurrentTerm, input)
	s.node.LogChannel <- *logData
}

// Close closes the node and returns an error on failure.
func (s *simpleServer) Close() error {
	// Maintaining idempotency of the close function.
	if s.node == nil {
		return network.ErrClosed
	}
	// TODO: must close all operations gracefully.
	return nil
}

// NewRaftNode initialises a raft cluster with the given configuration.
func NewRaftNode(cluster Cluster) *Node {
	node := &Node{
		State:      StateCandidate.String(),
		LogChannel: make(chan message.LogData),
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
		VolatileStateLeader: &VolatileStateLeader{},
	}
	return node
}

// randomTimer returns tickers ranging from 150ms to 300ms.
func randomTimer() *time.Timer {
	randomInt := rand.Intn(150) + 150
	ticker := time.NewTimer(time.Duration(randomInt) * time.Millisecond)
	return ticker
}

// processIncomingData is responsible for parsing the incoming data and calling
// appropriate functions based on the request type.
func processIncomingData(data *incomingData, node *Node) error {

	ctx := context.TODO()

	switch data.msg.Kind() {
	case message.KindRequestVoteRequest:
		requestVoteRequest := data.msg.(*message.RequestVoteRequest)
		requestVoteResponse := RequestVoteResponse(node, requestVoteRequest)
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
		appendEntriesResponse := AppendEntriesResponse(node, appendEntriesRequest)
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
