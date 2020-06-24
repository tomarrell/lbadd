package types

//go:generate stringer -type=BaseType

// BaseType is an underlying type for parameterized types.
type BaseType uint8

// Known base types.
const (
	BaseTypeUnknown BaseType = iota
	BaseTypeBool
	BaseTypeBinary
	BaseTypeString
	BaseTypeNumeric
	BaseTypeFunction
)

// TypeDescriptor describes a type in more detail than just the type. Every type
// has a type descriptor, which holds the base type and parameterization. Based
// on the parameterization, the type may be interpreted differently.
//
// Example: The simple type INTEGER would have a type descriptor that describes
// a baseType=number with no further parameterization, whereas the more complex
// type VARCHAR(50) would have a type descriptor, that describes a
// baseType=string and a max length of 50.
type TypeDescriptor interface {
	Base() BaseType
	// TODO: parameters to be done
}

// genericTypeDescriptor is a type descriptor that has no parameterization and
// just a base type.
type genericTypeDescriptor struct {
	baseType BaseType
}

// Base returns the base type of this type descriptor.
func (td genericTypeDescriptor) Base() BaseType {
	return td.baseType
}
