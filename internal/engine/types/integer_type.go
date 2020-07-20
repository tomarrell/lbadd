package types

import "math"

var (
	// Integer is the date type. Integers are comparable. The name of this type
	// is "Integer".
	Integer = IntegerType{
		typ: typ{
			name: "Integer",
		},
	}
)

// IntegerType is a comparable type.
type IntegerType struct {
	typ
}

// Compare compares two date values. For this to succeed, both values must be of
// type IntegerValue and be not nil. A date later than another date is considered
// larger. This method will return 1 if left>right, 0 if left==right, and -1 if
// left<right.
func (t IntegerType) Compare(left, right Value) (int, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return 0, err
	}

	if left.IsNull() {
		return -1, nil
	} else if right.IsNull() {
		return 1, nil
	}

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	if leftInteger < rightInteger {
		return -1, nil
	} else if rightInteger > leftInteger {
		return 1, nil
	}
	return 0, nil
}

// Add adds the left and right value, producing a new integer value. This only
// works, if left and right are of type integer.
func (t IntegerType) Add(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	return NewInteger(leftInteger + rightInteger), nil
}

// Sub subtracts the right from the left value, producing a new integer value.
// This only works, if left and right are of type integer.
func (t IntegerType) Sub(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	return NewInteger(leftInteger - rightInteger), nil
}

// Mul multiplicates the left and right value, producing a new integer value.
// This only works, if left and right are of type integer.
func (t IntegerType) Mul(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	return NewInteger(leftInteger * rightInteger), nil
}

// Div divides the left by the right value, producing a new real value. This
// only works, if left and right are of type integer.
func (t IntegerType) Div(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	return NewReal(float64(leftInteger) / float64(rightInteger)), nil
}

// Mod modulates the left and right value, producing a new integer value. This
// only works, if left and right are of type integer.
func (t IntegerType) Mod(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	return NewInteger(leftInteger % rightInteger), nil
}

// Pow exponentiates the left and right value, producing a new integer value.
// This only works, if left and right are of type integer.
func (t IntegerType) Pow(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	return NewInteger(int64(math.Pow(float64(leftInteger), float64(rightInteger)))), nil
}
