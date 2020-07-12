package types

var (
	// Bool is the Bool type. Bools are comparable with true>false. The name of
	// this type is "Bool".
	Bool = BoolType{
		typ: typ{
			name: "Bool",
		},
	}
)

var _ Type = (*BoolType)(nil)       // BoolType is a type
var _ Comparator = (*BoolType)(nil) // BoolType is comparable
var _ Serializer = (*BoolType)(nil) // BoolType is serializable

// BoolType is a basic type. Values of this type describe a boolean value,
// either true or false.
type BoolType struct {
	typ
}

// Compare compares two bool values. For this to succeed, both values must be of
// type BoolValue and be not nil. The bool value true is considered larger than
// false. This method will return 1 if left>right, 0 if left==right, and -1 if
// left<right.
func (t BoolType) Compare(left, right Value) (int, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return 0, err
	}

	if left.IsNull() {
		return -1, nil
	} else if right.IsNull() {
		return 1, nil
	}

	leftBool := left.(BoolValue).Value
	rightBool := right.(BoolValue).Value

	if leftBool && !rightBool {
		return 1, nil
	} else if !leftBool && rightBool {
		return -1, nil
	}
	return 0, nil
}

// Serialize can serialize bool values to a single byte. false will be
// serialized to 0x00, while true will be serialized to 0x01.
func (t BoolType) Serialize(v Value) ([]byte, error) {
	if err := t.ensureHasThisType(v); err != nil {
		return nil, err
	}

	val, ok := v.(BoolValue)
	if !ok {
		return nil, ErrTypeMismatch(Bool, v.Type())
	}
	result := []byte{0x00}
	if val.Value {
		result[0]++
	}
	return result, nil
}

// Deserialize can deserialize a single byte to a bool value. If the length of
// the given data is not 1, an error will be returned. A zero byte will be
// deserialized to the value false, while everything else will be deserialized
// to the value true.
func (BoolType) Deserialize(data []byte) (Value, error) {
	if len(data) != 1 {
		return nil, ErrDataSizeMismatch(1, len(data))
	}
	return NewBool(data[0] != 0), nil
}
