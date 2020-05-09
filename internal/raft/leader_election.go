package raft

import "github.com/tomarrell/lbadd/internal/raft/message"

// StartElection enables a node in the cluster to start the election.
func StartElection(server Node) {
	server.State = CandidateState
	server.PersistentState.CurrentTerm++

	var votes int

	for i := range server.PersistentState.PeerIPs {
		// parallely request votes from all the other peers.
		go func(i int) {
			if server.PersistentState.PeerIPs[i] != server.PersistentState.SelfIP {
				// send a requestVotesRPC
				req := message.NewRequestVoteRequest(
					int32(server.PersistentState.CurrentTerm),
					server.PersistentState.SelfID,
					int32(len(server.PersistentState.Log)),
					int32(server.PersistentState.Log[len(server.PersistentState.Log)-1].Term),
				)
				res, err := RequestVote(req)
				// If they are (un)/marshalling errors, we probably should retry.
				// Because it doesnt mean that the server denied the vote.
				// Opposing view - failure is a failure, network or software,
				// we can assume the error is an error for whatever reasona and
				// proceed without having this vote.
				if err != nil {
					if res.VoteGranted {
						votes++
					}
				}
			}
		}(i)
	}
}
