package types

import (
	"strconv"
)

var _ Value = (*BoolValue)(nil)

// BoolValue is a value of type Bool.
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

func (v BoolValue) String() string {
	return strconv.FormatBool(v.Value)
}
