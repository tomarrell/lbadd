package types

import "fmt"

type (
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

func (t typ) ensureHaveThisType(left, right Value) error {
	if err := t.ensureHasThisType(left); err != nil {
		return err
	}
	if err := t.ensureHasThisType(right); err != nil {
		return err
	}
	return nil
}

func (t typ) ensureHasThisType(v Value) error {
	if v == nil {
		return ErrTypeMismatch(t, nil)
	}
	if !v.Is(t) {
		return ErrTypeMismatch(t, v.Type())
	}
	return nil
}
