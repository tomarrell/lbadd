package types

import "fmt"

// Value describes an object that can be used as a value in the engine. A value
// has a type, which can be used to compare this value to another one.
type Value interface {
	Type() Type
	IsTyper
	fmt.Stringer
}

// IsTyper wraps the basic method Is, which can check the type of a value.
type IsTyper interface {
	Is(Type) bool
}

type value struct {
	typ Type
}

func (v value) Type() Type {
	return v.typ
}

// Is checks whether this value is of the given type.
func (v value) Is(t Type) bool {
	return v.typ.Name() == t.Name()
}
