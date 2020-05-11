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

// Server is a description of a raft server.
type Server interface {
	Start() error
	OnReplication(ReplicationHandler)
	Input(string)
	io.Closer
}

// ReplicationHandler is a handler setter.
type ReplicationHandler func(string)

// Cluster is a description of a cluster of servers.
type Cluster interface {
	Nodes() []network.Conn
	Receive(context.Context) (network.Conn, message.Message, error)
	Broadcast(context.Context, message.Message) error
}

// Available states
const (
	LeaderState    = "leader"
	CandidateState = "candidate"
	FollowerState  = "follower"
)

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
	VotedFor    id.ID
	Log         []message.LogData

	SelfID  id.ID
	SelfIP  network.Conn
	PeerIPs []network.Conn
	mu      sync.Mutex
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

// Start starts the raft servers.
// It creates a new cluster and returns an error,
// if there was one in the process.
// This function starts the leader election and keeps a check on whether
// regular heartbeats exist. It restarts leader election on failure to do so.
// This function also continuously listens on all the connections to the nodes
// and routes the requests to appropriate functions.
func (s *simpleServer) Start() error {
	nodes := NewRaftCluster(s.cluster)
	_ = nodes
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
func NewRaftCluster(cluster Cluster) []*Node {
	var clusterNodes []*Node

	for i := range cluster.Nodes() {
		node := &Node{
			PersistentState: &PersistentState{
				CurrentTerm: 0,
				VotedFor:    nil,
				SelfIP:      cluster.Nodes()[i],
				PeerIPs:     cluster.Nodes(),
			},
			VolatileState: &VolatileState{
				CommitIndex: -1,
				LastApplied: -1,
			},
			VolatileStateLeader: &VolatileStateLeader{},
		}

		clusterNodes = append(clusterNodes, node)
	}
	return clusterNodes
}
