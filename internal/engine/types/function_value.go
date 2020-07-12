package types

import (
	"fmt"
	"strings"
)

var _ Value = (*FunctionValue)(nil)

// FunctionValue is a value of type Function. This can not be called, it is
// simply a shell that holds a function name and the arguments, that were used
// in the SQL statement. It is the engine's responsibility, to execute the
// appropriate code to make this function call happen.
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
