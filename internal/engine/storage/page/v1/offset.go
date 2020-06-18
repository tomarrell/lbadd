package v1

import "unsafe"

const (
	// OffsetSize is the size of an Offset, in the Go memory layout as well as
	// in the serialized form.
	OffsetSize = uint16(unsafe.Sizeof(Offset{})) // #nosec
)

// Offset represents a cell Offset in the page data.
type Offset struct {
	// Offset is the Offset of the data in the page data slice. If overflow page
	// support is added, this might need to be changed to an uint32.
	Offset uint16
	// Size is the length of the data segment in the page data slice. If
	// overflow page support is added, this might need to be changed to an
	// uint32.
	Size uint16
}

func (o Offset) encodeInto(target []byte) {
	/* very simple way to avoid a new 4 byte allocation, should probably also be
	applied to cells */
	_ = target[3]
	byteOrder.PutUint16(target[0:], o.Offset)
	byteOrder.PutUint16(target[2:], o.Size)
}

func decodeOffset(data []byte) Offset {
	_ = data[3]
	return Offset{
		Offset: byteOrder.Uint16(data[0:]),
		Size:   byteOrder.Uint16(data[2:]),
	}
}
