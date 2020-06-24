package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

type builtinFunction = func(...types.Value) (types.Value, error)

func builtinRandom(args ...types.Value) (types.Value, error) {
	return nil, ErrUnimplemented
}

func builtinCount(args ...types.Value) (types.Value, error) {
	return types.NumericValue{Value: float64(len(args))}, nil
}

func builtinUCase(args ...types.Value) (types.Value, error) {
	return nil, ErrUnimplemented
}

func builtinLCase(args ...types.Value) (types.Value, error) {
	return nil, ErrUnimplemented
}

func builtinNow(args ...types.Value) (types.Value, error) {
	return nil, ErrUnimplemented
}

func builtinMax(args ...types.Value) (types.Value, error) {
	if len(args) == 0 {
		return nil, nil
	}

	if err := ensureSameType(args...); err != nil {
		return nil, err
	}

	largest := args[0]
	t := largest.Type()
	for i := 1; i < len(args); i++ {
		res, err := t.Compare(largest, args[i])
		if err != nil {
			return nil, fmt.Errorf("compare: %w", err)
		}
		if res < 0 {
			largest = args[i]
		}
	}
	return largest, nil
}

func builtinMin(args ...types.Value) (types.Value, error) {
	if len(args) == 0 {
		return nil, nil
	}

	if err := ensureSameType(args...); err != nil {
		return nil, err
	}

	smallest := args[0]
	t := smallest.Type()
	for i := 1; i < len(args); i++ {
		res, err := t.Compare(smallest, args[i])
		if err != nil {
			return nil, fmt.Errorf("compare: %w", err)
		}
		if res > 0 {
			smallest = args[i]
		}
	}
	return smallest, nil
}

func ensureSameType(args ...types.Value) error {
	if len(args) == 0 {
		return nil
	}

	base := args[0]
	for i := 1; i < len(args); i++ {
		if !base.Is(args[i].Type()) { // Is is transitive
			return types.ErrTypeMismatch(base.Type(), args[i].Type())
		}
	}
	return nil
}
