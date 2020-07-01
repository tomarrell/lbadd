package types

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

	leftInteger := left.(IntegerValue).Value
	rightInteger := right.(IntegerValue).Value

	if leftInteger < rightInteger {
		return -1, nil
	} else if rightInteger > leftInteger {
		return 1, nil
	}
	return 0, nil
}
