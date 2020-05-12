package raft

import (
	"context"
	"fmt"
	"time"

	"github.com/tomarrell/lbadd/internal/id"
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

// RequestVoteResponse function is called on a request from a candidate for a vote. This function
// generates the response for the responder node to send back to the candidate node.
func RequestVoteResponse(node *Node, req *message.RequestVoteRequest) *message.RequestVoteResponse {
	node.PersistentState.mu.Lock()

	if node.PersistentState.VotedFor == nil {
		cID, err := id.Parse(req.CandidateID)
		if err != nil {
			// no point in handling this because I really need that to parse into ID.
			fmt.Println(err)
		}
		node.PersistentState.VotedFor = cID
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
