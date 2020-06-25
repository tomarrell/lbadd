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

// BoolType is a comparable type.
type BoolType struct {
	typ
}

// Compare compares two bool values. For this to succeed, both values must be of
// type BoolValue and be not nil. The bool value true is considered larger than
// false. This method will return 1 if left>right, 0 if left==right, and -1 if
// left<right.
func (t BoolType) Compare(left, right Value) (int, error) {
	if err := t.ensureCanCompare(left, right); err != nil {
		return 0, err
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
