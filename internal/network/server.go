package network

import (
	"fmt"
	"io"
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

	Open(string) error
	OnConnect(ConnHandler)
}

// Conn describes a network connection. One can send a message with Conn.Send,
// and receive one with Conn.Receive. Unlike an io.Writer, the data that is
// passed into Send is guaranteed to be returned in a single Receive call on the
// other end, meaning that you don't have to worry about where your messages
// end. Maximum message length is 2GiB.
type Conn interface {
	io.Closer

	ID() ID
	Send([]byte) error
	Receive() ([]byte, error)
}

// ID describes an identifier that is used for connections. An ID has to be
// unique application-wide. IDs must not be re-used.
type ID interface {
	fmt.Stringer
}
