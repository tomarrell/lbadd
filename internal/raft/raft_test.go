package raft

import (
	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
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
	assert := assert.New(t)
	ctx := context.Background()
	log := zerolog.Nop()

	// create a new cluster
	cluster := new(raftmocks.Cluster)

	conn1 := new(network.Conn)
	conn2 := new(network.Conn)
	conn3 := new(network.Conn)
	conn4 := new(network.Conn)
	conn5 := new(network.Conn)

	connSlice := []network.Conn{
		*conn1,
		*conn2,
		*conn3,
		*conn4,
		*conn5,
	}

	// set up cluster to return the slice of connections on demand.
	cluster.
		On("Nodes").
		Return(connSlice)

	clusterID := id.Create()

	// return cluster ID
	cluster.
		On("OwnID").
		Return(clusterID)

	cluster.
		On("Receive", mock.IsType(ctx)).
		Return(nil, nil, nil)
	server := NewServer(
		log,
		cluster,
	)

	err := server.Start()
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
