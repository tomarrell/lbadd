package compile

import "github.com/tomarrell/lbadd/internal/parser/ast"

// Compiler describes a component that can convert an SQL statement to the
// intermediary representation, called the IR, which is represented by a
// command. An executor will be able to execute such a command.
type Compiler interface {
	// Compile converts the statement to a command, or returns an error, if it
	// is unable to do so.
	Compile(*ast.SQLStmt) (Command, error)
}
