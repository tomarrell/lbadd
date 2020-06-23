package engine

import "fmt"

var (
	numericType = NumericType{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeString,
		},
	}
)

var _ Type = (*NumericType)(nil)
var _ Value = (*NumericValue)(nil)

type (
	// NumericType is the type for parameterized and non-parameterized numeric
	// types, such as DECIMAL.
	NumericType struct {
		genericTypeDescriptor
	}

	// NumericValue is a value of type NumericType.
	NumericValue struct {
		// Value is the underlying primitive value.
		Value float64
	}
)

// Compare for the NumericType is defined as the lexicographical comparison of
// the two underlying primitive values.
func (NumericType) Compare(left, right Value) (int, error) {
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

// Type returns a string type.
func (NumericValue) Type() Type { return numericType }
