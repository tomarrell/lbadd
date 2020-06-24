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
func (e Engine) evaluateExpression(expr command.Expr) (types.Value, error) {
	switch ex := expr.(type) {
	case command.ConstantBooleanExpr:
		return types.BoolValue{Value: ex.Value}, nil
	case command.LiteralExpr:
		return types.StringValue{Value: ex.Value}, nil
	}
	return nil, fmt.Errorf("cannot evaluate expression of type %T", expr)
}
