package types

import "fmt"

var _ Value = (*NullValue)(nil)

// NullValue is a value of type null. It has no actual value.
type NullValue struct {
	value
}

// NewNull creates a new value of type null, with the given type.
func NewNull(t Type) NullValue {
	return NullValue{
		value: value{
			typ:    t,
			isNull: true,
		},
	}
}

// String "NULL", appended to the type of this value in parenthesis.
func (v NullValue) String() string {
	return fmt.Sprintf("(%v)NULL", v.typ.Name())
}
