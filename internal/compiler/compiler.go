package compiler

import (
	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/parser/ast"
)

// Compiler describes a component that is able to convert an (*ast.SQLStmt) to a
// (command.Command).
type Compiler interface {
	// Compile compiles an SQLStmt to a command, which can be used by an
	// executor. If an error occurs during the compiling, it will be returned.
	// If such an error is not fatal to the compile process and there occur more
	// errors, the returned error will contain all occurred errors until a fatal
	// error.
	Compile(*ast.SQLStmt) (command.Command, error)
}
