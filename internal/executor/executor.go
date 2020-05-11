package executor

import (
	"io"

	"github.com/tomarrell/lbadd/internal/compile"
)

// Executor describes a component that can execute a command. A command is the
// intermediate representation of an SQL statement, meaning that it has been
// parsed.
type Executor interface {
	// Execute executes a command. The result of the computation is returned
	// together with an error, if one occurred.
	Execute(compile.Command) (Result, error)

	io.Closer
}
