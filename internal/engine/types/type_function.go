package types

import "fmt"

var (
	// Function is the function type. Functions are not comparable.
	Function = FunctionTypeDescriptor{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeFunction,
		},
	}
)

type (
	// FunctionTypeDescriptor is the function type of this engine.
	FunctionTypeDescriptor struct {
		genericTypeDescriptor
	}

	functionValue struct {
		Name  string
		value func(...Value) (Value, error)
	}
)

// Compare will always return an error, indicating that functions are not
// comparable. Both arguments have to have a funtion type, otherwise a type
// mismatch error will be returned.
func (FunctionTypeDescriptor) Compare(left, right Value) (int, error) {
	if !left.Is(Function) {
		return 0, ErrTypeMismatch(Function, left.Type())
	}
	if !right.Is(Function) {
		return 0, ErrTypeMismatch(Function, right.Type())
	}
	return -2, fmt.Errorf("functions are not comparable")
}

func (FunctionTypeDescriptor) String() string { return "Function" }

// NewFunctionValue creates a new function value with the given name and
// underlying function.
func NewFunctionValue(name string, fn func(...Value) (Value, error)) Value {
	return functionValue{
		Name:  name,
		value: fn,
	}
}

// CallWithArgs will call the underlying function with the given arguments.
func (f functionValue) CallWithArgs(args ...Value) (Value, error) {
	result, err := f.value(args...)
	if err != nil {
		return nil, fmt.Errorf("call %v: %w", f.Name, err)
	}
	return result, nil
}

// Type returns a function type.
func (functionValue) Type() Type { return Function }

// Is checks if this value is of type function.
func (functionValue) Is(t Type) bool { return t == Function }
