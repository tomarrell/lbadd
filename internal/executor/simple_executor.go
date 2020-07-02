package executor

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/compiler/command"
)

var _ Executor = (*simpleExecutor)(nil)

type simpleExecutor struct {
	log          zerolog.Logger
	databaseFile string
}

// NewSimpleExecutor creates a new ready to use executor, that operates on the
// given database file.
func NewSimpleExecutor(log zerolog.Logger, databaseFile string) Executor {
	return &simpleExecutor{
		log:          log,
		databaseFile: databaseFile,
	}
}

func (e *simpleExecutor) Execute(cmd command.Command) (Result, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (e *simpleExecutor) Close() error {
	return nil
}
