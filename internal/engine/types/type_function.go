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

	// FunctionValue represents a callable function. It consists of a nameand
	// arguments. How the function is evaluated and what code is actually
	// executed, is decided by the engine.
	FunctionValue struct {
		Name string
		Args []Value
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
func NewFunctionValue(name string, args ...Value) FunctionValue {
	return FunctionValue{
		Name: name,
		Args: args,
	}
}

// Type returns a function type.
func (FunctionValue) Type() Type { return Function }

// Is checks if this value is of type function.
func (FunctionValue) Is(t Type) bool { return t == Function }

func (f FunctionValue) String() string { return f.Name + "()" }
