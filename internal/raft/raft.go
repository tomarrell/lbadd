package raft

// LogData is a single log entry
type LogData struct {
}

// State describes the current state of a raft node.
type State struct {
	PersistentState     PersistentState
	VolatileState       VolatileState
	VolatileStateLeader VolatileStateLeader
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int
	VotedFor    int
	Log         []LogData
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
