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
	switch {
	case ast.SelectStmt != nil:
		cmd, err := c.compileSelect(ast.SelectStmt)
		if err != nil {
			return nil, fmt.Errorf("select: %w", err)
		}
		return cmd, nil
	case ast.DeleteStmt != nil:
		cmd, err := c.compileDelete(ast.DeleteStmt)
		if err != nil {
			return nil, fmt.Errorf("delete: %w", err)
		}
		return cmd, nil
	case ast.DropTableStmt != nil:
		cmd, err := c.compileDropTable(ast.DropTableStmt)
		if err != nil {
			return nil, fmt.Errorf("drop table: %w", err)
		}
		return cmd, nil
	case ast.DropIndexStmt != nil:
		cmd, err := c.compileDropIndex(ast.DropIndexStmt)
		if err != nil {
			return nil, fmt.Errorf("drop index: %w", err)
		}
		return cmd, nil
	case ast.DropTriggerStmt != nil:
		cmd, err := c.compileDropTrigger(ast.DropTriggerStmt)
		if err != nil {
			return nil, fmt.Errorf("drop trigger: %w", err)
		}
		return cmd, nil
	case ast.DropViewStmt != nil:
		cmd, err := c.compileDropView(ast.DropViewStmt)
		if err != nil {
			return nil, fmt.Errorf("drop view: %w", err)
		}
		return cmd, nil
	}
	return nil, fmt.Errorf("statement type: %w", ErrUnsupported)
}

func (c *simpleCompiler) compileDropTable(stmt *ast.DropTableStmt) (command.Command, error) {
	cmd := command.DropTable{
		IfExists: stmt.If != nil,
		Name:     stmt.TableName.Value(),
	}
	if stmt.SchemaName != nil {
		cmd.Schema = stmt.SchemaName.Value()
	}
	return cmd, nil
}

func (c *simpleCompiler) compileDropIndex(stmt *ast.DropIndexStmt) (command.Command, error) {
	cmd := command.DropIndex{
		IfExists: stmt.If != nil,
		Name:     stmt.IndexName.Value(),
	}
	if stmt.SchemaName != nil {
		cmd.Schema = stmt.SchemaName.Value()
	}
	return cmd, nil
}

func (c *simpleCompiler) compileDropTrigger(stmt *ast.DropTriggerStmt) (command.Command, error) {
	cmd := command.DropTrigger{
		IfExists: stmt.If != nil,
		Name:     stmt.TriggerName.Value(),
	}
	if stmt.SchemaName != nil {
		cmd.Schema = stmt.SchemaName.Value()
	}
	return cmd, nil
}

func (c *simpleCompiler) compileDropView(stmt *ast.DropViewStmt) (command.Command, error) {
	cmd := command.DropView{
		IfExists: stmt.If != nil,
		Name:     stmt.ViewName.Value(),
	}
	if stmt.SchemaName != nil {
		cmd.Schema = stmt.SchemaName.Value()
	}
	return cmd, nil
}

func (c *simpleCompiler) compileDelete(stmt *ast.DeleteStmt) (command.Command, error) {
	if stmt.WithClause != nil {
		return nil, fmt.Errorf("with: %w", ErrUnsupported)
	}

	var filter command.Expr
	if stmt.Where != nil {
		compiled, err := c.compileExpr(stmt.Expr)
		if err != nil {
			return nil, fmt.Errorf("expr: %w", err)
		}
		filter = compiled
	} else {
		filter = command.ConstantBooleanExpr{Value: true} // constant true
	}

	table, err := c.compileQualifiedTableName(stmt.QualifiedTableName)
	if err != nil {
		return nil, fmt.Errorf("qualified table name: %w", err)
	}
	return command.Delete{
		Table:  table,
		Filter: filter,
	}, nil
}

func (c *simpleCompiler) compileQualifiedTableName(tableName *ast.QualifiedTableName) (command.Table, error) {
	table := command.SimpleTable{
		Table: tableName.TableName.Value(),
	}
	if tableName.SchemaName != nil {
		table.Schema = tableName.SchemaName.Value()
	}
	if tableName.As != nil {
		table.Alias = tableName.Alias.Value()
	}
	if tableName.By != nil {
		table.Indexed = true
		table.Index = tableName.IndexName.Value()
	}
	return table, nil
}

func (c *simpleCompiler) compileSelect(stmt *ast.SelectStmt) (command.Command, error) {
	if len(stmt.SelectCore) != 1 {
		return nil, fmt.Errorf("compound select: %w", ErrUnsupported)
	}

	var cmd command.Command
	// compile the select core
	core, err := c.compileSelectCore(stmt.SelectCore[0])
	if err != nil {
		return nil, fmt.Errorf("core: %w", err)
	}
	cmd = core

	// compile ORDER BY
	if stmt.Order != nil {
		return nil, fmt.Errorf("order: %w", ErrUnsupported)
	}

	// compile LIMIT
	if stmt.Limit != nil {
		// if there is an offset specified, wrap the command in an offset
		if stmt.Expr2 != nil {
			offset, err := c.compileExpr(stmt.Expr2)
			if err != nil {
				return nil, fmt.Errorf("limit offset: %w", err)
			}
			cmd = command.Offset{
				Offset: offset,
				Input:  cmd.(command.List),
			}
		}

		// wrap the command into a limit
		limit, err := c.compileExpr(stmt.Expr1)
		if err != nil {
			return nil, fmt.Errorf("limit from: %w", err)
		}
		cmd = command.Limit{
			Limit: limit,
			Input: cmd.(command.List),
		}
	}

	return cmd, nil
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
