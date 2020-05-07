package network

import (
	"net"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPConnSendReceive(t *testing.T) {
	assert := assert.New(t)
	conn1, conn2 := net.Pipe()
	tcpConn1, tcpConn2 := newTCPConn(conn1), newTCPConn(conn2)

	payload := []byte("Hello, World!")
	recv := make([]byte, len(payload))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var err error
		recv, err = tcpConn2.Receive()
		assert.NoError(err)
		wg.Done()
	}()

	err := tcpConn1.Send(payload)
	assert.NoError(err)

	wg.Wait()
	assert.Equal(payload, recv)
}

func TestDialTCP(t *testing.T) {
	assert := assert.New(t)
	payload := []byte("Hello, World!")

	lis, err := net.Listen("tcp", ":0")
	assert.NoError(err)

	var srvConnID string
	go func() {
		conn, err := lis.Accept()
		assert.NoError(err)

		tcpConn := newTCPConn(conn)
		srvConnID = tcpConn.ID().String()
		assert.NoError(tcpConn.Send(tcpConn.ID().Bytes()))
		assert.NoError(tcpConn.Send(payload))
	}()

	port := lis.Addr().(*net.TCPAddr).Port

	conn, err := DialTCP(":" + strconv.Itoa(port))
	assert.NoError(err)
	defer func() { assert.NoError(conn.Close()) }()
	assert.Equal(srvConnID, conn.ID().String())

	recv, err := conn.Receive()
	assert.NoError(err)
	assert.Equal(payload, recv)
}
