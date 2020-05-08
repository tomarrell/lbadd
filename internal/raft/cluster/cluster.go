package cluster

import (
	"context"

	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

// Cluster describes a raft cluster. It sometimes has a leader and consists of
// nodes.
type Cluster interface {
	// Leader returns the current cluster leader, or nil if no leader has
	// elected or this node is the leader.
	Leader() network.Conn
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
}
