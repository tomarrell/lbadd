package raft

import (
	"context"
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
func startLeader(node *Node) {

	node.log.
		Debug().
		Str("self-id", node.PersistentState.SelfID.String()).
		Msg("starting leader election proceedings")
	go func() {
		// The loop that the leader stays in until it's functioning properly.
		// The goal of this loop is to maintain raft in it's working phase;
		// periodically sending heartbeats/appendEntries.
		// This loop goes on until this node is the leader.
		for {
			// Send heartbeats every 50ms.
			<-time.NewTimer(50 * time.Millisecond).C

			sendHeartBeats(node)

			node.PersistentState.mu.Lock()
			if node.State != StateLeader.String() {
				node.PersistentState.mu.Unlock()
				return
			}
			node.PersistentState.mu.Unlock()
		}
	}()
}

func sendHeartBeats(node *Node) {
	ctx := context.TODO()

	node.PersistentState.mu.Lock()
	savedCurrentTerm := node.PersistentState.CurrentTerm
	node.PersistentState.mu.Unlock()

	var appendEntriesRequest *message.AppendEntriesRequest

	// Parallely send AppendEntriesRPC to all followers.
	for i := range node.PersistentState.PeerIPs {
		node.log.
			Debug().
			Str("self-id", node.PersistentState.SelfID.String()).
			Msg("sending heartbeats")
		go func(i int) {
			node.PersistentState.mu.Lock()
			nextIndex := node.VolatileStateLeader.NextIndex[i]
			prevLogIndex := nextIndex
			prevLogTerm := -1
			if prevLogIndex >= 0 {
				prevLogTerm = int(node.PersistentState.Log[prevLogIndex].Term)
			}

			// Logs are included from the nextIndex value to the current appended values
			// in the leader node. If there are none, no logs will be appended.
			entries := node.PersistentState.Log[nextIndex:]

			appendEntriesRequest = message.NewAppendEntriesRequest(
				node.PersistentState.CurrentTerm,
				node.PersistentState.SelfID,
				int32(prevLogIndex),
				int32(prevLogTerm),
				entries,
				node.VolatileState.CommitIndex,
			)
			node.PersistentState.mu.Unlock()

			payload, err := message.Marshal(appendEntriesRequest)
			if err != nil {
				node.log.
					Err(err).
					Str("Node", node.PersistentState.SelfID.String()).
					Msg("error")
				return
			}
			err = node.PersistentState.PeerIPs[i].Send(ctx, payload)
			if err != nil {
				node.log.
					Err(err).
					Str("Node", node.PersistentState.SelfID.String()).
					Msg("error")
				return
			}

			node.log.
				Debug().
				Str("self-id", node.PersistentState.SelfID.String()).
				Str("sent to", node.PersistentState.PeerIPs[i].RemoteID().String()).
				Msg("sent heartbeat to peer")

			res, err := node.PersistentState.PeerIPs[i].Receive(ctx)
			if err != nil {
				node.log.
					Err(err).
					Str("Node", node.PersistentState.SelfID.String()).
					Msg("error")
				return
			}

			resP, err := message.Unmarshal(res)
			if err != nil {
				node.log.
					Err(err).
					Str("Node", node.PersistentState.SelfID.String()).
					Msg("error")
				return
			}

			appendEntriesResponse := resP.(*message.AppendEntriesResponse)

			// If the term in the other node is greater than this node's term,
			// it means that this node is not up to date and has to step down
			// from being a leader.
			if appendEntriesResponse.Term > savedCurrentTerm {
				node.log.Debug().
					Str(node.PersistentState.SelfID.String(), "stale term").
					Str("following newer node", node.PersistentState.PeerIPs[i].RemoteID()) // TODO
				becomeFollower(node)
				return
			}

			node.PersistentState.mu.Lock()

			if node.State == StateLeader.String() && appendEntriesResponse.Term == savedCurrentTerm {
				if appendEntriesResponse.Success {
					node.VolatileStateLeader.NextIndex[i] = nextIndex + len(entries)
				} else {
					// If this appendEntries request failed,
					// proceed and retry in the next cycle.
					node.log.
					Debug().
					Str("self-id",node.PersistentState.SelfID.String()).
					Str("received failure to append entries from",node.PersistentState.PeerIPs[i].RemoteID()).
					Msg("failed to append entries")
				}
			}
		}(i)
	}
}
