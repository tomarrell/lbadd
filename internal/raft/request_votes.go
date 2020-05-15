package raft

import (
	"context"
	"fmt"
	"time"

	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// RequestVote enables a node to send out the RequestVotes RPC.
// This function requests a vote from one node and returns that node's response.
// It opens a connection to the intended node using the network layer and waits for a response.
func RequestVote(nodeConn network.Conn, req *message.RequestVoteRequest) (*message.RequestVoteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// fmt.Println("DD")
	// fmt.Println(req)
	// fmt.Println("DD")

	payload, err := message.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = nodeConn.Send(ctx, payload)
	if err != nil {
		return nil, err
	}

	res, err := nodeConn.Receive(ctx)
	if err != nil {
		return nil, err
	}
	// fmt.Println("CC")
	// fmt.Println(string(res))
	// fmt.Println("CC")
	msg, err := message.Unmarshal(res)
	if err != nil {
		fmt.Printf("There is an err: %v\n", err)
		return nil, err
	}

	return msg.(*message.RequestVoteResponse), nil
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
