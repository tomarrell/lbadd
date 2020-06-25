package types

import "fmt"

type (
	// Comparator is the interface that wraps the basic compare method. The
	// compare method compares the left and right value as follows. -1 if
	// left<right, 0 if left==right, 1 if left>right. What exectly is considered
	// to be <, ==, > is up to the implementation.
	Comparator interface {
		Compare(Value, Value) (int, error)
	}

	// Caster wraps the Cast method, which is used to transform the input value
	// into an output value. Types can implement this interface. E.g. if the
	// type String implements Caster, any value passed into the Cast method
	// should be attempted to be cast to String, or an error should be returned.
	Caster interface {
		Cast(Value) (Value, error)
	}

	// Type is a data type that consists of a type descriptor and a comparator.
	// The comparator forces types to define relations between two values of
	// this type. A type is only expected to be able to handle values of its own
	// type.
	Type interface {
		TypeDescriptor
		Comparator
		fmt.Stringer
	}
)
