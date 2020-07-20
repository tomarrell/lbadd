package raft

import (
	"context"
	"net"
	"os"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
	"gotest.tools/assert/cmp"
)

func Test_LeaderElection(t *testing.T) {
	t.SkipNow()
	assert := assert.New(t)

	zerolog.New(os.Stdout).With().
		Str("foo", "bar").
		Logger()

	ctx := context.TODO()
	log := zerolog.New(os.Stdout).With().Logger().Level(zerolog.GlobalLevel())
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

	node := NewRaftNode(cluster)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		res, err := tcp1ext.Receive(ctx)
		assert.Nil(err)

		msg, err := message.Unmarshal(res)
		assert.Nil(err)
		_ = msg
		_ = res
		resP := message.NewRequestVoteResponse(1, true)

		payload, err := message.Marshal(resP)
		assert.Nil(err)

		err = tcp1ext.Send(ctx, payload)
		assert.Nil(err)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		res, err := tcp2ext.Receive(ctx)
		assert.Nil(err)

		msg, err := message.Unmarshal(res)
		assert.Nil(err)
		_ = msg
		_ = res
		resP := message.NewRequestVoteResponse(1, true)

		payload, err := message.Marshal(resP)
		assert.Nil(err)
		err = tcp2ext.Send(ctx, payload)
		assert.Nil(err)
		wg.Done()
	}()

	server := SimpleServer{
		node:    node,
		cluster: cluster,
		log:     log,
	}

	server.StartElection(ctx)

	wg.Wait()

	node.PersistentState.mu.Lock()
	assert.True(cmp.Equal(node.PersistentState.SelfID, node.PersistentState.LeaderID)().Success())
	node.PersistentState.mu.Unlock()
}
