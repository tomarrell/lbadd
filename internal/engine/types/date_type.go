package types

var (
	// Date is the date type. Dates are comparable. A date that is later than
	// another date is considered larger. The name of this type is "Date".
	Date = DateType{
		typ: typ{
			name: "Date",
		},
	}
)

// DateType is a comparable type.
type DateType struct {
	typ
}

// Compare compares two date values. For this to succeed, both values must be of
// type DateValue and be not nil. A date later than another date is considered
// larger. This method will return 1 if left>right, 0 if left==right, and -1 if
// left<right.
func (t DateType) Compare(left, right Value) (int, error) {
	if err := t.ensureHaveThisType(left, right); err != nil {
		return 0, err
	}

	if left.IsNull() {
		return -1, nil
	} else if right.IsNull() {
		return 1, nil
	}

	leftDate := left.(DateValue).Value
	rightDate := right.(DateValue).Value

	if leftDate.After(rightDate) {
		return 1, nil
	} else if rightDate.After(leftDate) {
		return -1, nil
	}
	return 0, nil
}
