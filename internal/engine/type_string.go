package engine

import "strings"

var (
	stringType = StringType{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeString,
		},
	}
)

var _ Type = (*StringType)(nil)
var _ Value = (*StringValue)(nil)

type (
	// StringType is the type for parameterized and non-parameterized string
	// types, such as VARCHAR or VARCHAR(n).
	StringType struct {
		genericTypeDescriptor
	}

	// StringValue is a value of type StringType.
	StringValue struct {
		// Value is the underlying primitive value.
		Value string
	}
)

// Compare for the StringType is defined as the lexicographical comparison of
// the two underlying primitive values.
func (StringType) Compare(left, right Value) (int, error) {
	leftString := left.(StringValue).Value
	rightString := right.(StringValue).Value
	return strings.Compare(leftString, rightString), nil
}

// Type returns a string type.
func (StringValue) Type() Type { return blobType }
