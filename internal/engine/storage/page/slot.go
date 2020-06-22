package page

import "unsafe"

const (
	// SlotByteSize is the size of an Offset, in the Go memory layout as well as
	// in the serialized form.
	SlotByteSize = uint16(unsafe.Sizeof(Slot{})) // #nosec
)

// Slot represents a data slot in the page. A slot consists of an offset and a
// size.
type Slot struct {
	// Offset is the Offset of the data in the page data slice. If overflow page
	// support is added, this might need to be changed to an uint32.
	Offset uint16
	// Size is the length of the data segment in the page data slice. If
	// overflow page support is added, this might need to be changed to an
	// uint32.
	Size uint16
}

func (s Slot) encodeInto(target []byte) {
	/* very simple way to avoid a new 4 byte allocation, should probably also be
	applied to cells */
	_ = target[3]
	byteOrder.PutUint16(target[0:], s.Offset)
	byteOrder.PutUint16(target[2:], s.Size)
}

func decodeOffset(data []byte) Slot {
	_ = data[3]
	return Slot{
		Offset: byteOrder.Uint16(data[0:]),
		Size:   byteOrder.Uint16(data[2:]),
	}
}
