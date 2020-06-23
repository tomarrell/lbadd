package page

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"
)

const (
	// Size is the fix size of a page, which is 16KB or 16384 bytes.
	Size = 1 << 14
	// HeaderSize is the fix size of a page header, which is 10 bytes.
	HeaderSize = 6
)

// Header field offset in page data.
const (
	idOffset        = 0 // byte 1,2,3,4: byte page ID
	cellCountOffset = 4 // byte 5,6: cell count
)

var (
	byteOrder = binary.BigEndian
)

// ID is the type of a page ID. This is mainly to avoid any confusion.
// Changing this will break existing database files, so only change during major
// version upgrades.
type ID = uint32

// Page is a page implementation that does not support overflow pages. It is not
// meant for that. Since we want to separate index and data, records should not
// contain datasets, but rather enough information, to find the corresponding
// dataset in a data file.
type Page struct {
	// data is the underlying data byte slice, which holds the header, offsets
	// and cells.
	data []byte

	dirty bool
}

// New creates a new page with the given ID.
func New(id ID) *Page {
	data := make([]byte, Size)
	byteOrder.PutUint32(data[idOffset:], id)
	return &Page{
		data: data,
	}
}

// Load loads the given data into the page. The length of the given data byte
// slice may differ from v1.PageSize, however, it cannot exceed ^uint16(0)-1
// (65535 or 64KB), and must be larger than 22 (HeaderSize(=10) + 1 Offset(=4) +
// 1 empty cell(=8)).
func Load(data []byte) (*Page, error) {
	return load(data)
}

// ID returns the ID of this page. This value must be constant.
func (p *Page) ID() ID { return byteOrder.Uint32(p.data[idOffset:]) }

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
func (p *Page) StorePointerCell(cell PointerCell) error {
	return p.storePointerCell(cell)
}

// StoreRecordCell stores a record cell in this page. A record cell holds
// arbitrary, variable size data.
func (p *Page) StoreRecordCell(cell RecordCell) error {
	return p.storeRecordCell(cell)
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
	p.zero(offsetIndex, SlotByteSize)
	// delete cell data
	p.zero(cellOffset.Offset, cellOffset.Size)
	// close gap in offsets due to offset deletion
	from := offsetIndex + SlotByteSize // lower bound, right next to gap
	to := offsetIndex                  // upper bound of the offset data
	cellCount := p.CellCount()
	p.moveAndZero(from, HeaderSize+cellCount*SlotByteSize-from, to) // actually move the data
	// update cell count
	p.decrementCellCount(1)
	return true, nil
}

// Cell returns a cell from this page with the given key, or false if no such
// cell exists in this page. In that case, the returned page is also nil.
func (p *Page) Cell(key []byte) (CellTyper, bool) {
	_, _, cell, found := p.findCell(key)
	return cell, found
}

// Cells decodes all cells in this page, which can be expensive, and returns all
// of them. The returned cells do not point back to the original page data, so
// don't modify them. Instead, delete the old cell and store a new one.
func (p *Page) Cells() (result []CellTyper) {
	for _, offset := range p.OccupiedSlots() {
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

// OccupiedSlots returns all occupied slots in the page. The slots all point to
// cells in the page. The amount of slots will always be equal to the amount of
// cells stored in a page. The amount of slots in the page depends on the cell
// count of this page, not the other way around.
func (p *Page) OccupiedSlots() (result []Slot) {
	cellCount := p.CellCount()
	offsetsWidth := cellCount * SlotByteSize
	offsetData := p.data[HeaderSize : HeaderSize+offsetsWidth]
	for i := uint16(0); i < cellCount; i++ {
		result = append(result, decodeOffset(offsetData[i*SlotByteSize:i*SlotByteSize+SlotByteSize]))
	}
	return
}

// FreeSlots computes all free addressable cell slots in this page. The free
// slots are sorted in ascending order by the offset in the page.
func (p *Page) FreeSlots() (result []Slot) {
	offsets := p.OccupiedSlots()
	if len(offsets) == 0 {
		// if there are no offsets at all, that means that the page is empty,
		// and one slot is returned, which reaches from 0+OffsetSize until the
		// end of the page
		off := HeaderSize + SlotByteSize
		return []Slot{{
			Offset: HeaderSize + SlotByteSize,
			Size:   uint16(len(p.data)) - off,
		}}
	}

	sort.Slice(offsets, func(i, j int) bool {
		return offsets[i].Offset < offsets[j].Offset
	})
	// first slot, from end of offset data until first cell
	firstOff := HeaderSize + uint16(len(offsets)+1)*SlotByteSize // +1 because we always need space to store one more offset, so if that space is blocked, there is no free slot that is addressable
	firstSize := offsets[0].Offset - firstOff
	if firstSize > 0 {
		result = append(result, Slot{
			Offset: firstOff,
			Size:   firstSize,
		})
	}
	// rest of the spaces between cells
	for i := 0; i < len(offsets)-1; i++ {
		off := offsets[i].Offset + offsets[i].Size
		size := offsets[i+1].Offset - off
		if size > 0 {
			result = append(result, Slot{
				Offset: off,
				Size:   size,
			})
		}
	}
	return
}

// FindFreeSlotForSize searches for a free slot in this page, matching or
// exceeding the given data size. This is done by using a best-fit algorithm.
func (p *Page) FindFreeSlotForSize(dataSize uint16) (Slot, bool) {
	// sort all free slots by size
	slots := p.FreeSlots()
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].Size < slots[j].Size
	})
	// search for the best fitting slot, i.e. the first slot, whose size is greater
	// than or equal to the given data size
	index := sort.Search(len(slots), func(i int) bool {
		return slots[i].Size >= dataSize
	})
	if index == len(slots) {
		return Slot{}, false
	}
	return slots[index], true
}

// Defragment will move the cells in the page in a way that after defragmenting,
// there is only a single free block, and that is located between the offsets
// and the cell data. After calling this method, (*Page).Fragmentation() will
// return 0.
func (p *Page) Defragment() {
	occupied := p.OccupiedSlots()
	var newSlots []Slot
	nextLeftBound := uint16(len(p.data))
	for i := len(occupied) - 1; i >= 0; i-- {
		slot := occupied[i]
		newOffset := nextLeftBound - slot.Size
		p.moveAndZero(slot.Offset, slot.Size, newOffset)
		nextLeftBound = newOffset
		newSlots = append(newSlots, Slot{
			Offset: newOffset,
			Size:   slot.Size,
		})
	}
	// no need to sort new slots, as the order was not modified during
	// defragmentation, but we need to reverse it
	for i, j := 0, len(newSlots)-1; i < j; i, j = i+1, j-1 {
		newSlots[i], newSlots[j] = newSlots[j], newSlots[i]
	}

	for i, slot := range newSlots {
		slot.encodeInto(p.data[HeaderSize+uint16(i)*SlotByteSize:])
	}
}

// Fragmentation computes the page fragmentation, which is defined by 1 -
// (largest free block / total free size). Multiply with 100 to get the
// fragmentation percentage.
func (p *Page) Fragmentation() float32 {
	slots := p.FreeSlots()
	if len(slots) == 0 {
		return 0
	}

	var largestFree, totalFree uint16
	for _, slot := range slots {
		totalFree += slot.Size
		if slot.Size > largestFree {
			largestFree = slot.Size
		}
	}
	return 1 - (float32(largestFree) / float32(totalFree))
}

func load(data []byte) (*Page, error) {
	if len(data) > int(^uint16(0))-1 {
		return nil, fmt.Errorf("page size too large: %v (max %v)", len(data), int(^uint16(0))-1)
	}
	if len(data) < HeaderSize {
		return nil, fmt.Errorf("page size too small: %v (min %v)", len(data), HeaderSize)
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
func (p *Page) findCell(key []byte) (offsetIndex uint16, cellSlot Slot, cell CellTyper, found bool) {
	offsets := p.OccupiedSlots()
	result := sort.Search(len(offsets), func(i int) bool {
		cell := p.cellAt(offsets[i])
		switch c := cell.(type) {
		case RecordCell:
			return bytes.Compare(c.Key, key) >= 0
		case PointerCell:
			return bytes.Compare(c.Key, key) >= 0
		}
		return false
	})
	if result == len(offsets) {
		return 0, Slot{}, nil, false
	}
	return HeaderSize + uint16(result)*SlotByteSize, offsets[result], p.cellAt(offsets[result]), true
}

func (p *Page) storePointerCell(cell PointerCell) error {
	return p.storeRawCell(cell.Key, encodePointerCell(cell))
}

func (p *Page) storeRecordCell(cell RecordCell) error {
	return p.storeRawCell(cell.Key, encodeRecordCell(cell))
}

func (p *Page) storeRawCell(key, rawCell []byte) error {
	size := uint16(len(rawCell))
	slot, ok := p.FindFreeSlotForSize(size)
	if !ok {
		return ErrPageFull
	}
	p.storeCellSlot(Slot{
		Offset: slot.Offset + slot.Size - size,
		Size:   size,
	}, key)
	copy(p.data[slot.Offset+slot.Size-size:], rawCell)
	p.incrementCellCount(1)
	return nil
}

func (p *Page) storeCellSlot(offset Slot, cellKey []byte) {
	offsets := p.OccupiedSlots()
	if len(offsets) == 0 {
		// directly into the start of the page content, after the header
		offset.encodeInto(p.data[HeaderSize:])
		return
	}

	index := sort.Search(len(offsets), func(i int) bool {
		cell := p.cellAt(offsets[i])
		switch c := cell.(type) {
		case RecordCell:
			return bytes.Compare(cellKey, c.Key) < 0
		case PointerCell:
			return bytes.Compare(cellKey, c.Key) < 0
		}
		return false
	})

	offsetOffset := HeaderSize + uint16(index)*SlotByteSize
	if index != len(offsets) {
		// make room if neccessary
		allOffsetsEnd := HeaderSize + uint16(len(offsets))*SlotByteSize
		p.moveAndZero(offsetOffset, allOffsetsEnd-offsetOffset, offsetOffset+SlotByteSize)
	}
	offset.encodeInto(p.data[offsetOffset:])
}

func (p *Page) cellAt(slot Slot) CellTyper {
	return decodeCell(p.data[slot.Offset : slot.Offset+slot.Size])
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
	if target == offset {
		// no-op when offset and target are the same
		return
	}

	_ = p.data[offset+size-1] // bounds check
	_ = p.data[target+size-1] // bounds check

	copy(p.data[target:target+size], p.data[offset:offset+size])

	// area needs zeroing
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
