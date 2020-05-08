package raft

import (
	"github.com/rs/zerolog"
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

// State describes the current state of a raft node.
type State struct {
	Name string

	PersistentState     *PersistentState
	VolatileState       *VolatileState
	VolatileStateLeader *VolatileStateLeader

	log zerolog.Logger
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int
	VotedFor    int
	Log         []LogData

	SelfID  int
	SelfIP  network.Conn
	PeerIPs []network.Conn
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
func NewRaftCluster(cluster Cluster) []*State {
	var ClusterStates []*State
	sampleState := &State{
		PersistentState:     &PersistentState{},
		VolatileState:       &VolatileState{},
		VolatileStateLeader: &VolatileStateLeader{},
	}

	for i := range cluster.Nodes() {
		var state *State
		state = sampleState
		state.PersistentState.CurrentTerm = 0
		state.PersistentState.VotedFor = -1
		state.PersistentState.SelfIP = cluster.Nodes()[i]
		state.PersistentState.PeerIPs = cluster.Nodes()

		state.VolatileState.CommitIndex = -1
		state.VolatileState.LastApplied = -1

		ClusterStates = append(ClusterStates, state)
	}
	return ClusterStates
}
