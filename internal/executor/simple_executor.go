package executor

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor/command"
)

var _ Executor = (*simpleExecutor)(nil)

type simpleExecutor struct {
	log zerolog.Logger
}

func newSimpleExecutor(log zerolog.Logger) *simpleExecutor {
	return &simpleExecutor{
		log: log,
	}
}

func (e *simpleExecutor) Execute(cmd command.Command) (Result, error) {
	return nil, fmt.Errorf("unimplemented")
}
