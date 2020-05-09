package raft

import (
	"sync"

	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
)

// Server representsa a raft server.
type Server interface {
	LeaderElection()
	RequestVotes()
	AppendEnttries()
}

// Available states
const (
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

	// log zerolog.Logger
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
