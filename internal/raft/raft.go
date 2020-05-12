package raft

import (
	"context"
	"io"
	"sync"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// Available states
const (
	LeaderState    = "leader"
	CandidateState = "candidate"
	FollowerState  = "follower"
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
	State string

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

type simpleServer struct {
	cluster       Cluster
	onReplication ReplicationHandler
	log           zerolog.Logger
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
	// Initialise all raft variables in this node.
	node := NewRaftCluster(s.cluster)

	ctx := context.Background()
	// Listen forever on all node connections. This block of code checks what kind of
	// request has to be serviced and calls the necessary function to complete it.
	go func() {
		conn, msg, err := s.cluster.Receive(ctx)
		if err != nil {
			return
		}
		switch msg.Kind() {
		case message.KindRequestVoteRequest:
			requestVoteRequest := msg.(*message.RequestVoteRequest)
			requestVoteResponse := RequestVoteResponse(node, requestVoteRequest)
			payload, err := message.Marshal(requestVoteResponse)
			if err != nil {
				return
			}
			err = conn.Send(ctx, payload)
			if err != nil {
				return
			}
		case message.KindAppendEntriesRequest:
			appendEntriesRequest := msg.(*message.AppendEntriesRequest)
			appendEntriesResponse := AppendEntriesResponse(node, appendEntriesRequest)
			payload, err := message.Marshal(appendEntriesResponse)
			if err != nil {
				return
			}
			err = conn.Send(ctx, payload)
			if err != nil {
				return
			}
		}
	}()

	go func() {
		StartElection(node)
	}()

	// check for heartbeats

	return nil
}

func (s *simpleServer) OnReplication(handler ReplicationHandler) {
	s.onReplication = handler
}

func (s *simpleServer) Input(string) {

}

func (s *simpleServer) Close() error {
	return nil
}

// NewRaftCluster initialises a raft cluster with the given configuration.
func NewRaftCluster(cluster Cluster) *Node {
	node := &Node{
		PersistentState: &PersistentState{
			CurrentTerm: 0,
			VotedFor:    nil,
			SelfID:      nil, // TODO: add node's global ID once done in NW layer.
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
