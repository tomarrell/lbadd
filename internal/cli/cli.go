package cli

import (
	"io"

	"github.com/tomarrell/lbadd/internal/parser/ast"
)

type Cli interface {
	Start()
	io.Closer
}

type Executor interface {
	Execute(*ast.SQLStmt) error
}

func New(in io.Reader, out io.Writer, exec Executor) Cli {
	return newSimpleCli(in, out, exec)
}
