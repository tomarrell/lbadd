package master

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor"
)

// Master is a database master node.
//
//  m := master.New(log, executor)
//  err := m.ListenAndServe(ctx, ":34213")
type Master struct {
	log  zerolog.Logger
	exec executor.Executor
}

// New creates a new master node that is executing commands on the given
// executor.
func New(log zerolog.Logger, exec executor.Executor) *Master {
	return &Master{
		log:  log,
		exec: exec,
	}
}

// ListenAndServe starts the master node on the given address. The given context
// must be used to stop the server, since there is no stop function. Canceling
// the context or a context timeout will cause the server to attempt a graceful
// shutdown.
func (m *Master) ListenAndServe(ctx context.Context, addr string) error {
	m.log.Info().
		Str("addr", addr).
		Msg("listen and serve")
	return fmt.Errorf("unimplemented")
}
