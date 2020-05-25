package raft

import (
	"net"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// TestAppendEntries function checks the correctnes of AppendEntriesResponse
// function. In this test function, we check how the function will respond to
// if node Term is less than leader node, node Log Index is less than leader
// commitIndex and checks if logs are appended correctly to node Log.
func TestAppendEntries(t *testing.T) {
	assert := assert.New(t)

	log := zerolog.Nop()
	cluster := cluster.NewTCPCluster(log)

	conn1, conn2 := net.Pipe()
	conn3, conn4 := net.Pipe()
	tcp1int, tcp1ext := network.NewTCPConn(conn1), network.NewTCPConn(conn2)
	tcp2int, tcp2ext := network.NewTCPConn(conn3), network.NewTCPConn(conn4)
	defer func() {
		_ = tcp1int.Close()
		_ = tcp1ext.Close()
		_ = tcp2int.Close()
		_ = tcp2ext.Close()
	}()
	cluster.AddConnection(tcp1int)
	cluster.AddConnection(tcp2int)

	// Created a mock node with default values for PersistentState
	// and Volatile State.
	// For Volatile State, CommitIndex given -1 to show no commit is
	// applied as Log Index start with 0, so -1 show value before any
	// log committed and same logic goes with LastApplied as -1 show
	// no logs with given Index is applied to State Machine.
	node := &Node{
		State: StateFollower.String(),
		PersistentState: &PersistentState{
			CurrentTerm: 0,
			VotedFor:    nil,
			SelfID:      cluster.OwnID(),
			PeerIPs:     cluster.Nodes(),
		},
		VolatileState: &VolatileState{
			CommitIndex: -1,
			LastApplied: -1,
		},
	}

	entries := []*message.LogData{
		message.NewLogData(2, "execute cmd3"),
		message.NewLogData(2, "execute cmd4"),
	}
	// Creating a mock msg AppendEntriesRequest with default values
	// Leader commit specifies the Index of Log commited by leader and
	// entries include msg LogData sent to nodes.
	msg := &message.AppendEntriesRequest{
		Term:         1,
		PrevLogIndex: -1,
		PrevLogTerm:  1,
		Entries:      entries,
		LeaderCommit: 3,
	}

	node.PersistentState.CurrentTerm = 3
	res := AppendEntriesResponse(node, msg)
	assert.False(res.Success, "Node Term is lesser than leader term")
	msg.Term = 3
	msg.PrevLogIndex = 3
	node.VolatileState.CommitIndex = 2
	res = AppendEntriesResponse(node, msg)
	assert.False(res.Success, "Node Log Index is lesser than"+
		"leader commit Index")
	msg.Term = 2
	node.PersistentState.CurrentTerm = 2
	msg.PrevLogIndex = 1
	msg.PrevLogTerm = 1
	node.VolatileState.CommitIndex = 1
	node.PersistentState.Log = []*message.LogData{message.NewLogData(1,
		"execute cmd1"), message.NewLogData(1, "execute cmd2")}
	numberOfPersistentLog := len(node.PersistentState.Log)
	res = AppendEntriesResponse(node, msg)
	assert.True(res.Success, "Msg isn't appended to the node Logs")
	assert.Equal(node.PersistentState.CurrentTerm, res.GetTerm(),
		"Node doesn't have same term as leader")
	assert.Equal(len(node.PersistentState.Log),
		numberOfPersistentLog+len(entries), "LogData hasn't appended to the node ")
}
