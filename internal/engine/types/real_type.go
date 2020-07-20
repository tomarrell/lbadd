package types

import "math"

var (
	// Real is the date type. Reals are comparable. The name of this type
	// is "Real".
	Real = RealType{
		typ: typ{
			name: "Real",
		},
	}
)

// RealType is a comparable type.
type RealType struct {
	typ
}

// Compare compares two real values. For this to succeed, both values must be of
// type RealValue and be not nil.
func (t RealType) Compare(left, right Value) (int, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return 0, err
	}

	if left.IsNull() {
		return -1, nil
	} else if right.IsNull() {
		return 1, nil
	}

	leftReal := left.(RealValue).Value
	rightReal := right.(RealValue).Value

	if leftReal < rightReal {
		return -1, nil
	} else if rightReal > leftReal {
		return 1, nil
	}
	return 0, nil
}

// Add adds the left and right value, producing a new real value. This only
// works, if left and right are of type real.
func (t RealType) Add(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftReal := left.(RealValue).Value
	rightReal := right.(RealValue).Value

	return NewReal(leftReal + rightReal), nil
}

// Sub subtracts the right from the left value, producing a new real value.
// This only works, if left and right are of type real.
func (t RealType) Sub(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftReal := left.(RealValue).Value
	rightReal := right.(RealValue).Value

	return NewReal(leftReal - rightReal), nil
}

// Mul multiplicates the left and right value, producing a new real value.
// This only works, if left and right are of type real.
func (t RealType) Mul(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftReal := left.(RealValue).Value
	rightReal := right.(RealValue).Value

	return NewReal(leftReal * rightReal), nil
}

// Div divides the left by the right value, producing a new real value. This
// only works, if left and right are of type real.
func (t RealType) Div(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftReal := left.(RealValue).Value
	rightReal := right.(RealValue).Value

	return NewReal(float64(leftReal) / float64(rightReal)), nil
}

// Pow exponentiates the left and right value, producing a new real value.
// This only works, if left and right are of type real.
func (t RealType) Pow(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftReal := left.(RealValue).Value
	rightReal := right.(RealValue).Value

	return NewReal(math.Pow(leftReal, rightReal)), nil
}
