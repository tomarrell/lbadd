package worker

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor"
)

type Worker struct {
	log  zerolog.Logger
	exec executor.Executor
}

func New(log zerolog.Logger, exec executor.Executor) *Worker {
	return &Worker{
		log:  log,
		exec: exec,
	}
}

func (w *Worker) Connect(ctx context.Context, addr string) error {
	w.log.Info().
		Str("addr", addr).
		Msg("connect")
	return fmt.Errorf("unimplemented")
}
