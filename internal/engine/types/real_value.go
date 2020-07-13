package types

import (
	"strconv"
)

var _ Value = (*RealValue)(nil)

// RealValue is a value of type Real.
type RealValue struct {
	value

	// Value is the underlying primitive value.
	Value float64
}

// NewReal creates a new value of type Real.
func NewReal(v float64) RealValue {
	return RealValue{
		value: value{
			typ: Real,
		},
		Value: v,
	}
}

func (v RealValue) String() string {
	return strconv.FormatFloat(v.Value, 'e', -1, 8)
}
