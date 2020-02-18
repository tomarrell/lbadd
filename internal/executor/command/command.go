package command

import "github.com/tomarrell/lbadd/internal/parser/ast"

type Command struct {
}

func From(stmt *ast.SQLStmt) (Command, error) {
	return Command{}, nil
}
