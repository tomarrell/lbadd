package network_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
)

// TestTCPServerHandshake ensures that the server logon handshake with DialTCP
// works correctly. The handshake is responsible for sending the client the
// server ID, and then receive the client ID and remember it in its connection.
// After the handshake, the ID of the connection on the server side must be
// equal to the client ID, and the remote ID of the connection created with
// DialTCP must be equal to the server ID.
func TestTCPServerHandshake(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// create the server
	server := network.NewTCPServer(zerolog.Nop())
	serverID := server.OwnID()
	assert.NotNil(serverID)
	var serverConnsLock sync.Mutex
	var serverConns []network.Conn
	server.OnConnect(func(conn network.Conn) {
		serverConnsLock.Lock()
		defer serverConnsLock.Unlock()
		serverConns = append(serverConns, conn)
	})

	// open the server in separate goroutine
	go func() {
		err := server.Open(":0")
		assert.NoError(err)
	}()

	// enforce timeout for server open
	select {
	case <-ctx.Done():
		_ = server.Close()
		t.Error("timeout")
	case <-server.Listening():
	}

	// dial the server
	conn1ID := id.Create() // create a connection ID
	conn1, err := network.DialTCP(ctx, conn1ID, server.Addr().String())
	assert.NoError(err)

	// check the client side connection
	assert.Equal(serverID, conn1.RemoteID()) // ensure that the remote ID of this connection is equal to the own ID of the server

	// check the server side connections
	serverConnsLock.Lock()
	assert.Len(serverConns, 1)
	assert.Equal(conn1ID, serverConns[0].RemoteID())
	serverConnsLock.Unlock()
}
