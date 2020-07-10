package engine

import (
	"fmt"
	"strconv"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

// evaluateExpression evaluates the given expression to a runtime-constant
// value, meaning that it can only be evaluated to a constant value with a given
// execution context.
func (e Engine) evaluateExpression(ctx ExecutionContext, expr command.Expr) (types.Value, error) {
	switch ex := expr.(type) {
	case command.ConstantBooleanExpr:
		return types.NewBool(ex.Value), nil
	case command.LiteralExpr:
		return e.evaluateLiteralExpr(ctx, ex)
	case command.FunctionExpr:
		return e.evaluateFunctionExpr(ctx, ex)
	}
	return nil, ErrUnimplemented(fmt.Sprintf("evaluate %T", expr))
}

func (e Engine) evaluateMultipleExpressions(ctx ExecutionContext, exprs []command.Expr) ([]types.Value, error) {
	var vals []types.Value
	for _, expr := range exprs {
		evaluated, err := e.evaluateExpression(ctx, expr)
		if err != nil {
			return nil, err
		}
		vals = append(vals, evaluated)
	}
	return vals, nil
}

// evaluateLiteralExpr evaluates the given literal expression based on the
// current execution context. The returned value will either be a numeric value
// (integer or real) or a string value.
func (e Engine) evaluateLiteralExpr(ctx ExecutionContext, expr command.LiteralExpr) (types.Value, error) {
	// Check whether the expression value is a numeric literal. In the future,
	// this evaluation might depend on the execution context.
	if numVal, ok := ToNumericValue(expr.Value); ok {
		return numVal, nil
	}
	// if not a numeric literal, remove quotes and resolve escapes
	resolved, err := strconv.Unquote(expr.Value)
	if err != nil {
		// use the original string
		return types.NewString(expr.Value), nil
	}
	return types.NewString(resolved), nil
}

func (e Engine) evaluateFunctionExpr(ctx ExecutionContext, expr command.FunctionExpr) (types.Value, error) {
	exprs, err := e.evaluateMultipleExpressions(ctx, expr.Args)
	if err != nil {
		return nil, fmt.Errorf("arguments: %w", err)
	}

	function := types.NewFunction(expr.Name, exprs...)
	return e.evaluateFunction(ctx, function)
}
