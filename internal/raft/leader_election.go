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
// TODO: Logging.
func StartElection(node *Node) {
	node.State = StateCandidate.String()
	node.PersistentState.CurrentTerm++

	var votes int32

	for i := range node.PersistentState.PeerIPs {
		// Parallely request votes from the peers.
		go func(i int) {
			// send a requestVotesRPC
			lastLogTerm := 1 // TODO: index issue here
			req := message.NewRequestVoteRequest(
				int32(node.PersistentState.CurrentTerm),
				node.PersistentState.SelfID,
				int32(len(node.PersistentState.Log)),
				int32(lastLogTerm), //int32(node.PersistentState.Log[len(node.PersistentState.Log)].Term),
			)
			node.log.
				Debug().
				Str("self-id", node.PersistentState.SelfID.String()).
				Str("request-vote sent to", node.PersistentState.PeerIPs[i].RemoteID().String()).
				Msg("request vote")
			res, err := RequestVote(node.PersistentState.PeerIPs[i], req)
			// If there's an error, the vote is considered to be not casted by the node.
			// Worst case, there will be a re-election; the errors might be from network or
			// data consistency errors, which will be sorted by a re-election.
			// This decision was taken because, StartElection returning an error is not feasible.
			if res.VoteGranted && err == nil {
				node.log.
					Debug().
					Str("received vote from", node.PersistentState.PeerIPs[i].RemoteID().String()).
					Msg("voting from peer")
				votesRecieved := atomic.AddInt32(&votes, 1)
				// Check whether this node has already voted.
				// Else it can vote for itself.
				node.PersistentState.mu.Lock()
				if node.PersistentState.VotedFor == nil {
					node.PersistentState.VotedFor = node.PersistentState.SelfID
					node.log.
						Debug().
						Str("self-id", node.PersistentState.SelfID.String()).
						Msg("node voting for itself")
					votesRecieved++
				}
				node.PersistentState.mu.Unlock()

				if votesRecieved > int32(len(node.PersistentState.PeerIPs)/2) {
					// This node has won the election.
					node.State = StateLeader.String()
					node.PersistentState.LeaderID = node.PersistentState.SelfID
					node.log.
						Debug().
						Str("self-id", node.PersistentState.SelfID.String()).
						Msg("node elected leader")
					startLeader(node)
				}
			}
		}(i)
	}
}
