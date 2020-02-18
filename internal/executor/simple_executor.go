package executor

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/executor/command"
)

var _ Executor = (*simpleExecutor)(nil)

type simpleExecutor struct {
}

func newSimpleExecutor() *simpleExecutor {
	return &simpleExecutor{}
}

func (e *simpleExecutor) Execute(cmd command.Command) error {
	return fmt.Errorf("unimplemented")
}
