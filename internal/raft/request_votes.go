package raft

import (
	"context"
	"fmt"

	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// RequestVote enables a node to send out the RequestVotes RPC.
// This function requests a vote from one node and returns that node's response.
// It opens a connection to the intended node using the network layer and waits for a response.
func RequestVote(nodeConn network.Conn, req *message.RequestVoteRequest) (*message.RequestVoteResponse, error) {
	ctx := context.Background()

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

	msg, err := message.Unmarshal(res)
	if err != nil {
		return nil, err
	}

	return msg.(*message.RequestVoteResponse), nil
}

// RequestVoteResponse function is called on a request from a candidate for a vote. This function
// generates the response for the responder node to send back to the candidate node.
func (node *Node) RequestVoteResponse(req *message.RequestVoteRequest) *message.RequestVoteResponse {
	node.PersistentState.mu.Lock()
	defer node.PersistentState.mu.Unlock()

	// If the candidate is not up to date with the term, reject the vote.
	if req.Term < node.PersistentState.CurrentTerm {
		return &message.RequestVoteResponse{
			Term:        node.PersistentState.CurrentTerm,
			VoteGranted: false,
		}
	}

	// If this node hasn't voted for any other node, vote only then.
	// TODO: Check whether candidate's log is atleast as up to date as mine only then grant vote.
	if node.PersistentState.VotedFor == nil {
		cID, err := id.Parse(req.CandidateID)
		if err != nil {
			// no point in handling this because I really need that to parse into ID.
			fmt.Println(err)
		}
		node.PersistentState.VotedFor = cID
		node.log.
			Debug().
			Str("self-id", node.PersistentState.SelfID.String()).
			Str("vote granted to", cID.String()).
			Msg("voting a peer")
		return &message.RequestVoteResponse{
			Term:        node.PersistentState.CurrentTerm,
			VoteGranted: true,
		}
	}

	return &message.RequestVoteResponse{
		Term:        node.PersistentState.CurrentTerm,
		VoteGranted: false,
	}
}
