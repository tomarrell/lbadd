package types

import "time"

var _ Value = (*DateValue)(nil)

// DateValue is a value of type Date.
type DateValue struct {
	value

	// Value is the underlying primitive value.
	Value time.Time
}

// NewDate creates a new value of type Date.
func NewDate(v time.Time) DateValue {
	return DateValue{
		value: value{
			typ: Date,
		},
		Value: v,
	}
}

func (v DateValue) String() string {
	return v.Value.Format(time.RFC3339)
}
