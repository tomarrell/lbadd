package raft

import (
	"context"
	"time"

	"github.com/tomarrell/lbadd/internal/network"
)

// RequestVoteRPCReq describes the data in a single RequestVotes request.
type RequestVoteRPCReq struct {
	Term         int // Candidate's term
	CandidateID  int
	LastLogIndex int
	LastLogTerm  int
}

// RequestVoteRPCRes describes the data in a single RequestVotes response.
type RequestVoteRPCRes struct {
	Term        int
	VoteGranted bool
}

// RequestVote enables a node to send out the RequestVotes RPC.
// This function requests a vote from one node and returns that node's response.
// It opens a connection to the intended node using the network layer and waits for a response.
func RequestVote(req *RequestVoteRPCReq) RequestVoteRPCRes {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := network.DialTCP(ctx, "x")
	if err != nil {

	}
	payload := []byte("protobuf serialised version of req")
	err = conn.Send(ctx, payload)
	if err != nil {

	}

	res, err := conn.Receive(ctx)
	if err != nil {

	}

	return unSerialise(res)
}

func unSerialise(res []byte) RequestVoteRPCRes {
	return RequestVoteRPCRes{}
}
