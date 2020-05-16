package raft

import (
	"context"

	"github.com/tomarrell/lbadd/internal/id"
	"github.com/tomarrell/lbadd/internal/network"
	"github.com/tomarrell/lbadd/internal/raft/message"
)

//go:generate mockery -case=snake -name=Cluster

// Cluster is a description of a cluster of servers.
type Cluster interface {
	OwnID() id.ID
	Nodes() []network.Conn
	Receive(context.Context) (network.Conn, message.Message, error)
	Broadcast(context.Context, message.Message) error
}
