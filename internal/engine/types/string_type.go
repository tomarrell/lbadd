package types

import "strings"

var (
	// String is the string type. Strings are comparable. Comparison is done
	// lexicographically. The name of this type is "String".
	String = StringType{
		typ: typ{
			name: "String",
		},
	}
)

var _ Type = (*StringType)(nil)
var _ Value = (*StringValue)(nil)
var _ Comparator = (*StringType)(nil)
var _ Caster = (*StringType)(nil)

// StringType is a comparable type.
type StringType struct {
	typ
}

// Compare for the String is defined as the lexicographical comparison of the
// two underlying primitive values. This method will return 1 if left>right, 0
// if left==right, and -1 if left<right.
func (t StringType) Compare(left, right Value) (int, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return 0, err
	}

	leftString := left.(StringValue).Value
	rightString := right.(StringValue).Value
	return strings.Compare(leftString, rightString), nil
}

// Cast attempts to cast the given value to a String. This is done by returning
// a string representing the string value of the given value.
func (StringType) Cast(v Value) (Value, error) {
	if v.Is(String) {
		return v, nil
	}
	return NewString(v.String()), nil
}
