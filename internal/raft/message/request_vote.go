package message

import (
	"github.com/tomarrell/lbadd/internal/id"
)

//go:generate protoc --go_out=. request_vote.proto

var _ Message = (*RequestVoteRequest)(nil)
var _ Message = (*RequestVoteResponse)(nil)

func NewRequestVoteRequest(term int32, candidateID id.ID, lastLogIndex int32, lastLogTerm int32) *RequestVoteRequest {
	return &RequestVoteRequest{
		Term:         term,
		CandidateId:  candidateID.Bytes(),
		LastLogIndex: lastLogIndex,
		LastLogTerm:  lastLogTerm,
	}
}

func (*RequestVoteRequest) Kind() Kind {
	return KindRequestVoteRequest
}

func NewRequestVoteResponse(term int32, voteGranted bool) *RequestVoteResponse {
	return &RequestVoteResponse{
		Term:        term,
		VoteGranted: voteGranted,
	}
}

func (*RequestVoteResponse) Kind() Kind {
	return KindRequestVoteResponse
}
