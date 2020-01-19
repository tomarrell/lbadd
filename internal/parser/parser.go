package parser

import "github.com/tomarrell/lbadd/internal/parser/ast"

type Parser interface {
	HasNext() bool
	Next() *ast.SqlStmt
}
