package types

import "bytes"

var (
	// Blob is the blob type. A Blob is a Binary Large OBject, and its base type
	// is a byte slice.
	Blob = BlobTypeDescriptor{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeBinary,
		},
	}
)

var _ Type = (*BlobTypeDescriptor)(nil)
var _ Value = (*BlobValue)(nil)

type (
	// BlobTypeDescriptor is the type descriptor for Binary Large OBjects. The
	// value is basically a byte slice.
	BlobTypeDescriptor struct {
		genericTypeDescriptor
	}

	// BlobValue is a value of type Blob.
	BlobValue struct {
		// Value is the primitive value of this value object.
		Value []byte
	}
)

// Compare for the Blob is defined the lexicographical comparison between
// the primitive underlying values.
func (BlobTypeDescriptor) Compare(left, right Value) (int, error) {
	if !left.Is(Blob) {
		return 0, ErrTypeMismatch(Blob, left.Type())
	}
	if !right.Is(Blob) {
		return 0, ErrTypeMismatch(Blob, right.Type())
	}

	leftBlob := left.(BlobValue).Value
	rightBlob := right.(BlobValue).Value
	return bytes.Compare(leftBlob, rightBlob), nil
}

func (BlobTypeDescriptor) String() string { return "Blob" }

// Type returns a blob type.
func (BlobValue) Type() Type { return Blob }

// Is checks if this value is of type Blob.
func (BlobValue) Is(t Type) bool { return t == Blob }
