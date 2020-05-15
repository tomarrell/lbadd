package raft

import (
	"context"
	"fmt"
	"net"
	"sync"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/cluster"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

func Test_LeaderElection(t *testing.T) {
	assert := assert.New(t)

	ctx := context.TODO()
	log := zerolog.Nop()
	cluster := cluster.NewTCPCluster(log)
	// server := NewServer(
	// 	log,
	// 	cluster,
	// )

	conn1, conn2 := net.Pipe()
	tcp1int, tcp1ext := network.NewTCPConn(conn1), network.NewTCPConn(conn2)
	tcp2int, tcp2ext := network.NewTCPConn(conn1), network.NewTCPConn(conn2)
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
		fmt.Printf("Entered first gofunc\n")
		res, err := tcp1ext.Receive(ctx)
		if err != nil {
			fmt.Printf("Error01:%v\n", err)
		}
		// msg, err := message.Unmarshal(res)
		// _ = msg
		_ = res
		resP := message.NewRequestVoteResponse(1, true)

		payload, err := message.Marshal(resP)
		if err != nil {
			fmt.Printf("Error1:%v\n", err)
		}
		err = tcp1ext.Send(ctx, payload)
		if err != nil {
			fmt.Printf("Error11:%v\n", err)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Printf("Entered second gofunc\n")
		res, err := tcp2ext.Receive(ctx)
		if err != nil {
			fmt.Printf("Error02:%v\n", err)
		}
		// msg, err := message.Unmarshal(res)
		// _ = msg
		_ = res
		resP := message.NewRequestVoteResponse(1, true)

		payload, err := message.Marshal(resP)
		if err != nil {
			fmt.Printf("Error2:%v\n", err)
		}
		err = tcp2ext.Send(ctx, payload)
		if err != nil {
			fmt.Printf("Error21:%v\n", err)
		}
		wg.Done()
	}()

	err := StartElection(node)

	wg.Wait()

	assert.NoError(err)
}
