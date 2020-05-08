package raft

// StartElection enables a node in the cluster to start the election.
func StartElection(Server State) {
	Server.Name = CandidateState
	Server.PersistentState.CurrentTerm++

	var votes int

	for i := range Server.PersistentState.PeerIPs {
		// parallely request votes from all the other peers.
		go func(i int) {
			if Server.PersistentState.PeerIPs[i] != Server.PersistentState.SelfIP {
				// send a requestVotesRPC
				req := &RequestVoteRPCReq{
					Term:         Server.PersistentState.CurrentTerm,
					CandidateID:  Server.PersistentState.SelfID,
					LastLogIndex: len(Server.PersistentState.Log),
					LastLogTerm:  Server.PersistentState.Log[len(Server.PersistentState.Log)-1].Term,
				}
				res := RequestVote(req)
				if res.VoteGranted {
					votes++
				}
			}
		}(i)
	}
}
