package types

import "fmt"

// Value describes an object that can be used as a value in the engine. A value
// has a type, which can be used to compare this value to another one.
type Value interface {
	Type() Type
	IsTyper
	IsNuller
	fmt.Stringer
}

// IsNuller wraps the method IsNull, which determines whether the value is a
// null value or not.
type IsNuller interface {
	IsNull() bool
}

// IsTyper wraps the basic method Is, which can check the type of a value.
type IsTyper interface {
	Is(Type) bool
}

type value struct {
	typ    Type
	isNull bool
}

func (v value) IsNull() bool {
	return v.isNull
}

func (v value) Type() Type {
	return v.typ
}

// Is checks whether this value is of the given type.
func (v value) Is(t Type) bool {
	return v.typ.Name() == t.Name()
}
