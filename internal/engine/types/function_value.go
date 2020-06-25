package types

import (
	"fmt"
	"strings"
)

var _ Value = (*FunctionValue)(nil)

// FunctionValue is a value of type Function.
type FunctionValue struct {
	value

	Name string
	Args []Value
}

// NewFunction creates a new value of type Function.
func NewFunction(name string, args ...Value) FunctionValue {
	return FunctionValue{
		value: value{
			typ: Function,
		},
		Name: name,
		Args: args,
	}
}

func (v FunctionValue) String() string {
	var args []string
	for _, arg := range v.Args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%v(%v)", v.Name, strings.Join(args, ", "))
}
