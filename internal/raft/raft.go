package raft

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
)

var (
	LeaderState    = "leader"
	CandidateState = "candidate"
	FollowerState  = "follower"
)

// LogData is a single log entry
type LogData struct {
	Term int // Term  where this log was appended
	Data string
}

// Node describes the current state of a raft node.
// The raft paper describes this as a "State" but node
// seemed more intuitive.
type Node struct {
	State string

	PersistentState     *PersistentState
	VolatileState       *VolatileState
	VolatileStateLeader *VolatileStateLeader

	log zerolog.Logger
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int32
	VotedFor    []byte
	Log         []LogData

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
func NewRaftCluster(cluster Cluster) []*Node {
	var ClusterNodes []*Node
	sampleState := &Node{
		PersistentState:     &PersistentState{},
		VolatileState:       &VolatileState{},
		VolatileStateLeader: &VolatileStateLeader{},
	}

	for i := range cluster.Nodes() {
		var node *Node
		node = sampleState
		node.PersistentState.CurrentTerm = 0
		node.PersistentState.VotedFor = nil
		node.PersistentState.SelfIP = cluster.Nodes()[i]
		node.PersistentState.PeerIPs = cluster.Nodes()

		node.VolatileState.CommitIndex = -1
		node.VolatileState.LastApplied = -1

		ClusterNodes = append(ClusterNodes, node)
	}
	return ClusterNodes
}
