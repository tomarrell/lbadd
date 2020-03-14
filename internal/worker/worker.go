package worker

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor"
)

// Worker is a database worker node.
//
//  w := worker.New(log, exec)
//  err := w.Connect(ctx, ":34213")
type Worker struct {
	log  zerolog.Logger
	exec executor.Executor
}

// New creates a new worker with the given logger, that will replicate master
// commands with the given executor.
func New(log zerolog.Logger, exec executor.Executor) *Worker {
	return &Worker{
		log:  log,
		exec: exec,
	}
}

// Connect connects the worker to a running and available master node. Use the
// given context to close the connection and terminate the worker node.
// Canceling the context will cause the worker to attempt a graceful shutdown.
func (w *Worker) Connect(ctx context.Context, addr string) error {
	w.log.Info().
		Str("addr", addr).
		Msg("connect")
	return fmt.Errorf("unimplemented")
}
