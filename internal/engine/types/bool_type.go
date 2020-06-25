package types

var (
	Bool = BoolType{
		typ: typ{
			name: "Bool",
		},
	}
)

type BoolType struct {
	typ
}

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
