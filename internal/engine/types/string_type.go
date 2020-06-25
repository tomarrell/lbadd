package types

import "strings"

var (
	// String is the string type. Its base type is a string.
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

// StringType is the type descriptor for a string value.
type StringType struct {
	typ
}

// Compare for the String is defined as the lexicographical comparison of
// the two underlying primitive values.
func (t StringType) Compare(left, right Value) (int, error) {
	if err := t.ensureCanCompare(left, right); err != nil {
		return 0, err
	}

	leftString := left.(StringValue).Value
	rightString := right.(StringValue).Value
	return strings.Compare(leftString, rightString), nil
}

func (StringType) Cast(v Value) (Value, error) {
	if v.Is(String) {
		return v, nil
	}
	return NewString(v.String()), nil
}
