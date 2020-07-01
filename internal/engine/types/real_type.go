package types

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

	leftReal := left.(RealValue).Value
	rightReal := right.(RealValue).Value

	if leftReal < rightReal {
		return -1, nil
	} else if rightReal > leftReal {
		return 1, nil
	}
	return 0, nil
}
