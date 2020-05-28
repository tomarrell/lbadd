package raft

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	networkmocks "github.com/tomarrell/lbadd/internal/network/mocks"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
	raftmocks "github.com/tomarrell/lbadd/internal/raft/mocks"
)

// Raft integration tests go here.
func Test_NewServer(t *testing.T) {
	t.SkipNow()
	assert := assert.New(t)

	log := zerolog.Nop()
	ctx := context.Background()
	cluster := cluster.NewTCPCluster(log)
	err := cluster.Open(ctx, ":0")
	server := NewServer(
		log,
		cluster,
	)
	assert.NoError(err)
	err = server.Start()
	assert.NoError(err)
}

// Test_Raft tests the entire raft operation.
func Test_Raft(t *testing.T) {
	t.SkipNow()
	zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Logger()

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

	conn1.On("Receive", ctx).Return(payload1, nil).Once()
	conn2.On("Receive", ctx).Return(payload1, nil).Once()
	conn3.On("Receive", ctx).Return(payload1, nil).Once()
	conn4.On("Receive", ctx).Return(payload1, nil).Once()

	appERes1 := message.NewAppendEntriesResponse(1, true)
	payload2, err := message.Marshal(appERes1)

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

	server := NewServer(
		log,
		cluster,
	)

	_ = server
	err = server.Start()
	<-time.NewTimer(time.Duration(1000) * time.Second).C
	assert.NoError(err)

	// msg1 := message.NewAppendEntriesResponse(12, true)
	// msg2 := message.NewAppendEntriesResponse(12, true)
	// // instead of mocking this connection, you can also use a real connection if
	// // you need
	// conn := new(networkmocks.Conn)
	// conn.
	// 	On("Send", mock.IsType(ctx), mock.IsType([]byte{})).
	// 	Return(nil)
	// // cluster := new(raftmocks.Cluster)
	// cluster.
	// 	On("Receive", mock.Anything).
	// 	Return(conn, msg1, nil).
	// 	Once()
	// cluster.
	// 	On("Receive", mock.Anything).
	// 	Return(conn, msg2, nil).
	// 	Once()
	// cluster.
	// 	On("Broadcast", mock.IsType(ctx), mock.IsType(msg1)).
	// 	Return(nil)
	// err := cluster.Broadcast(ctx, msg1)
	// assert.NoError(err)
	// cluster.AssertNumberOfCalls(t, "Broadcast", 1)
	// cluster.AssertCalled(t, "Broadcast", ctx, msg1)
}

// func addRemoteID(conn *networkmocks.Conn) *networkmocks.Conn {
// 	cID := id.Create()
// 	conn.On("RemoteID").Return(cID)
// 	return conn
// }
