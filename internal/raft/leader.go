package raft

import (
	"context"
	"fmt"

	"github.com/tomarrell/lbadd/internal/raft/message"
)

// startLeader begins the leaders operations.
// The node passed as argument is the leader node.
// The leader begins by sending append entries RPC to the nodes.
// The leader sends periodic append entries request to the
// followers to keep them alive.

// TODO: Handle errors.
func startLeader(node *Node) {
	var log []message.LogData
	log = append(log, <-node.LogChannel)
	ctx := context.Background()
	appendEntriesRequest := message.NewAppendEntriesRequest(1, nil, 1, 1, nil, 1) // dummy request until I understand.
	for i := range node.PersistentState.PeerIPs {
		go func(i int) {
			payload, err := message.Marshal(appendEntriesRequest)
			if err != nil {
				fmt.Println(err)
			}
			err = node.PersistentState.PeerIPs[i].Send(ctx, payload)
			if err != nil {
				fmt.Println(err)
			}

			res, err := node.PersistentState.PeerIPs[i].Receive(ctx)
			if err != nil {
				fmt.Println(err)
			}

			resP, err := message.Unmarshal(res)
			if err != nil {
				fmt.Println(err)
			}

			appendEntriesResponse := resP.(*message.AppendEntriesResponse)
			fmt.Println(appendEntriesResponse)
		}(i)
	}
}
