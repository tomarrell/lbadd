package raft

import (
	"context"
	"sync/atomic"

	"github.com/tomarrell/lbadd/internal/raft/message"
)

// StartElection enables a node in the cluster to start the election.
// It returns an error based on what happened if it cannot start the election.(?)
// The function caller doesn't need to wait for a voting response from this function,
// the function triggers the necessary functions responsible to continue the raft cluster
// into it's working stage if the node won the election.
func (s *SimpleServer) StartElection(ctx context.Context) {

	s.lock.Lock()
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
	numNodes := s.node.PersistentState.PeerIPs
	s.node.log.
		Debug().
		Str("self-id", selfID.String()).
		Int32("term", s.node.PersistentState.CurrentTerm+1).
		Msg("starting election")
	s.node.PersistentState.mu.Unlock()
	s.lock.Unlock()

	var votes int32

	for i := range numNodes {
		// Parallely request votes from the peers.
		go func(i int) {
			req := message.NewRequestVoteRequest(
				savedCurrentTerm,
				selfID,
				lastLogIndex,
				lastLogTerm,
			)
			s.lock.Lock()
			if s.node == nil {
				return
			}
			s.node.log.
				Debug().
				Str("self-id", selfID.String()).
				Str("request-vote sent to", s.node.PersistentState.PeerIPs[i].RemoteID().String()).
				Msg("request vote")

			nodeConn := s.node.PersistentState.PeerIPs[i]
			s.lock.Unlock()

			res, err := s.RequestVote(ctx, nodeConn, req)
			// If there's an error, the vote is considered to be not casted by the node.
			// Worst case, there will be a re-election; the errors might be from network or
			// data consistency errors, which will be sorted by a re-election.
			// This decision was taken because, StartElection returning an error is not feasible.
			if err == nil && res.VoteGranted {
				s.lock.Lock()
				s.node.log.
					Debug().
					Str("received vote from", s.node.PersistentState.PeerIPs[i].RemoteID().String()).
					Msg("voting from peer")
				s.lock.Unlock()
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
					s.node.PersistentState.LeaderID = selfID
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
