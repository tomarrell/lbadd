package engine

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/compiler/command"
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func (e Engine) evaluate(c command.Command) (Result, error) {
	switch cmd := c.(type) {
	case command.Values:
		_, _ = e.evaluateValues(cmd)
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
