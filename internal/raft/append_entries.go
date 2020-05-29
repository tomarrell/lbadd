package raft

import "github.com/tomarrell/lbadd/internal/raft/message"

// AppendEntriesResponse function is called on a request from the leader to append log data
// to the follower node. This function generates the response to be sent to the leader node.
// This is the response to the contact by the leader to assert it's leadership.
func AppendEntriesResponse(node *Node, req *message.AppendEntriesRequest) *message.AppendEntriesResponse {
	leaderTerm := req.GetTerm()
	nodePersistentState := node.PersistentState
	nodeTerm := nodePersistentState.CurrentTerm
	// Return false if term is greater than currentTerm,
	// if msg Log Index is greater than node commit Index,
	// if term of msg at PrevLogIndex doesn't match prev Log Term stored by Leader.
	if nodeTerm > leaderTerm ||
		req.GetPrevLogIndex() > node.VolatileState.CommitIndex ||
		nodePersistentState.Log[req.PrevLogIndex].Term != req.GetPrevLogTerm() {
		node.log.
			Debug().
			Str("self-id", node.PersistentState.SelfID.String()).
			Str("returning failure to append entries to", string(req.GetLeaderID())).
			Msg("append entries failure")
		return &message.AppendEntriesResponse{
			Term:    nodeTerm,
			Success: false,
		}
	}

	entries := req.GetEntries()
	// if heartbeat, skip adding entries
	if len(entries) > 0 {
		nodePersistentState.mu.Lock()
		if req.GetPrevLogIndex() < node.VolatileState.CommitIndex {
			node.PersistentState.Log = node.PersistentState.Log[:req.GetPrevLogIndex()]
		}
		node.PersistentState.Log = append(node.PersistentState.Log, entries...)
		node.PersistentState.mu.Unlock()
	}

	if req.GetLeaderCommit() > node.VolatileState.CommitIndex {
		nodeCommitIndex := req.GetLeaderCommit()
		if int(req.GetLeaderCommit()) > len(node.PersistentState.Log) {
			nodeCommitIndex = int32(len(node.PersistentState.Log))
		}
		node.VolatileState.CommitIndex = nodeCommitIndex
		// TODO: Issue #152 apply the log command & update lastApplied
	}

	node.log.
		Debug().
		Str("self-id", node.PersistentState.SelfID.String()).
		Str("returning success to append entries to", string(req.GetLeaderID())).
		Msg("append entries success")

	return &message.AppendEntriesResponse{
		Term:    nodeTerm,
		Success: true,
	}

}
