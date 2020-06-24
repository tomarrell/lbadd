package types

import (
	"fmt"
	"strconv"
)

var (
	// Numeric is the numeric type. Its base type is a numeric type.
	Numeric = NumericTypeDescriptor{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeNumeric,
		},
	}
)

var _ Type = (*NumericTypeDescriptor)(nil)
var _ Value = (*NumericValue)(nil)

type (
	// NumericTypeDescriptor is the type descriptor for parameterized and non-parameterized
	// numeric types, such as DECIMAL.
	NumericTypeDescriptor struct {
		genericTypeDescriptor
	}

	// NumericValue is a value of type Numeric.
	NumericValue struct {
		// Value is the underlying primitive value.
		Value float64
	}
)

// Compare for the Numeric is defined as the lexicographical comparison of the
// two underlying primitive values.
func (NumericTypeDescriptor) Compare(left, right Value) (int, error) {
	if !left.Is(Numeric) {
		return 0, ErrTypeMismatch(Numeric, left.Type())
	}
	if !right.Is(Numeric) {
		return 0, ErrTypeMismatch(Numeric, right.Type())
	}

	leftNum := left.(NumericValue).Value
	rightNum := right.(NumericValue).Value
	switch {
	case leftNum < rightNum:
		return -1, nil
	case leftNum == rightNum:
		return 0, nil
	case leftNum > rightNum:
		return 1, nil
	}
	return -2, fmt.Errorf("unhandled constellation: %v <-> %v", leftNum, rightNum)
}

func (NumericTypeDescriptor) String() string { return "Numeric" }

// Type returns a string type.
func (NumericValue) Type() Type { return Numeric }

// Is checks if this value is of type Numeric.
func (NumericValue) Is(t Type) bool { return t == Numeric }

func (v NumericValue) String() string { return strconv.FormatFloat(v.Value, 'f', -1, 64) }
