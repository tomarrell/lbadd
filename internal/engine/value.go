package engine

// Value describes an object that can be used as a value in the engine. A value
// has a type, which can be used to compare this value to another one.
type Value interface {
	Type() Type
}
