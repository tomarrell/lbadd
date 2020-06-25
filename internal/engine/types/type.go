package types

import "fmt"

type (
	// Comparator is the interface that wraps the basic compare method. The
	// compare method compares the left and right value as follows. -1 if
	// left<right, 0 if left==right, 1 if left>right. What exectly is considered
	// to be <, ==, > is up to the implementation.
	Comparator interface {
		// Compare compares the given to values left and right as follows. -1 if
		// left<right, 0 if left==right, 1 if left>right.
		Compare(left, right Value) (int, error)
	}

	// Caster wraps the Cast method, which is used to transform the input value
	// into an output value. Types can implement this interface. E.g. if the
	// type String implements Caster, any value passed into the Cast method
	// should be attempted to be cast to String, or an error should be returned.
	Caster interface {
		Cast(Value) (Value, error)
	}

	// Codec describes a component that can encode and decode values. Types
	// embed codec, but are only expected to be able to encode and decode values
	// of their own type. If that is not the case, a type mismatch may be
	// returned.
	Codec interface {
		Encode(Value) ([]byte, error)
		Decode([]byte) (Value, error)
	}

	// Type is a data type that consists of a type descriptor and a comparator.
	// The comparator forces types to define relations between two values of
	// this type. A type is only expected to be able to handle values of its own
	// type.
	Type interface {
		Name() string
		fmt.Stringer
	}
)

type typ struct {
	name string
}

func (t typ) Name() string   { return t.name }
func (t typ) String() string { return t.name }

func (t typ) ensureCanCompare(left, right Value) error {
	if left == nil || right == nil {
		return ErrTypeMismatch(t, nil)
	}
	if !left.Is(t) {
		return ErrTypeMismatch(t, left.Type())
	}
	if !right.Is(t) {
		return ErrTypeMismatch(t, right.Type())
	}
	return nil
}
