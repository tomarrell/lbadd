package network

import (
	"context"
	"io"
	"net"

	"github.com/tomarrell/lbadd/internal/id"
)

// ConnHandler is a handler function for handling new connections. It will be
// called with a fully initialized Conn, and is used as a callback in the
// server.
type ConnHandler func(Conn)

// Server describes a server component, that listens for connecting clients.
// Before opening, it is recommended that one sets a connect handler with
// Server.OnConnect. A server can only be opened once. Closing a server must not
// close the accepted connections, but must only stop accepting new connections
// and release the allocated address.
type Server interface {
	io.Closer

	// Open opens the server on the given address. To make the server choose a
	// random free port for you, specify a port ":0".
	Open(string) error
	// Listening can be used to get a signal when the server has allocated a
	// port and is now actively listening for incoming connections.
	Listening() <-chan struct{}
	// Addr returns the address that this server is listening to.
	Addr() net.Addr

	// OwnID returns the ID of this server. The remote ID of any connection is
	// the own ID of another server.
	OwnID() id.ID
	// OnConnect sets a callback that will be executed whenever a new connection
	// connects to this server.
	OnConnect(ConnHandler)
}

//go:generate mockery -case=snake -name=Conn

// Conn describes a network connection. One can send a message with Conn.Send,
// and receive one with Conn.Receive. Unlike an io.Writer, the data that is
// passed into Send is guaranteed to be returned in a single Receive call on the
// other end, meaning that you don't have to worry about where your messages
// end. Maximum message length is 2GiB.
type Conn interface {
	io.Closer

	// RemoteID returns the own ID of the server that this connection points to.
	RemoteID() id.ID
	// Send sends the given payload to the remote part of this connection. The
	// message will not be chunked, and can be read with a single call to
	// Conn.Receive.
	Send(context.Context, []byte) error
	// Receive reads a whole message and returns it in a byte slice. A message
	// is a byte slice that was sent with a single call to Conn.Send.
	Receive(context.Context) ([]byte, error)
}
