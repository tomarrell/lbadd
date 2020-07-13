package raft

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	networkmocks "github.com/tomarrell/lbadd/internal/network/mocks"
	"github.com/tomarrell/lbadd/internal/raft/message"
	raftmocks "github.com/tomarrell/lbadd/internal/raft/mocks"
)

// Test_Raft tests the entire raft operation.
func Test_Raft(t *testing.T) {

	assert := assert.New(t)
	ctx := context.Background()

	log := zerolog.New(os.Stdout).With().Logger().Level(zerolog.GlobalLevel())

	// Create a new cluster.
	cluster := new(raftmocks.Cluster)
	clusterID := id.Create()

	// Mock 4 other nodes in the cluster.
	conn1 := new(networkmocks.Conn)
	conn2 := new(networkmocks.Conn)
	conn3 := new(networkmocks.Conn)
	conn4 := new(networkmocks.Conn)

	connSlice := []network.Conn{
		conn1,
		conn2,
		conn3,
		conn4,
	}

	conn1 = addRemoteID(conn1)
	conn2 = addRemoteID(conn2)
	conn3 = addRemoteID(conn3)
	conn4 = addRemoteID(conn4)

	conn1.On("Send", ctx, mock.IsType([]byte{})).Return(nil)
	conn2.On("Send", ctx, mock.IsType([]byte{})).Return(nil)
	conn3.On("Send", ctx, mock.IsType([]byte{})).Return(nil)
	conn4.On("Send", ctx, mock.IsType([]byte{})).Return(nil)

	reqVRes1 := message.NewRequestVoteResponse(1, true)
	payload1, err := message.Marshal(reqVRes1)
	assert.NoError(err)

	conn1.On("Receive", ctx).Return(payload1, nil).Once()
	conn2.On("Receive", ctx).Return(payload1, nil).Once()
	conn3.On("Receive", ctx).Return(payload1, nil).Once()
	conn4.On("Receive", ctx).Return(payload1, nil).Once()

	appERes1 := message.NewAppendEntriesResponse(1, true)
	payload2, err := message.Marshal(appERes1)
	assert.NoError(err)

	conn1.On("Receive", ctx).Return(payload2, nil)
	conn2.On("Receive", ctx).Return(payload2, nil)
	conn3.On("Receive", ctx).Return(payload2, nil)
	conn4.On("Receive", ctx).Return(payload2, nil)

	// set up cluster to return the slice of connections on demand.
	cluster.
		On("Nodes").
		Return(connSlice)

	// return cluster ID
	cluster.
		On("OwnID").
		Return(clusterID)

	cluster.
		On("Receive", ctx).
		Return(conn1, nil, nil).After(time.Duration(1000) * time.Second)

	cluster.On("Close").Return(nil)

	server := newServer(
		log,
		cluster,
		timeoutProvider,
	)

	server.OnRequestVotes(func(msg *message.RequestVoteRequest) {})
	server.OnLeaderElected(func() {})
	server.OnAppendEntries(func() {
		err = server.Close()
		if err != network.ErrClosed {
			assert.NoError(err)
		}
	})
	err = server.Start()
	assert.NoError(err)
}

func addRemoteID(conn *networkmocks.Conn) *networkmocks.Conn {
	cID := id.Create()
	conn.On("RemoteID").Return(cID)
	return conn
}

func timeoutProvider(node *Node) *time.Timer {
	node.log.
		Debug().
		Str("self-id", node.PersistentState.SelfID.String()).
		Int("random timer set to", 150).
		Msg("heart beat timer")
	return time.NewTimer(time.Duration(150) * time.Millisecond)
}

func Test_Integration(t *testing.T) {
	log := zerolog.New(os.Stdout).With().Logger().Level(zerolog.GlobalLevel())

	assert := assert.New(t)
	operations := []OpData{
		{
			Op: SendData,
			Data: &OpSendData{
				Data: []*command.Command{},
			},
		},
		{
			Op:   StopNode,
			Data: &OpStopNode{},
		},
	}
	opParams := OperationParameters{
		Rounds:     4,
		TimeLimit:  5,
		Operations: operations,
	}

	cfg := NetworkConfiguration{}

	raftTest := NewSimpleRaftTest(log, opParams, cfg)

	go func() {
		err := raftTest.BeginTest()
		assert.Nil(err)
	}()

	<-time.After(time.Duration(2*opParams.TimeLimit) * time.Second)
}
