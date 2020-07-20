package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

func (e Engine) add(ctx ExecutionContext, left, right types.Value) (types.Value, error) {
	if left == nil || right == nil {
		return nil, fmt.Errorf("cannot add %T and %T", left, right)
	}

	if adder, ok := left.Type().(types.ArithmeticAdder); ok {
		result, err := adder.Add(left, right)
		if err != nil {
			return nil, fmt.Errorf("add: %w", err)
		}
		return result, nil
	}
	return nil, fmt.Errorf("%v does not support addition", left.Type())
}

func (e Engine) sub(ctx ExecutionContext, left, right types.Value) (types.Value, error) {
	if left == nil || right == nil {
		return nil, fmt.Errorf("cannot subtract %T and %T", left, right)
	}

	if subtractor, ok := left.Type().(types.ArithmeticSubtractor); ok {
		result, err := subtractor.Sub(left, right)
		if err != nil {
			return nil, fmt.Errorf("sub: %w", err)
		}
		return result, nil
	}
	return nil, fmt.Errorf("%v does not support subtraction", left.Type())
}

func (e Engine) mul(ctx ExecutionContext, left, right types.Value) (types.Value, error) {
	if left == nil || right == nil {
		return nil, fmt.Errorf("cannot multiplicate %T and %T", left, right)
	}

	if multiplicator, ok := left.Type().(types.ArithmeticMultiplicator); ok {
		result, err := multiplicator.Mul(left, right)
		if err != nil {
			return nil, fmt.Errorf("mul: %w", err)
		}
		return result, nil
	}
	return nil, fmt.Errorf("%v does not support multiplication", left.Type())
}

func (e Engine) div(ctx ExecutionContext, left, right types.Value) (types.Value, error) {
	if left == nil || right == nil {
		return nil, fmt.Errorf("cannot divide %T and %T", left, right)
	}

	if divider, ok := left.Type().(types.ArithmeticDivider); ok {
		result, err := divider.Div(left, right)
		if err != nil {
			return nil, fmt.Errorf("div: %w", err)
		}
		return result, nil
	}
	return nil, fmt.Errorf("%v does not support division", left.Type())
}

func (e Engine) mod(ctx ExecutionContext, left, right types.Value) (types.Value, error) {
	if left == nil || right == nil {
		return nil, fmt.Errorf("cannot modulo %T and %T", left, right)
	}

	if modulator, ok := left.Type().(types.ArithmeticModulator); ok {
		result, err := modulator.Mod(left, right)
		if err != nil {
			return nil, fmt.Errorf("mod: %w", err)
		}
		return result, nil
	}
	return nil, fmt.Errorf("%v does not support modulo", left.Type())
}

func (e Engine) pow(ctx ExecutionContext, left, right types.Value) (types.Value, error) {
	if left == nil || right == nil {
		return nil, fmt.Errorf("cannot exponentiate %T and %T", left, right)
	}

	if exponentiator, ok := left.Type().(types.ArithmeticExponentiator); ok {
		result, err := exponentiator.Pow(left, right)
		if err != nil {
			return nil, fmt.Errorf("pow: %w", err)
		}
		return result, nil
	}
	return nil, fmt.Errorf("%v does not support exponentiation", left.Type())
}
