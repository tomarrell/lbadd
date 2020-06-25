package types

var _ Value = (*StringValue)(nil)

// StringValue is a value of type String.
type StringValue struct {
	value

	// Value is the underlying primitive value.
	Value string
}

// NewString creates a new value of type String.
func NewString(v string) StringValue {
	return StringValue{
		value: value{
			typ: String,
		},
		Value: v,
	}
}

func (v StringValue) String() string {
	return v.Value
}
