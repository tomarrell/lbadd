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
func (s *SimpleServer) StartElection() {

	s.node.PersistentState.mu.Lock()
	s.node.State = StateCandidate.String()
	s.node.PersistentState.CurrentTerm++
	var lastLogTerm, lastLogIndex int32
	savedCurrentTerm := s.node.PersistentState.CurrentTerm
	if len(s.node.PersistentState.Log) == 0 {
		lastLogTerm = 0
	} else {
		lastLogTerm = s.node.PersistentState.Log[len(s.node.PersistentState.Log)].Term
	}
	lastLogIndex = int32(len(s.node.PersistentState.Log))
	selfID := s.node.PersistentState.SelfID
	s.node.log.
		Debug().
		Str("self-id", selfID.String()).
		Int32("term", s.node.PersistentState.CurrentTerm+1).
		Msg("starting election")
	s.node.PersistentState.mu.Unlock()

	var votes int32

	for i := range s.node.PersistentState.PeerIPs {
		// Parallely request votes from the peers.
		go func(i int) {
			req := message.NewRequestVoteRequest(
				savedCurrentTerm,
				s.node.PersistentState.SelfID,
				lastLogIndex,
				lastLogTerm,
			)
			s.node.log.
				Debug().
				Str("self-id", selfID.String()).
				Str("request-vote sent to", s.node.PersistentState.PeerIPs[i].RemoteID().String()).
				Msg("request vote")

			// send a requestVotesRPC
			res, err := RequestVote(s.node.PersistentState.PeerIPs[i], req)
			// If there's an error, the vote is considered to be not casted by the node.
			// Worst case, there will be a re-election; the errors might be from network or
			// data consistency errors, which will be sorted by a re-election.
			// This decision was taken because, StartElection returning an error is not feasible.
			if res.VoteGranted && err == nil {
				s.node.log.
					Debug().
					Str("received vote from", s.node.PersistentState.PeerIPs[i].RemoteID().String()).
					Msg("voting from peer")
				votesRecieved := atomic.AddInt32(&votes, 1)

				// Check whether this node has already voted.
				// Else it can vote for itself.
				s.node.PersistentState.mu.Lock()
				defer s.node.PersistentState.mu.Unlock()

				if s.node.PersistentState.VotedFor == nil {
					s.node.PersistentState.VotedFor = selfID
					s.node.log.
						Debug().
						Str("self-id", selfID.String()).
						Msg("node voting for itself")
					votesRecieved++
				}

				if votesRecieved > int32(len(s.node.PersistentState.PeerIPs)/2) && s.node.State != StateLeader.String() {
					// This node has won the election.
					s.node.State = StateLeader.String()
					s.node.PersistentState.LeaderID = s.node.PersistentState.SelfID
					s.node.log.
						Debug().
						Str("self-id", selfID.String()).
						Msg("node elected leader")
					s.startLeader()
					return
				}
			}
		}(i)
	}
}
