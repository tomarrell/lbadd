package compile

import "github.com/tomarrell/lbadd/internal/parser/ast"

type Compiler interface {
	Compile(*ast.SQLStmt) (Command, error)
}
