package types

import (
	"strconv"
)

var _ Value = (*BoolValue)(nil)

// BoolValue is a value of type Bool. The boolean value is held as a primitive
// Go bool.
type BoolValue struct {
	value

	// Value is the underlying primitive value.
	Value bool
}

// NewBool creates a new value of type Bool.
func NewBool(v bool) BoolValue {
	return BoolValue{
		value: value{
			typ: Bool,
		},
		Value: v,
	}
}

// String returns "true" or "false", depending on the value of this bool value.
func (v BoolValue) String() string {
	return strconv.FormatBool(v.Value)
}
