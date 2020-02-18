package executor

import "github.com/tomarrell/lbadd/internal/executor/command"

type Executor interface {
	Execute(command.Command) error
}

func New() Executor {
	return newSimpleExecutor()
}
