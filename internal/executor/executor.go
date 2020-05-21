package executor

import (
	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor/command"
)

// Executor describes a component that can execute a command. A command is the
// intermediate representation of an SQL statement, meaning that it has been
// parsed.
type Executor interface {
	// Execute executes a command. The result of the computation is returned
	// together with an error, if one occurred.
	Execute(command.Command) (Result, error)
}

// New creates a new, ready to use Executor.
func New(log zerolog.Logger, databaseFile string) Executor {
	return newSimpleExecutor(log, databaseFile)
}
