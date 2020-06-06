package v1

import "unsafe"

const (
	// OffsetSize is the constant size of an encoded Offset using encodeOffset.
	OffsetSize = uint16(unsafe.Sizeof(Offset{})) // #nosec
)

// Offset describes a memory segment relative to the page start (=0, =before the
// header).
type Offset struct {
	// Location is the target location of this offset, relative to the page.
	Location uint16
	// Size is the size of the memory segment that is located at Location.
	Size uint16
}
