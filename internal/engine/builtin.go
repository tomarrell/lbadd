package engine

import (
	"fmt"
	"strings"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

var (
	// suppress warnings, TODO: remove
	_ = builtinCount
	_ = builtinUCase
	_ = builtinLCase
	_ = builtinMin
)

func builtinNow(tp timeProvider) (types.Value, error) {
	return types.NewDate(tp()), nil
}

func builtinCount(args ...types.Value) (types.Value, error) {
	return nil, ErrUnimplemented
}

func builtinUCase(args ...types.StringValue) ([]types.StringValue, error) {
	var output []types.StringValue
	for _, arg := range args {
		output = append(output, types.StringValue{Value: strings.ToUpper(arg.Value)})
	}
	return output, nil
}

func builtinLCase(args ...types.StringValue) ([]types.StringValue, error) {
	var output []types.StringValue
	for _, arg := range args {
		output = append(output, types.StringValue{Value: strings.ToLower(arg.Value)})
	}
	return output, nil
}

func builtinMax(args ...types.Value) (types.Value, error) {
	if len(args) == 0 {
		return nil, nil
	}

	if err := ensureSameType(args...); err != nil {
		return nil, err
	}

	largest := args[0] // start at 0 and compare on
	t := largest.Type()
	comparator, ok := t.(types.Comparator)
	if !ok {
		return nil, ErrUncomparable(t)
	}
	for i := 1; i < len(args); i++ {
		res, err := comparator.Compare(largest, args[i])
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
	comparator, ok := t.(types.Comparator)
	if !ok {
		return nil, ErrUncomparable(t)
	}
	for i := 1; i < len(args); i++ {
		res, err := comparator.Compare(smallest, args[i])
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
