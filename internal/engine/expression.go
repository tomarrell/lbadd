package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

// evaluateExpression evaluates the given expression to a runtime-constant
// value, meaning that it can only be evaluated to a constant value with a given
// execution context. This execution context must be inferred from the engine
// receiver.
func (e Engine) evaluateExpression(ctx ExecutionContext, expr command.Expr) (types.Value, error) {
	switch ex := expr.(type) {
	case command.ConstantBooleanExpr:
		return types.BoolValue{Value: ex.Value}, nil
	case command.LiteralExpr:
		return e.evaluateLiteralExpr(ctx, ex)
	case command.FunctionExpr:
		return e.evaluateFunctionExpr(ctx, ex)
	}
	return nil, fmt.Errorf("cannot evaluate expression of type %T", expr)
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

func (e Engine) evaluateLiteralExpr(ctx ExecutionContext, expr command.LiteralExpr) (types.Value, error) {
	return types.StringValue{Value: expr.Value}, nil
}

func (e Engine) evaluateFunctionExpr(ctx ExecutionContext, expr command.FunctionExpr) (types.Value, error) {
	exprs, err := e.evaluateMultipleExpressions(ctx, expr.Args)
	if err != nil {
		return nil, fmt.Errorf("arguments: %w", err)
	}

	function := types.NewFunctionValue(expr.Name, exprs...)
	return e.evaluateFunction(ctx, function)
}
