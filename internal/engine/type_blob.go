package engine

import "bytes"

var (
	blobType = BlobType{
		genericTypeDescriptor: genericTypeDescriptor{
			baseType: BaseTypeBinary,
		},
	}
)

var _ Type = (*BlobType)(nil)
var _ Value = (*BlobValue)(nil)

type (
	// BlobType is the type for Binary Large OBjects. The value is basically a
	// byte slice.
	BlobType struct {
		genericTypeDescriptor
	}

	// BlobValue is a value of type BlobType.
	BlobValue struct {
		// Value is the primitive value of this value object.
		Value []byte
	}
)

// Compare for the BlobType is defined the lexicographical comparison between
// the primitive underlying values.
func (BlobType) Compare(left, right Value) (int, error) {
	leftBlob := left.(BlobValue).Value
	rightBlob := right.(BlobValue).Value
	return bytes.Compare(leftBlob, rightBlob), nil
}

// Type returns a blob type.
func (BlobValue) Type() Type { return blobType }
