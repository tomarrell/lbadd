package raft

// AppendEntriesRPCReq describes the data in an AppendEntries request.
type AppendEntriesRPCReq struct {
	Term         int
	LeaderID     int
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogData // The log entries.
	LeaderCommit int       // Leader's commit index.
}

// AppendEntriesRPCRes describes the data in an AppendEntries response.
type AppendEntriesRPCRes struct {
	Term    int  // The node's current term
	Success bool // Returns true if log matching property holds good, else false.
}
