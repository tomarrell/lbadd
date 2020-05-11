package raft

import "github.com/tomarrell/lbadd/internal/raft/message"

// AppendEntriesResponse function is called on a request from the leader to append log data
// to the follower node. This function generates the response to be sent to the leader node.
func AppendEntriesResponse(node Node, req *message.AppendEntriesRequest) *message.AppendEntriesResponse {

	return nil
}
