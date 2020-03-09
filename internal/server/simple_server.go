package server

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor"
)

var _ Server = (*simpleServer)(nil)

type simpleServer struct {
	log zerolog.Logger

	executor executor.Executor
}

func newSimpleServer(log zerolog.Logger, executor executor.Executor) *simpleServer {
	return &simpleServer{
		log:      log,
		executor: executor,
	}
}

func (s *simpleServer) ListenAndServe(ctx context.Context, addr string) error {
	s.log.Info().Str("addr", addr).Msg("start listening")
	return fmt.Errorf("unimplemented")
}
