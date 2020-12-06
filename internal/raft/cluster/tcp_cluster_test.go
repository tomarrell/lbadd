package cluster_test

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

func TestTCPClusterCommunication(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cluster := cluster.NewTCPCluster(zerolog.Nop())
	defer func() {
		_ = cluster.Close()
	}()
	assert.Empty(cluster.Nodes())

	err := cluster.Open(ctx, ":0")
	assert.NoError(err)

	conn1, conn2 := net.Pipe()
	tcp1, tcp2 := network.NewTCPConn(conn1), network.NewTCPConn(conn2)
	defer func() {
		_ = tcp1.Close()
		_ = tcp2.Close()
	}()

	cluster.AddConnection(tcp1)
	assert.Len(cluster.Nodes(), 1)

	t.Run("Broadcast", _TestTCPClusterBroadcast(ctx, cluster, tcp2))
	t.Run("Receive", _TestTCPClusterReceive(ctx, cluster, tcp1, tcp2))
}

func _TestTCPClusterBroadcast(ctx context.Context, cluster cluster.Cluster, externalConn network.Conn) func(*testing.T) {
	return func(t *testing.T) {
		assert := assert.New(t)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			err := cluster.Broadcast(ctx, message.NewTestMessage("Hello, World!"))
			assert.NoError(err)
			wg.Done()
		}()

		data, err := externalConn.Receive(ctx)
		assert.NoError(err)

		wg.Wait()
		msg, err := message.Unmarshal(data)
		assert.NoError(err)
		assert.Equal(message.KindTestMessage, msg.Kind())
		assert.IsType(&message.TestMessage{}, msg)
		assert.Equal("Hello, World!", msg.(*message.TestMessage).GetData())
	}
}

func _TestTCPClusterReceive(ctx context.Context, cluster cluster.Cluster, internalConn, externalConn network.Conn) func(*testing.T) {
	return func(t *testing.T) {
		assert := assert.New(t)

		data, err := message.Marshal(message.NewTestMessage("Hello, World!"))
		assert.NoError(err)

		// This should not block, since the cluster uses a buffered message
		// queue. If it blocks however, it will run into a timeout.
		err = externalConn.Send(ctx, data)
		assert.NoError(err)

		conn, msg, err := cluster.Receive(ctx)
		assert.NoError(err)
		// The external conn ID is not equal to the conn.ID(), because it did
		// not connect to a network.Server with network.DialTCP, and thus had no
		// chance of exchanging the ID. When connecting to the cluster with
		// cluster.Join or network.DialTCP however, this ID will be the same.
		assert.Equal(internalConn.RemoteID(), conn.RemoteID())
		assert.Equal(message.KindTestMessage, msg.Kind())
		assert.IsType(&message.TestMessage{}, msg)
		assert.Equal("Hello, World!", msg.(*message.TestMessage).GetData())
	}
}
