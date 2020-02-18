package cli

import (
	"io"

	"github.com/tomarrell/lbadd/internal/executor/command"
)

// Cli describes a command line interface that can be started and closed. It
// should process commands from a defined input as long as it is running.
// Processing must stop, when the cli is closed.
type Cli interface {
	Start()
	io.Closer
}

// Executor describes a component that can execute the AST that is produced when
// parsing the input.
type Executor interface {
	Execute(command.Command) error
}

// New creates a new Cli that can immediately be started.
func New(in io.Reader, out io.Writer, exec Executor) Cli {
	return newSimpleCli(in, out, exec)
}
