package raft

import "github.com/tomarrell/lbadd/internal/raft/message"

// AppendEntriesResponse provides the response that a node must generate for an append entries request.
func AppendEntriesResponse(node Node, req *message.AppendEntriesRequest) *message.AppendEntriesResponse {

	return nil
}
