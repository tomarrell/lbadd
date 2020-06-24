package raft

import (
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// AppendEntriesResponse function is called on a request from the leader to append log data
// to the follower node. This function generates the response to be sent to the leader node.
// This is the response to the contact by the leader to assert it's leadership.
func (s *simpleServer) AppendEntriesResponse(req *message.AppendEntriesRequest) *message.AppendEntriesResponse {
	leaderTerm := req.GetTerm()
	nodePersistentState := s.node.PersistentState
	nodeTerm := nodePersistentState.CurrentTerm
	// Return false if term is greater than currentTerm,
	// if msg Log Index is greater than node commit Index,
	// if term of msg at PrevLogIndex doesn't match prev Log Term stored by Leader.
	if nodeTerm > leaderTerm ||
		req.GetPrevLogIndex() > s.node.VolatileState.CommitIndex ||
		nodePersistentState.Log[req.PrevLogIndex].Term != req.GetPrevLogTerm() {
		s.node.log.
			Debug().
			Str("self-id", s.node.PersistentState.SelfID.String()).
			Str("returning failure to append entries to", string(req.GetLeaderID())).
			Msg("append entries failure")
		return &message.AppendEntriesResponse{
			Term:    nodeTerm,
			Success: false,
		}
	}

	entries := req.GetEntries()
	if len(entries) > 0 {
		nodePersistentState.mu.Lock()
		if req.GetPrevLogIndex() < s.node.VolatileState.CommitIndex {
			s.node.PersistentState.Log = s.node.PersistentState.Log[:req.GetPrevLogIndex()]
		}
		s.node.PersistentState.Log = append(s.node.PersistentState.Log, entries...)
		s.node.PersistentState.mu.Unlock()
	}

	if req.GetLeaderCommit() > s.node.VolatileState.CommitIndex {
		nodeCommitIndex := req.GetLeaderCommit()
		if int(req.GetLeaderCommit()) > len(s.node.PersistentState.Log) {
			nodeCommitIndex = int32(len(s.node.PersistentState.Log))
		}
		s.node.VolatileState.CommitIndex = nodeCommitIndex
		/* FIX ISSUE #152 from this
		commandEntries := getCommandFromLogs(entries)
		succeeded := s.onReplication(commandEntries)
		_ = succeeded
		// succeeded returns the number of applied entries.
		*/
	}

	s.node.log.
		Debug().
		Str("self-id", s.node.PersistentState.SelfID.String()).
		Str("returning success to append entries to", string(req.GetLeaderID())).
		Msg("append entries success")

	return &message.AppendEntriesResponse{
		Term:    nodeTerm,
		Success: true,
	}

}

func getCommandFromLogs(entries []*message.LogData) []*message.Command {
	var commandEntries []*message.Command
	for i := range entries {
		commandEntries = append(commandEntries, entries[i].Entry)
	}
	return commandEntries
}
