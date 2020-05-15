package raft

import (
	"sync/atomic"

	"github.com/tomarrell/lbadd/internal/raft/message"
)

// StartElection enables a node in the cluster to start the election.
// It returns an error based on what happened if it cannot start the election.(?)
// The function caller doesn't need to wait for a voting response from this function,
// the function triggers the necessary functions responsible to continue the raft cluster
// into it's working stage if the node won the election.
func StartElection(node *Node) {
	node.State = StateCandidate.String()
	node.PersistentState.CurrentTerm++

	var votes int32

	for i := range node.PersistentState.PeerIPs {
		// Parallely request votes from the peers.
		go func(i int) {
			// send a requestVotesRPC
			req := message.NewRequestVoteRequest(
				int32(node.PersistentState.CurrentTerm),
				node.PersistentState.SelfID,
				int32(len(node.PersistentState.Log)),
				int32(node.PersistentState.Log[len(node.PersistentState.Log)-1].Term),
			)
			res, err := RequestVote(node.PersistentState.PeerIPs[i], req)
			// If they are (un)/marshalling errors, we probably should retry.
			// Because it doesnt mean that the node denied the vote.
			// Opposing view - failure is a failure, network or software,
			// we can assume the error is an error for whatever reasona and
			// proceed without having this vote.
			if err != nil {
				if res.VoteGranted {
					votesRecieved := atomic.AddInt32(&votes, 1)
					// Check whether this node has already voted.
					// Else it can vote for itself.
					node.PersistentState.mu.Lock()
					if node.PersistentState.VotedFor == nil {
						node.PersistentState.VotedFor = node.PersistentState.SelfID
						votesRecieved++
					}
					node.PersistentState.mu.Unlock()

					if votesRecieved > int32(len(node.PersistentState.PeerIPs)/2) {
						// This node has won the election.
						node.State = StateLeader.String()
						startLeader(node)
					}
				}
			}
		}(i)
	}

}
