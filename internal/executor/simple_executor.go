package executor

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor/command"
)

var _ Executor = (*simpleExecutor)(nil)

type simpleExecutor struct {
	log          zerolog.Logger
	databaseFile string
}

func newSimpleExecutor(log zerolog.Logger, databaseFile string) *simpleExecutor {
	return &simpleExecutor{
		log:          log,
		databaseFile: databaseFile,
	}
}

func (e *simpleExecutor) Execute(cmd command.Command) (Result, error) {
	return nil, fmt.Errorf("unimplemented")
}
