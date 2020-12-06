package cluster

import (
	"context"
	"io"

	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// ConnHandler is a function that handles a connection and performs a
// handshake. If an error occurs, considering closing the connection. If you
// want the connection to be remembered by the cluster as node, you must add it
// with (cluster.Cluster).AddConnection(network.Conn).
type ConnHandler func(Cluster, network.Conn)

// Cluster describes a raft cluster. It sometimes has a leader and consists of
// nodes.
type Cluster interface {
	// Nodes returns all nodes in the cluster (except this one), including the
	// leader node.
	Nodes() []network.Conn
	// Receive blocks until any connection in the cluster has sent a message to
	// this node. It will return the connection and the message, with respect to
	// the given context.
	Receive(context.Context) (network.Conn, message.Message, error)
	// Broadcast sends the given message to all other nodes in this cluster,
	// with respect to the given context.
	Broadcast(context.Context, message.Message) error

	// OwnID returns the global ID of this node.
	OwnID() id.ID
	// Join joins the cluster at the given address. The given address may be the
	// address and port of any of the nodes in the existing cluster.
	Join(context.Context, string) error
	// Open creates a new cluster and opens it on the given address. This
	// creates a server that will listen for incoming connections.
	Open(context.Context, string) error
	// AddConnection adds the connection to the cluster. It is considered
	// another node in the cluster.
	AddConnection(network.Conn)
	// RemoveConnection closes the connection and removes it from the cluster.
	RemoveConnection(network.Conn)
	// OnConnect allows to set a connection hook. This is useful when
	// implementing a custom handshake for connecting to the cluster. By default
	// on connect will just remember the connection as cluster node. When this
	// is set explicitely, (cluster.Cluster).AddConnection(network.Conn) must be
	// called, otherwise the connection will not be added to the cluster.
	OnConnect(ConnHandler)

	io.Closer
}
