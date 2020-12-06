package network

import (
	"context"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/id"
)

func TestTCPConnSendReceive(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn1, conn2 := net.Pipe()
	tcpConn1, tcpConn2 := NewTCPConn(conn1), NewTCPConn(conn2)

	payload := []byte("Hello, World!")
	recv := make([]byte, len(payload))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var err error
		recv, err = tcpConn2.Receive(ctx)
		assert.NoError(err)
		wg.Done()
	}()

	err := tcpConn1.Send(ctx, payload)
	assert.NoError(err)

	wg.Wait()
	assert.Equal(payload, recv)
}

func TestDialTCP(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	payload := []byte("Hello, World!")

	lis, err := net.Listen("tcp", ":0")
	assert.NoError(err)

	clientID := id.Create()
	srvID := id.Create()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		conn, err := lis.Accept()
		assert.NoError(err)

		// default handshake
		tcpConn := NewTCPConn(conn)
		assert.NoError(tcpConn.Send(ctx, srvID.Bytes()))   // send server ID
		recvID, err := tcpConn.Receive(ctx)                // receive client ID
		assert.NoError(err)                                //
		parsedID, err := id.Parse(recvID)                  // parse client ID
		assert.NoError(err)                                //
		assert.Equal(clientID.String(), parsedID.String()) // parsed ID must be equal to actual client ID

		assert.NoError(tcpConn.Send(ctx, payload))

		wg.Done()
	}()

	port := lis.Addr().(*net.TCPAddr).Port

	conn, err := DialTCP(ctx, clientID, ":"+strconv.Itoa(port))
	assert.NoError(err)
	defer func() { assert.NoError(conn.Close()) }()
	assert.Equal(srvID.String(), conn.RemoteID().String())

	recv, err := conn.Receive(ctx)
	assert.NoError(err)
	assert.Equal(payload, recv)

	wg.Wait()
}

func TestTCPConnWriteContext(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	conn1, conn2 := net.Pipe()
	tcpConn1, _ := NewTCPConn(conn1), NewTCPConn(conn2)

	err := tcpConn1.Send(ctx, []byte("Hello")) // will not be able to write within 10ms, because noone is reading
	assert.Equal(ErrTimeout, err)
}

func TestTCPConnReadContext(t *testing.T) {
	assert := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	conn1, conn2 := net.Pipe()
	tcpConn1, _ := NewTCPConn(conn1), NewTCPConn(conn2)

	data, err := tcpConn1.Receive(ctx) // will not be able to receive within 10ms, because noone is writing
	assert.Equal(ErrTimeout, err)
	assert.Nil(data)
}
