package node

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

// Node is a database node.
//
//  m := node.New(log)
//  err := m.ListenAndServe(ctx, ":34213")
type Node struct {
	log zerolog.Logger
}

// New creates a new node that is executing commands on the given executor.
func New(log zerolog.Logger) *Node {
	return &Node{
		log: log,
	}
}

// ListenAndServe starts the node on the given address. The given context must
// be used to stop the server, since there is no stop function. Canceling the
// context or a context timeout will cause the server to attempt a graceful
// shutdown.
func (m *Node) ListenAndServe(ctx context.Context, addr string) error {
	m.log.Info().
		Str("addr", addr).
		Msg("listen and serve")
	return fmt.Errorf("unimplemented")
}
