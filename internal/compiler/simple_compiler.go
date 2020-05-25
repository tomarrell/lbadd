package compiler

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/compiler/optimization"
	"github.com/tomarrell/lbadd/internal/parser/ast"
)

const (
	// decimalPoint is the rune '.', and it cannot be ',' because of ambiguity
	// in the grammar (could be interpreted as a token separator).
	decimalPoint rune = '.'
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
	if ast.SelectStmt != nil {
		cmd, err := c.compileSelect(ast.SelectStmt)
		if err != nil {
			return nil, fmt.Errorf("select: %w", err)
		}
		return cmd, nil
	}
	return nil, fmt.Errorf("not select: %w", ErrUnsupported)
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
	if core.CompoundOperator != nil {
		return nil, fmt.Errorf("compound statements: %w", ErrUnsupported)
	}

	if core.Values != nil {
		return c.compileSelectCoreValues(core)
	}
	return c.compileSelectCoreSelect(core)
}

func (c *simpleCompiler) compileSelectCoreValues(core *ast.SelectCore) (command.Command, error) {
	var datasets [][]command.Expr
	for _, parExpr := range core.ParenthesizedExpressions {
		var values []command.Expr
		for _, expr := range parExpr.Exprs {
			compiled, err := c.compileExpr(expr)
			if err != nil {
				return nil, fmt.Errorf("expr: %w", err)
			}
			values = append(values, compiled)
		}
		datasets = append(datasets, values)
	}
	return command.Values{Values: datasets}, nil
}

func (c *simpleCompiler) compileSelectCoreSelect(core *ast.SelectCore) (command.Command, error) {
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

	// only wrap into select if there is a filter, otherwise there is no need
	// for the select
	input := selectionInput
	if filter != nil {
		input = command.Select{
			Filter: filter,
			Input:  input,
		}
	}

	// wrap columns and input into projection
	var list command.List
	list = command.Project{
		Cols:  cols,
		Input: input,
	}

	// wrap list into distinct if needed
	if core.Distinct != nil {
		list = command.Distinct{
			Input: list,
		}
	}
	return list, nil
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
	switch {
	case expr.LiteralValue != nil:
		return command.LiteralExpr{Value: expr.LiteralValue.Value()}, nil
	case expr.UnaryOperator != nil:
		val, err := c.compileExpr(expr.Expr1)
		if err != nil {
			return nil, fmt.Errorf("expr1: %w", err)
		}
		return command.UnaryExpr{
			Operator: expr.UnaryOperator.Value(),
			Value:    val,
		}, nil
	case expr.BinaryOperator != nil:
		left, err := c.compileExpr(expr.Expr1)
		if err != nil {
			return nil, fmt.Errorf("expr1: %w", err)
		}
		right, err := c.compileExpr(expr.Expr2)
		if err != nil {
			return nil, fmt.Errorf("expr2: %w", err)
		}
		return command.BinaryExpr{
			Operator: expr.BinaryOperator.Value(),
			Left:     left,
			Right:    right,
		}, nil
	case expr.FunctionName != nil:
		if !(expr.FilterClause == nil && expr.OverClause == nil) {
			return nil, fmt.Errorf("filter or over on function: %w", ErrUnsupported)
		}
		if expr.Asterisk != nil {
			return nil, fmt.Errorf("function_name(*): %w", ErrUnsupported)
		}

		var args []command.Expr
		for _, arg := range expr.Expr {
			compiledArg, err := c.compileExpr(arg)
			if err != nil {
				return nil, fmt.Errorf("expr: %w", err)
			}
			args = append(args, compiledArg)
		}

		return command.FunctionExpr{
			Name:     expr.FunctionName.Value(),
			Distinct: expr.Distinct != nil,
			Args:     args,
		}, nil
	}

	return nil, ErrUnsupported
}

func (c *simpleCompiler) compileJoin(join *ast.JoinClause) (command.List, error) {
	left, err := c.compileTableOrSubquery(join.TableOrSubquery)
	if err != nil {
		return command.Join{}, fmt.Errorf("table or subquery: %w", err)
	}

	var prev command.List
	prev = command.Scan{
		Table: left,
	}

	for _, part := range join.JoinClausePart {
		if part.JoinConstraint != nil && part.JoinConstraint.Using != nil {
			return command.Join{}, fmt.Errorf("using: %w", ErrUnsupported)
		}

		op := part.JoinOperator
		// evaluate join type
		var typ command.JoinType
		var natural bool
		if op.Natural != nil {
			natural = true
		}
		if op.Left != nil {
			if op.Outer != nil {
				typ = command.JoinLeftOuter
			} else {
				typ = command.JoinLeft
			}
		} else if op.Inner != nil {
			typ = command.JoinInner
		} else if op.Cross != nil {
			typ = command.JoinCross
		}

		var filter command.Expr
		if part.JoinConstraint != nil && part.JoinConstraint.On != nil {
			filter, err = c.compileExpr(part.JoinConstraint.Expr)
			if err != nil {
				return nil, fmt.Errorf("expressoin: %w", err)
			}
		}

		table, err := c.compileTableOrSubquery(part.TableOrSubquery)
		if err != nil {
			return command.Join{}, fmt.Errorf("table or subquery: %w", err)
		}

		prev = command.Join{
			Natural: natural,
			Type:    typ,
			Filter:  filter,
			Left:    prev,
			Right: command.Scan{
				Table: table,
			},
		}
	}

	return prev, nil
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
		Indexed: tos.By != nil,
		Index:   index,
	}, nil
}
