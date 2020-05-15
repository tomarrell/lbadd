package raft

import (
	"context"
	"fmt"
	"time"

	"github.com/tomarrell/lbadd/internal/raft/message"
)

// startLeader begins the leaders operations.
// The node passed as argument is the leader node.
// The leader begins by sending append entries RPC to the nodes.
// The leader sends periodic append entries request to the
// followers to keep them alive.
// Empty append entries request are also called heartbeats.
// The data that goes in the append entries request is determined by
// existance of data in the LogChannel channel.

// TODO: Log errors.
func startLeader(node *Node) (err error) {
	var logs []*message.LogData

	ctx := context.TODO()

	var appendEntriesRequest *message.AppendEntriesRequest
	// The loop that the leader stays in until it's functioning properly.
	// The goal of this loop is to maintain raft in it's working phase;
	// periodically sending heartbeats/appendEntries.
	for {
		// Send heartbeats every 50ms.
		<-time.NewTimer(50 * time.Millisecond).C

		// If there is an input from the channel, grab it and add it to the outgoing request.
		select {
		case singleLog := <-node.LogChannel:
			logs = append(logs, &singleLog)
			appendEntriesRequest = message.NewAppendEntriesRequest(
				node.PersistentState.CurrentTerm,
				node.PersistentState.SelfID,
				1,
				1,
				logs,
				1,
			)
		default:
			appendEntriesRequest = message.NewAppendEntriesRequest(
				node.PersistentState.CurrentTerm,
				node.PersistentState.SelfID,
				1,
				1,
				logs,
				1,
			) // dummy request until I understand.
		}

		// Parallely send AppendEntriesRPC to all followers.
		for i := range node.PersistentState.PeerIPs {
			go func(i int) {
				payload, err := message.Marshal(appendEntriesRequest)
				if err != nil {
					return
				}
				err = node.PersistentState.PeerIPs[i].Send(ctx, payload)
				if err != nil {
					return
				}

				res, err := node.PersistentState.PeerIPs[i].Receive(ctx)
				if err != nil {
					return
				}

				resP, err := message.Unmarshal(res)
				if err != nil {
					return
				}

				appendEntriesResponse := resP.(*message.AppendEntriesResponse)
				// TODO: Based on the response, retries etc must be conducted.
				fmt.Println(appendEntriesResponse)
			}(i)
		}
	}

}
