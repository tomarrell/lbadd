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
var _ Serializer = (*StringType)(nil)

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

	if left.IsNull() {
		return -1, nil
	} else if right.IsNull() {
		return 1, nil
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

// Serialize serializes the internal string value as 4-byte-framed byte
// sequence.
func (t StringType) Serialize(v Value) ([]byte, error) {
	if err := t.ensureHasThisType(v); err != nil {
		return nil, err
	}

	str := v.(StringValue).Value
	payload := frame([]byte(str))
	payloadLength := len(payload)
	data := make([]byte, 1+payloadLength) // string byte length + 1 for the indicator
	data[0] = byte(TypeIndicatorString)
	copy(data[1:], str)

	return data, nil
}

// Deserialize reads the data size from the first 4 passed-in bytes, and then
// converts the rest of the bytes to a string leveraging the Go runtime.
func (t StringType) Deserialize(data []byte) (Value, error) {
	payloadSize := int(byteOrder.Uint32(data[0:]))
	if payloadSize+4 != len(data) {
		return nil, ErrDataSizeMismatch(payloadSize+4, len(data))
	}
	return NewString(string(data[4:])), nil
}

// Add concatenates the left and right right value. This only works, if left and
// right are string values.
func (t StringType) Add(left, right Value) (Value, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return nil, err
	}

	leftString := left.(StringValue).Value
	rightString := right.(StringValue).Value
	return NewString(leftString + rightString), nil
}
