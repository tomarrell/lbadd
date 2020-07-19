package engine

import (
	"github.com/tomarrell/lbadd/internal/engine/types"
)

func (e Engine) evaluateFunction(ctx ExecutionContext, fn types.FunctionValue) (types.Value, error) {
	switch fn.Name {
	case "NOW":
		return builtinNow(e.timeProvider)
	case "RANDOM":
		return builtinRand(e.randomProvider)
	}
	return nil, ErrNoSuchFunction(fn.Name)
}
