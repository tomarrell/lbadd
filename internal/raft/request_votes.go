package raft

import (
	"context"
	"time"

	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
	"google.golang.org/protobuf/proto"
)

// RequestVote enables a node to send out the RequestVotes RPC.
// This function requests a vote from one node and returns that node's response.
// It opens a connection to the intended node using the network layer and waits for a response.
func RequestVote(req *message.RequestVoteRequest) (*message.RequestVoteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := network.DialTCP(ctx, "x")
	if err != nil {
		return nil, err
	}

	payload, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = conn.Send(ctx, payload)
	if err != nil {
		return nil, err
	}

	res, err := conn.Receive(ctx)
	if err != nil {
		return nil, err
	}

	var message *message.RequestVoteResponse
	err = proto.Unmarshal(res, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

// RequestVoteResponse provides the response that a node must generate for a vote request.
func RequestVoteResponse(node Node, req *message.RequestVoteRequest) *message.RequestVoteResponse {
	node.PersistentState.mu.Lock()

	if node.PersistentState.VotedFor == nil {
		node.PersistentState.VotedFor = req.CandidateID
		return &message.RequestVoteResponse{
			Term:        node.PersistentState.CurrentTerm,
			VoteGranted: true,
		}
	}

	node.PersistentState.mu.Unlock()

	return &message.RequestVoteResponse{
		Term:        node.PersistentState.CurrentTerm,
		VoteGranted: false,
	}
}
