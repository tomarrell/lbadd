package raft

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

	PersistentState     PersistentState
	VolatileState       VolatileState
	VolatileStateLeader VolatileStateLeader
}

// PersistentState describes the persistent state data on a raft node.
type PersistentState struct {
	CurrentTerm int
	VotedFor    int
	Log         []LogData

	SelfID  int
	PeerIDs []int
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

// 1. How to contact other server
// 2. About raft init
// 3. Is ID enough to contact another server
// 4. Avoiding circular dependency
// 5. A method to log
