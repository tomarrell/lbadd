package raft

import (
	"context"
	"sync"

	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// NewServer enables starting a raft server/cluster.
func NewServer(Cluster) Server

// Server is a description of a raft server.
type Server interface {
	Start() error
	OnReplication(ReplicationHandler)
	Input(string)
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

	// log zerolog.Logger
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

// NewRaftCluster initialises a raft cluster with the given configuration.
func NewRaftCluster(cluster cluster.Cluster) []*Node {
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
