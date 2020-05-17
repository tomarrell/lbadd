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
func StartElection(node *Node) (err error) {
	node.State = StateCandidate.String()
	node.PersistentState.CurrentTerm++

	var votes int32

	// fmt.Println(len(node.PersistentState.PeerIPs))
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
			// s.log.Printf("%v sent RequestVoteRPC to %v", node.PersistentState.SelfID, node.PersistentState.PeerIPs[i])
			res, err := RequestVote(node.PersistentState.PeerIPs[i], req)
			if err != nil {
				return
			}
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
					node.PersistentState.LeaderID = node.PersistentState.SelfID
					_ = startLeader(node)
				}
			}
		}(i)
	}

	return nil
}
