package compile

import "github.com/tomarrell/lbadd/internal/parser/ast"

var _ Compiler = (*simpleCompiler)(nil)

type simpleCompiler struct{}

// NewSimpleCompiler creates a new, ready to use compiler.
func NewSimpleCompiler() Compiler {
	return &simpleCompiler{}
}

func (c *simpleCompiler) Compile(stmt *ast.SQLStmt) (Command, error) {
	return Command{}, nil
}
