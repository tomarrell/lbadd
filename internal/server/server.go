package server

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor"
)

type Server interface {
	ListenAndServe(context.Context, string) error
}

func New(log zerolog.Logger, executor executor.Executor) Server {
	return newSimpleServer(log, executor)
}
