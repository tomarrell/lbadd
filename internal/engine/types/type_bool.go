package types

import (
	"fmt"
	"strconv"
)

var (
	// Bool is the boolean type. Its base type is a bool.
	Bool = BoolTypeDescriptor{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeBool,
		},
	}
)

var _ Type = (*BoolTypeDescriptor)(nil)
var _ Value = (*BoolValue)(nil)

type (
	// BoolTypeDescriptor is the boolean type of this engine.
	BoolTypeDescriptor struct {
		genericTypeDescriptor
	}

	// BoolValue is a value of type Bool.
	BoolValue struct {
		// Value is the underlying primitive value.
		Value bool
	}
)

// Compare for the type BoolType is defined as false<true. The comparison result
// for the bool comparison is therefore: false~true==-1, true~false==1,
// true~true==false~false==0.
func (BoolTypeDescriptor) Compare(left, right Value) (int, error) {
	if !left.Is(Bool) {
		return 0, ErrTypeMismatch(Bool, left.Type())
	}
	if !right.Is(Bool) {
		return 0, ErrTypeMismatch(Bool, right.Type())
	}

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

func (BoolTypeDescriptor) String() string { return "Bool" }

// Type returns a bool type.
func (BoolValue) Type() Type { return Bool }

// Is checks if this value is of type Bool.
func (BoolValue) Is(t Type) bool { return t == Bool }

func (v BoolValue) String() string { return strconv.FormatBool(v.Value) }
