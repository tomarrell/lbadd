package compiler

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/compiler/optimization"
	"github.com/tomarrell/lbadd/internal/parser/ast"
)

type simpleCompiler struct {
	optimizations []optimization.Optimization
}

// OptionEnableOptimization is used to enable the given optimization in a
// compiler.
func OptionEnableOptimization(opt optimization.Optimization) Option {
	return func(c *simpleCompiler) {
		c.optimizations = append(c.optimizations, opt)
	}
}

// New creates a new, ready to use compiler with the given options applied.
func New(opts ...Option) Compiler {
	c := &simpleCompiler{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *simpleCompiler) Compile(ast *ast.SQLStmt) (command.Command, error) {
	// compile the ast
	cmd, err := c.compileInternal(ast)
	if err != nil {
		return nil, err
	}
	// apply optimizations
	for _, opt := range c.optimizations {
		if optimized, ok := opt(cmd); ok {
			cmd = optimized
		}
	}
	if ast.Explain != nil {
		return command.Explain{
			Command: cmd,
		}, nil
	}
	return cmd, nil
}

func (c *simpleCompiler) compileInternal(ast *ast.SQLStmt) (command.Command, error) {
	if ast.SelectStmt == nil {
		return nil, fmt.Errorf("not select: %w", ErrUnsupported)
	}
	cmd, err := c.compileSelect(ast.SelectStmt)
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}
	return cmd, nil
}

func (c *simpleCompiler) compileSelect(stmt *ast.SelectStmt) (command.Command, error) {
	// This implementation is incomplete, it is missing everything else about
	// the select statement except the core.
	if len(stmt.SelectCore) != 1 {
		return nil, fmt.Errorf("compound select: %w", ErrUnsupported)
	}
	return c.compileSelectCore(stmt.SelectCore[0])
}

func (c *simpleCompiler) compileSelectCore(core *ast.SelectCore) (command.Command, error) {
	if core.Distinct != nil {
		return nil, fmt.Errorf("distince: %w", ErrUnsupported)
	}

	// compile the projection columns

	// cols are the projection columns.
	var cols []command.Column
	for _, resultColumn := range core.ResultColumn {
		col, err := c.compileResultColumn(resultColumn)
		if err != nil {
			return nil, fmt.Errorf("result column: %w", err)
		}
		cols = append(cols, col)
	}

	// selectionInput is the scan or join that is selected from.
	var selectionInput command.List
	// if there is only one table to select from, meaning that no join exists
	if len(core.TableOrSubquery) == 1 {
		table, err := c.compileTableOrSubquery(core.TableOrSubquery[0])
		if err != nil {
			return nil, fmt.Errorf("table or subquery: %w", err)
		}

		selectionInput = command.Scan{
			Table: table,
		}
	} else if len(core.TableOrSubquery) == 0 {
		if core.JoinClause == nil {
			return nil, fmt.Errorf("nothing to select from")
		}

		join, err := c.compileJoin(core.JoinClause)
		if err != nil {
			return nil, fmt.Errorf("join: %w", err)
		}
		selectionInput = join
	} else {
		return nil, fmt.Errorf("table and join constellation: %w", ErrUnsupported)
	}

	// filter is the filter expression extracted from the where clause.
	var filter command.Expr
	if core.Expr1 != nil { // WHERE expr1
		compiled, err := c.compileExpr(core.Expr1)
		if err != nil {
			return nil, fmt.Errorf("where: %w", err)
		}
		filter = compiled
	}

	return command.Project{
		Cols: cols,
		Input: command.Select{
			Filter: filter,
			Input:  selectionInput,
		},
	}, nil
}

func (c *simpleCompiler) compileResultColumn(col *ast.ResultColumn) (command.Column, error) {
	if col.Asterisk != nil {
		var tableName string
		if col.TableName != nil {
			tableName = col.TableName.Value()
		}
		return command.Column{
			Table:  tableName,
			Column: command.LiteralExpr{Value: "*"},
		}, nil
	}

	var alias string
	if col.ColumnAlias != nil {
		alias = col.ColumnAlias.Value()
	}

	expr, err := c.compileExpr(col.Expr)
	if err != nil {
		return command.Column{}, fmt.Errorf("expr: %w", err)
	}

	return command.Column{
		Alias:  alias,
		Column: expr,
	}, nil
}

func (c *simpleCompiler) compileExpr(expr *ast.Expr) (command.Expr, error) {
	if expr.LiteralValue == nil {
		return nil, fmt.Errorf("not literal: %w", ErrUnsupported)
	}

	return command.LiteralExpr{Value: expr.LiteralValue.Value()}, nil
}

func (c *simpleCompiler) compileJoin(join *ast.JoinClause) (command.Join, error) {
	if len(join.JoinClausePart) != 0 {
		return command.Join{}, fmt.Errorf("join part: %w", ErrUnsupported)
	}

	left, err := c.compileTableOrSubquery(join.TableOrSubquery)
	if err != nil {
		return command.Join{}, fmt.Errorf("table or subquery: %w", err)
	}
	return command.Join{
		Left: command.Scan{
			Table: left,
		},
	}, nil
}

func (c *simpleCompiler) compileTableOrSubquery(tos *ast.TableOrSubquery) (command.Table, error) {
	if tos.TableName == nil {
		return nil, fmt.Errorf("not simple table: %w", ErrUnsupported)
	}

	var index string
	if tos.Not == nil && tos.IndexName != nil {
		index = tos.IndexName.Value()
	}
	var schema string
	if tos.SchemaName != nil {
		schema = tos.SchemaName.Value()
	}
	return command.SimpleTable{
		Schema:  schema,
		Table:   tos.TableName.Value(),
		Indexed: tos.Not == nil,
		Index:   index,
	}, nil
}
