package message

import (
	"github.com/tomarrell/lbadd/internal/id"
)

//go:generate protoc --go_out=. request_vote.proto

var _ Message = (*RequestVoteRequest)(nil)
var _ Message = (*RequestVoteResponse)(nil)

// NewRequestVoteRequest creates a new request-vote-request message with the
// given parameters.
func NewRequestVoteRequest(term int32, candidateID id.ID, lastLogIndex int32, lastLogTerm int32) *RequestVoteRequest {
	return &RequestVoteRequest{
		Term:         term,
		CandidateID:  candidateID.Bytes(),
		LastLogIndex: lastLogIndex,
		LastLogTerm:  lastLogTerm,
	}
}

// Kind returns KindRequestVoteRequest.
func (*RequestVoteRequest) Kind() Kind {
	return KindRequestVoteRequest
}

// NewRequestVoteResponse creates a new request-vote-response message with the
// given parameters.
func NewRequestVoteResponse(term int32, voteGranted bool) *RequestVoteResponse {
	return &RequestVoteResponse{
		Term:        term,
		VoteGranted: voteGranted,
	}
}

// Kind returns KindRequestVoteResponse.
func (*RequestVoteResponse) Kind() Kind {
	return KindRequestVoteResponse
}
