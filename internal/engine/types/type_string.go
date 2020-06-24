package types

import "strings"

var (
	// String is the string type. Its base type is a string.
	String = StringTypeDescriptor{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeString,
		},
	}
)

var _ Type = (*StringTypeDescriptor)(nil)
var _ Value = (*StringValue)(nil)

type (
	// StringTypeDescriptor is the type descriptor for parameterized and
	// non-parameterized string types, such as VARCHAR or VARCHAR(n).
	StringTypeDescriptor struct {
		genericTypeDescriptor
	}

	// StringValue is a value of type String.
	StringValue struct {
		// Value is the underlying primitive value.
		Value string
	}
)

// Compare for the String is defined as the lexicographical comparison of
// the two underlying primitive values.
func (StringTypeDescriptor) Compare(left, right Value) (int, error) {
	if !left.Is(String) {
		return 0, ErrTypeMismatch(String, left.Type())
	}
	if !right.Is(String) {
		return 0, ErrTypeMismatch(String, right.Type())
	}

	leftString := left.(StringValue).Value
	rightString := right.(StringValue).Value
	return strings.Compare(leftString, rightString), nil
}

func (StringTypeDescriptor) String() string { return "String" }

// Type returns a string type.
func (StringValue) Type() Type { return String }

// Is checks if this value is of type String.
func (StringValue) Is(t Type) bool { return t == String }

func (v StringValue) String() string { return v.Value }
