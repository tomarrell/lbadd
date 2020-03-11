package master

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor"
)

type Master struct {
	log  zerolog.Logger
	exec executor.Executor
}

func New(log zerolog.Logger, exec executor.Executor) *Master {
	return &Master{
		log:  log,
		exec: exec,
	}
}

func (m *Master) ListenAndServe(ctx context.Context, addr string) error {
	m.log.Info().
		Str("addr", addr).
		Msg("listen and serve")
	return fmt.Errorf("unimplemented")
}
