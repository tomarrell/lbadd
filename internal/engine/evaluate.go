package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/result"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func (e Engine) evaluate(c command.Command) (result.Result, error) {
	switch cmd := c.(type) {
	case command.Values:
		values, err := e.evaluateValues(cmd)
		if err != nil {
			return nil, fmt.Errorf("values: %w", err)
		}
		res, err := result.FromValues(values)
		if err != nil {
			return nil, fmt.Errorf("from values: %w", err)
		}
		return res, nil
	}
	return nil, nil
}

func (e Engine) evaluateValues(v command.Values) ([][]types.Value, error) {
	result := make([][]types.Value, len(v.Values))
	for y, values := range v.Values {
		rowValues := make([]types.Value, len(values))
		for x, value := range values {
			internalValue, err := e.evaluateExpression(value)
			if err != nil {
				return nil, fmt.Errorf("expr: %w", err)
			}
			rowValues[x] = internalValue
		}
		result[y] = rowValues
	}
	return result, nil
}
