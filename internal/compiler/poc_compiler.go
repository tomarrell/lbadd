package compiler

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/parser/ast"
)

var _ Compiler = (*pocCompiler)(nil)

type pocCompiler struct {
}

func (c *pocCompiler) Compile(stmt *ast.SQLStmt) (command.Command, error) {
	if err := c.validate(stmt); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	cmd, err := c.compileSelect(stmt.SelectStmt)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	return cmd, nil
}

func (c *pocCompiler) validate(stmt *ast.SQLStmt) error {
	if !(stmt.Explain == nil && stmt.Query == nil && stmt.Plan == nil) {
		return fmt.Errorf("can't explain: %w", ErrUnsupported)
	}

	// poc compiler only supports compiling basic select statements
	if stmt.SelectStmt == nil {
		return fmt.Errorf("only select is supported by the poc compiler: %w", ErrUnsupported)
	}

	return nil
}

func (c *pocCompiler) compileSelect(stmt *ast.SelectStmt) (command.Command, error) {
	if stmt.WithClause != nil {
		return nil, fmt.Errorf("with: %w", ErrUnsupported)
	}
	if len(stmt.SelectCore) != 1 {
		return nil, fmt.Errorf("more than one core: %w", ErrUnsupported)
	}

	return c.compileSelectCore(stmt.SelectCore[0])
}

func (c *pocCompiler) compileSelectCore(core *ast.SelectCore) (command.Command, error) {
	if core.Distinct != nil {
		return nil, fmt.Errorf("distinct: %w", ErrUnsupported)
	}
	if core.All != nil {
		return nil, fmt.Errorf("all: %w", ErrUnsupported)
	}
	if len(core.ResultColumn) != 1 {
		return nil, fmt.Errorf("not single col: %w", ErrUnsupported)
	}
	// must be disabled because of #127
	// if core.JoinClause != nil {
	// 	return nil, fmt.Errorf("join: %w", ErrUnsupported)
	// }
	if core.JoinClause.JoinClausePart != nil { // workaround for #127
		return nil, fmt.Errorf("join: %w", ErrUnsupported)
	}
	if core.Group != nil {
		return nil, fmt.Errorf("group: %w", ErrUnsupported)
	}
	if core.Window != nil {
		return nil, fmt.Errorf("window: %w", ErrUnsupported)
	}
	if core.Values != nil {
		return nil, fmt.Errorf("values: %w", ErrUnsupported)
	}
	// must be disables because of #127
	// if len(core.TableOrSubquery) != 1 {
	// 	return nil, fmt.Errorf("not single table: %w", ErrUnsupported)
	// }
	var tableName string
	if len(core.TableOrSubquery) == 1 {
		tableName = core.TableOrSubquery[0].TableName.Value()
	} else {
		tableName = core.JoinClause.TableOrSubquery.TableName.Value()
	}

	return command.Select{
		Tables: []string{tableName},
		Cols:   []string{core.ResultColumn[0].Expr.LiteralValue.Value()},
	}, nil
}
