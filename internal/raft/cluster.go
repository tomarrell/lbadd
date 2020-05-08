package raft

import (
	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/network"
)

// Cluster describes a raft cluster
type Cluster interface {
	Leader() network.Conn
	Nodes() []network.Conn
	Log() zerolog.Logger
}
