package types

var (
	Date = DateType{
		typ: typ{
			name: "Date",
		},
	}
)

type DateType struct {
	typ
}

func (t DateType) Compare(left, right Value) (int, error) {
	if err := t.ensureCanCompare(left, right); err != nil {
		return 0, err
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
