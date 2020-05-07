package raft

// RequestVotesRPCReq describes the data in a single RequestVotes request.
type RequestVotesRPCReq struct {
	Term         int
	CandidateID  int
	LastLogIndex int
	LastLogTerm  int
}

// RequestVotesRPCRes describes the data in a single RequestVotes response.
type RequestVotesRPCRes struct {
	Term        int
	VoteGranted bool
}

// RequestVotes enables a node to send out the RequestVotes RPC.
func RequestVotes() {

}
