package network

import (
	"context"
	"fmt"
	"io"
	"net"
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

	// OnConnect sets a callback that will be executed whenever a new connection
	// connects to this server.
	OnConnect(ConnHandler)
}

// Conn describes a network connection. One can send a message with Conn.Send,
// and receive one with Conn.Receive. Unlike an io.Writer, the data that is
// passed into Send is guaranteed to be returned in a single Receive call on the
// other end, meaning that you don't have to worry about where your messages
// end. Maximum message length is 2GiB.
type Conn interface {
	io.Closer

	// ID returns the ID of this connection. It can be used to uniquely identify
	// this connection globally.
	ID() ID
	// Send sends the given payload to the remote part of this connection. The
	// message will not be chunked, and can be read with a single call to
	// Conn.Receive.
	Send(context.Context, []byte) error
	// Receive reads a whole message and returns it in a byte slice. A message
	// is a byte slice that was sent with a single call to Conn.Send.
	Receive(context.Context) ([]byte, error)
}

// ID describes an identifier that is used for connections. An ID has to be
// unique application-wide. IDs must not be re-used.
type ID interface {
	fmt.Stringer
	Bytes() []byte
}
