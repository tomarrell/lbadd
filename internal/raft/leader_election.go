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

	node.PersistentState.mu.Lock()

	node.State = StateCandidate.String()
	node.PersistentState.CurrentTerm++
	var lastLogTerm, lastLogIndex int32
	savedCurrentTerm := node.PersistentState.CurrentTerm
	if len(node.PersistentState.Log) == 0 {
		lastLogTerm = 0
	} else {
		lastLogTerm = node.PersistentState.Log[len(node.PersistentState.Log)].Term
	}
	lastLogIndex = int32(len(node.PersistentState.Log))

	node.PersistentState.mu.Unlock()

	var votes int32

	for i := range node.PersistentState.PeerIPs {
		// Parallely request votes from the peers.
		go func(i int) {
			req := message.NewRequestVoteRequest(
				savedCurrentTerm,
				node.PersistentState.SelfID,
				lastLogIndex,
				lastLogTerm,
			)

			node.log.
				Debug().
				Str("self-id", node.PersistentState.SelfID.String()).
				Str("request-vote sent to", node.PersistentState.PeerIPs[i].RemoteID().String()).
				Msg("request vote")

			// send a requestVotesRPC
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
				defer node.PersistentState.mu.Unlock()
				if node.PersistentState.VotedFor == nil {
					node.PersistentState.VotedFor = node.PersistentState.SelfID
					node.log.
						Debug().
						Str("self-id", node.PersistentState.SelfID.String()).
						Msg("node voting for itself")
					votesRecieved++
				}

				if votesRecieved > int32(len(node.PersistentState.PeerIPs)/2) {
					// This node has won the election.
					node.State = StateLeader.String()
					node.PersistentState.LeaderID = node.PersistentState.SelfID
					node.log.
						Debug().
						Str("self-id", node.PersistentState.SelfID.String()).
						Msg("node elected leader")
					// node.PersistentState.mu.Unlock()
					startLeader(node)
					return
				}
			}
		}(i)
	}
}
