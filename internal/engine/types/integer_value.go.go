package types

import (
	"strconv"
)

var _ Value = (*IntegerValue)(nil)

// IntegerValue is a value of type Integer.
type IntegerValue struct {
	value

	// Value is the underlying primitive value.
	Value int64
}

// NewInteger creates a new value of type Integer.
func NewInteger(v int64) IntegerValue {
	return IntegerValue{
		value: value{
			typ: Integer,
		},
		Value: v,
	}
}

func (v IntegerValue) String() string {
	return strconv.FormatInt(v.Value, 10)
}
