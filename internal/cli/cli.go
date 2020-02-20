package cli

import (
	"context"
	"io"

	"github.com/tomarrell/lbadd/internal/executor"
)

// Cli describes a command line interface that can be started. A Cli runs under
// a context. Processing must start when the Cli is started and stopped, when
// the context is canceled.
type Cli interface {
	Start()
}

// New creates a new Cli that can immediately be started.
func New(ctx context.Context, in io.Reader, out io.Writer, exec executor.Executor) Cli {
	return newSimpleCli(ctx, in, out, exec)
}
