package raft

import (
	"testing"
)

func Test_LeaderElection(t *testing.T) {
	// assert := assert.New(t)

	// zerolog.New(os.Stdout).With().
	// 	Str("foo", "bar").
	// 	Logger()

	// ctx := context.TODO()
	// log := zerolog.New(os.Stdout).With().Logger().Level(zerolog.GlobalLevel())
	// cluster := new(raftmocks.Cluster)
	// clusterID := id.Create()

	// conn1 := new(networkmocks.Conn)
	// conn2 := new(networkmocks.Conn)

	// connSlice := []network.Conn{
	// 	conn1,
	// 	conn2,
	// }

	// conn1 = addRemoteID(conn1)
	// conn2 = addRemoteID(conn2)

	// conn1.On("Send", ctx, mock.IsType([]byte{})).Return(nil)
	// conn2.On("Send", ctx, mock.IsType([]byte{})).Return(nil)

	// reqVRes1 := message.NewRequestVoteResponse(1, true)
	// payload1, err := message.Marshal(reqVRes1)
	// assert.Nil(err)

	// conn1.On("Receive", ctx).Return(payload1, nil).Once()
	// conn2.On("Receive", ctx).Return(payload1, nil).Once()

	// cluster.
	// 	On("Nodes").
	// 	Return(connSlice)

	// cluster.
	// 	On("OwnID").
	// 	Return(clusterID)

	// node := NewRaftNode(cluster)

	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	res, err := tcp1ext.Receive(ctx)
	// 	assert.Nil(err)

	// 	msg, err := message.Unmarshal(res)
	// 	assert.Nil(err)
	// 	_ = msg
	// 	_ = res
	// 	resP := message.NewRequestVoteResponse(1, true)

	// 	payload, err := message.Marshal(resP)
	// 	assert.Nil(err)

	// 	err = tcp1ext.Send(ctx, payload)
	// 	assert.Nil(err)
	// 	wg.Done()
	// }()

	// wg.Add(1)
	// go func() {
	// 	res, err := tcp2ext.Receive(ctx)
	// 	assert.Nil(err)

	// 	msg, err := message.Unmarshal(res)
	// 	assert.Nil(err)
	// 	_ = msg
	// 	_ = res
	// 	resP := message.NewRequestVoteResponse(1, true)

	// 	payload, err := message.Marshal(resP)
	// 	assert.Nil(err)
	// 	err = tcp2ext.Send(ctx, payload)
	// 	assert.Nil(err)
	// 	wg.Done()
	// }()

	// node.StartElection()

	// wg.Wait()

	// node.PersistentState.mu.Lock()
	// assert.True(cmp.Equal(node.PersistentState.SelfID, node.PersistentState.LeaderID))
	// node.PersistentState.mu.Unlock()
}
