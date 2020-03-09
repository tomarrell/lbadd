package server

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor"
)

// Server describes a component that can be started on a network address with a
// context, that is used to terminate the server, instead of a Close method.
type Server interface {
	ListenAndServe(context.Context, string) error
}

// New creates a new server that uses the given logger and executes statements
// with the help of the given executor.
func New(log zerolog.Logger, executor executor.Executor) Server {
	return newSimpleServer(log, executor)
}
