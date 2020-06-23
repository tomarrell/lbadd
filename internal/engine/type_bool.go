package engine

import "fmt"

var (
	boolType = BoolType{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeBool,
		},
	}
)

var _ Type = (*BoolType)(nil)
var _ Value = (*BoolValue)(nil)

type (
	// BoolType is the boolean type of this engine.
	BoolType struct {
		genericTypeDescriptor
	}

	// BoolValue is a value of type BoolType.
	BoolValue struct {
		// Value is the underlying primitive value.
		Value bool
	}
)

// Compare for the type BoolType is defined as false<true. The comparison result
// for the bool comparison is therefore: false~true==-1, true~false==1,
// true~true==false~false==0.
func (BoolType) Compare(left, right Value) (int, error) {
	leftBool := left.(BoolValue).Value
	rightBool := right.(BoolValue).Value
	if !leftBool && rightBool {
		return -1, nil
	} else if leftBool == rightBool {
		return 0, nil
	} else if leftBool && !rightBool {
		return 1, nil
	}
	return -2, fmt.Errorf("unhandled constellation: %v <-> %v", leftBool, rightBool)
}

// Type returns a bool type.
func (BoolValue) Type() Type { return boolType }
