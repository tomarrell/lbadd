package raft

// becomeFollower converts a leader to a follower.
// After this function is executed, the node goes back to the loop in raft.go
// thus resuming normal operations.
func becomeFollower(node *Node) {
	node.log.
		Debug().
		Str("self-id", node.PersistentState.SelfID.String()).
		Msg("becoming follower")
	node.PersistentState.LeaderID = nil
	node.PersistentState.VotedFor = nil
	node.State = StateFollower.String()
}
