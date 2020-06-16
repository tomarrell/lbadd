package v1

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

var _ page.Page = (*Page)(nil)
var _ page.Loader = Load

const (
	// PageSize is the fix size of a page, which is 16KB or 16384 bytes.
	PageSize = 1 << 14
	// HeaderSize is the fix size of a page header, which is 10 bytes.
	HeaderSize = 10
)

// Header field offset in page data.
const (
	versionOffset   = 0 // byte 1,2,3,4: version
	idOffset        = 4 // byte 5,6,7,8: byte page ID
	cellCountOffset = 8 // byte 9,10: cell count
)

var (
	byteOrder = binary.BigEndian
)

// Page is a page implementation that does not support overflow pages. It is not
// meant for that. Since we want to separate index and data into separate files,
// records should not contain datasets, but rather enough information, to find
// the corresponding dataset in a data file.
type Page struct {
	// data is the underlying data byte slice, which holds the header, offsets
	// and cells.
	data []byte

	dirty bool
}

// Load loads the given data into the page. The length of the given data byte
// slice may differ from v1.PageSize, however, it cannot exceed ^uint16(0)-1
// (65535 or 64KB), and must be larger than 22 (HeaderSize(=10) + 1 Offset(=4) +
// 1 empty cell(=8)).
func Load(data []byte) (page.Page, error) {
	return load(data)
}

// Version returns the version of this page. This should always be 1. This value
// must be constant.
func (p *Page) Version() uint32 { return byteOrder.Uint32(p.data[versionOffset:]) }

// ID returns the ID of this page. This value must be constant.
func (p *Page) ID() page.ID { return byteOrder.Uint32(p.data[idOffset:]) }

// CellCount returns the amount of stored cells in this page. This value is NOT
// constant.
func (p *Page) CellCount() uint16 { return byteOrder.Uint16(p.data[cellCountOffset:]) }

// Dirty returns whether the page is dirty (needs syncing with secondary
// storage).
func (p *Page) Dirty() bool { return p.dirty }

// MarkDirty marks this page as dirty.
func (p *Page) MarkDirty() { p.dirty = true }

// ClearDirty marks this page as NOT dirty.
func (p *Page) ClearDirty() { p.dirty = false }

// StorePointerCell stores a pointer cell in this page. A pointer cell points to
// other page IDs.
func (p *Page) StorePointerCell(cell page.PointerCell) error {
	if v1cell, ok := cell.(PointerCell); ok {
		return p.storePointerCell(v1cell)
	}
	return fmt.Errorf("can only store v1 pointer cells, but got %T", cell)
}

// StoreRecordCell stores a record cell in this page. A record cell holds
// arbitrary, variable size data.
func (p *Page) StoreRecordCell(cell page.RecordCell) error {
	if v1cell, ok := cell.(RecordCell); ok {
		return p.storeRecordCell(v1cell)
	}
	return fmt.Errorf("can only store v1 record cells, but got %T", cell)
}

// DeleteCell deletes a cell with the given key. If no such cell could be found,
// false is returned. In this implementation, an error can not occur while
// deleting a cell.
func (p *Page) DeleteCell(key []byte) (bool, error) {
	offsetIndex, cellOffset, _, found := p.findCell(key)
	if !found {
		return false, nil
	}

	// delete offset
	p.zero(offsetIndex*OffsetSize, OffsetSize)
	// delete cell data
	p.zero(cellOffset.Offset, cellOffset.Size)
	// close gap in offsets due to offset deletion
	from := offsetIndex*OffsetSize + OffsetSize   // lower bound, right next to gap
	to := p.CellCount() * OffsetSize              // upper bound of the offset data
	p.moveAndZero(from, to-from, from-OffsetSize) // actually move the data
	// update cell count
	p.decrementCellCount(1)
	return true, nil
}

// Cell returns a cell from this page with the given key, or false if no such
// cell exists in this page. In that case, the returned page is also nil.
func (p *Page) Cell(key []byte) (page.Cell, bool) {
	_, _, cell, found := p.findCell(key)
	return cell, found
}

// Cells decodes all cells in this page, which can be expensive, and returns all
// of them. The returned cells do not point back to the original page data, so
// don't modify them. Instead, delete the old cell and store a new one.
func (p *Page) Cells() (result []page.Cell) {
	for _, offset := range p.Offsets() {
		result = append(result, decodeCell(p.data[offset.Offset:offset.Offset+offset.Size]))
	}
	return
}

// RawData returns a copy of the page's internal data, so you can modify it at
// will, and it won't change the original page data.
func (p *Page) RawData() []byte {
	cp := make([]byte, len(p.data))
	copy(cp, p.data)
	return cp
}

// Offsets returns all offsets in the page. The offsets can be used to find all
// cells in the page. The amount of offsets will always be equal to the amount
// of cells stored in a page. The amount of offsets in the page depends on the
// cell count of this page, not the other way around.
func (p *Page) Offsets() (result []Offset) {
	cellCount := p.CellCount()
	offsetsWidth := cellCount * OffsetSize
	offsetData := p.data[HeaderSize : HeaderSize+offsetsWidth]
	for i := uint16(0); i < cellCount; i++ {
		result = append(result, decodeOffset(offsetData[i*OffsetSize:i*OffsetSize+OffsetSize]))
	}
	return
}

func load(data []byte) (*Page, error) {
	if len(data) > int(^uint16(0))-1 {
		return nil, fmt.Errorf("page size too large: %v (max %v)", len(data), int(^uint16(0))-1)
	}

	return &Page{
		data: data,
	}, nil
}

// findCell searches for a cell with the given key, as well as the corresponding
// offset and the corresponding offset index. The index is the index of the cell
// offset in all offsets, meaning that the byte location of the offset in the
// page can be obtained with offsetIndex*OffsetSize. The cellOffset is the
// offset that points to the cell. cell is the cell that was found, or nil if no
// cell with the given key could be found. If no cell could be found,
// found=false will be returned, as well as zero values for all other return
// arguments.
func (p *Page) findCell(key []byte) (offsetIndex uint16, cellOffset Offset, cell Cell, found bool) {
	offsets := p.Offsets()
	result := sort.Search(len(offsets), func(i int) bool {
		cell := p.cellAt(offsets[i])
		return bytes.Compare(cell.Key(), key) >= 0
	})
	if result == len(offsets) {
		return 0, Offset{}, nil, false
	}
	return uint16(result), offsets[result], p.cellAt(offsets[result]), true
}

func (p *Page) storePointerCell(cell PointerCell) error {
	return p.storeRawCell(encodePointerCell(cell))
}

func (p *Page) storeRecordCell(cell RecordCell) error {
	return p.storeRawCell(encodeRecordCell(cell))
}

func (p *Page) storeRawCell(rawCell []byte) error {
	p.incrementCellCount(1)
	_ = Offset{}.encodeInto // to remove linter error
	return fmt.Errorf("unimplemented")
}

func (p *Page) cellAt(offset Offset) Cell {
	return decodeCell(p.data[offset.Offset : offset.Offset+offset.Size])
}

// moveAndZero moves target bytes in the page's raw data from offset to target,
// and zeros all bytes from offset to offset+size, that do not overlap with the
// target area. Source and target area may overlap.
//
//  [1,1,2,2,2,1,1,1,1,1]
//  moveAndZero(2, 3, 6)
//  [1,1,0,0,0,1,2,2,2,1]
//
// or, with overlap
//
//  [1,1,2,2,2,1,1,1,1,1]
//  moveAndZero(2, 3, 4)
//  [1,1,0,0,2,2,2,1,1,1]
func (p *Page) moveAndZero(offset, size, target uint16) {
	_ = p.data[offset+size-1] // bounds check
	_ = p.data[target+size-1] // bounds check

	copy(p.data[target:target+size], p.data[offset:offset+size])

	// area needs zeroing
	if target != offset {
		if target > offset+size || target+size < offset {
			// no overlap
			p.zero(offset, size)
		} else {
			// overlap
			if target > offset && target <= offset+size {
				// move to right, zero non-overlapping area
				p.zero(offset, target-offset)
			} else if target < offset && target+size >= offset {
				// move to left, zero non-overlapping area
				p.zero(target+size, offset-target)
			}
		}
	}
}

// zero zeroes size bytes, starting at offset in the page's raw data.
func (p *Page) zero(offset, size uint16) {
	for i := uint16(0); i < size; i++ {
		p.data[offset+i] = 0
	}
}

func (p *Page) incrementCellCount(delta uint16) { p.incrementUint16(cellCountOffset, delta) }
func (p *Page) decrementCellCount(delta uint16) { p.decrementUint16(cellCountOffset, delta) }

func (p *Page) storeUint16(at, val uint16)  { byteOrder.PutUint16(p.data[at:], val) }
func (p *Page) loadUint16(at uint16) uint16 { return byteOrder.Uint16(p.data[at:]) }

func (p *Page) incrementUint16(at, delta uint16) { p.storeUint16(at, p.loadUint16(at)+delta) }
func (p *Page) decrementUint16(at, delta uint16) { p.storeUint16(at, p.loadUint16(at)-delta) }
