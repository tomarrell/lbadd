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
