package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/compiler/command"
)

// evaluateExpression evaluates the given expression to a runtime-constant
// value, meaning that it can only be evaluated to a constant value with a given
// execution context. This execution context must be inferred from the engine
// receiver.
func (e Engine) evaluateExpression(expr command.Expr) (Value, error) {
	switch ex := expr.(type) {
	case command.ConstantBooleanExpr:
		return BoolValue{Value: ex.Value}, nil
	}
	return nil, fmt.Errorf("cannot evaluate expression of type %T", expr)
}
