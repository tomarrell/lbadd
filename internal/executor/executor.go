package executor

import (
	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor/command"
)

type Executor interface {
	Execute(command.Command) (Result, error)
}

func New(log zerolog.Logger) Executor {
	return newSimpleExecutor(log)
}
