package raft

import (
	"context"
	"time"

	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// RequestVote enables a node to send out the RequestVotes RPC.
// This function requests a vote from one node and returns that node's response.
// It opens a connection to the intended node using the network layer and waits for a response.
func RequestVote(req *message.RequestVoteRequest) *message.RequestVoteResponse {
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
	_ = res

	return nil
}
