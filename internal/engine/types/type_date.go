package types

import (
	"time"
)

var (
	// Date is the date type. A date is a timestamp, and its base type is a byte
	// slice. Internally, the timestamp is represented by time.Time.
	Date = DateTypeDescriptor{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeBinary,
		},
	}
)

var _ Type = (*DateTypeDescriptor)(nil)
var _ Value = (*DateValue)(nil)

type (
	// DateTypeDescriptor is the type descriptor for date objects. The value is
	// a time.Time.
	DateTypeDescriptor struct {
		genericTypeDescriptor
	}

	// DateValue is a value of type Date.
	DateValue struct {
		// Value is the primitive value of this value object.
		Value time.Time
	}
)

// Compare for the Date is defined the lexicographical comparison between
// the primitive underlying values.
func (DateTypeDescriptor) Compare(left, right Value) (int, error) {
	if !left.Is(Date) {
		return 0, ErrTypeMismatch(Date, left.Type())
	}
	if !right.Is(Date) {
		return 0, ErrTypeMismatch(Date, right.Type())
	}

	leftDate := left.(DateValue).Value
	rightDate := right.(DateValue).Value
	if leftDate.After(rightDate) {
		return 1, nil
	} else if rightDate.After(leftDate) {
		return -1, nil
	}
	return 0, nil
}

func (DateTypeDescriptor) String() string { return "Date" }

// Type returns a blob type.
func (DateValue) Type() Type { return Date }

// Is checks if this value is of type Date.
func (DateValue) Is(t Type) bool { return t == Date }

func (v DateValue) String() string { return v.Value.Format(time.RFC3339Nano) }
