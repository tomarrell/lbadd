package v1

import (
	"bytes"
	"sort"
	"unsafe"

	"github.com/tomarrell/lbadd/internal/engine/converter"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

// Internal headers, defined from 255 downwards, as opposed to other headers,
// which are defined from 0 upwards.
const (
	InternalHeaderCellCount page.Header = page.Header(^uint8(0)) - iota
	InternalHeaderDirty
)

// Constants.
const (
	// PageSize is the fixed size of one page.
	PageSize uint16 = 1 << 14
	// HeaderSize is the fixed size of the header area in a page.
	HeaderSize uint16 = 1 << 9
	// HeaderOffset is the offset in the page's data, at which the header
	// starts.
	HeaderOffset uint16 = 0
	// ContentOffset is the offset in the page's data, at which the header ends
	// and the content starts. This is also, where the offsets are stored.
	ContentOffset uint16 = HeaderOffset + HeaderSize
	// ContentSize is the size of the content area, equivalent to the page size
	// minus the header size.
	ContentSize uint16 = PageSize - HeaderSize

	// FrameSizeSize is the byte size of a frame.
	FrameSizeSize = uint16(unsafe.Sizeof(uint16(0))) // #nosec

	HeaderVersionOffset = HeaderOffset
	HeaderVersionSize   = uint16(unsafe.Sizeof(uint16(0))) // #nosec
	HeaderVersionLo     = HeaderVersionOffset
	HeaderVersionHi     = HeaderVersionOffset + HeaderVersionSize

	HeaderIDOffset = HeaderVersionOffset + HeaderVersionSize
	HeaderIDSize   = uint16(unsafe.Sizeof(uint32(0))) // #nosec
	HeaderIDLo     = HeaderIDOffset
	HeaderIDHi     = HeaderIDOffset + HeaderIDSize

	InternalHeaderCellCountOffset = HeaderIDOffset + HeaderIDSize
	InternalHeaderCellCountSize   = uint16(unsafe.Sizeof(uint16(0))) // #nosec
	InternalHeaderCellCountLo     = InternalHeaderCellCountOffset
	InternalHeaderCellCountHi     = InternalHeaderCellCountOffset + InternalHeaderCellCountSize

	InternalHeaderDirtyOffset = InternalHeaderCellCountOffset + InternalHeaderCellCountSize
	InternalHeaderDirtySize   = uint16(unsafe.Sizeof(false)) // #nosec
	InternalHeaderDirtyIndex  = InternalHeaderDirtyOffset
)

var _ page.Loader = Load
var _ page.Page = (*Page)(nil)

// Page is an implementation of (page.Page). It implements the concept of
// slotted pages, also it holds all data in memory at all times. The byte slice
// that is passed in when loading the page, is the one that the implementation
// will operate on. It will always stick to that slice, and never replace it
// with another. All changes to the page are immediately reflected in that byte
// slice. This implementation is NOT safe for concurrent use. This
// implementation does not support extension pages.
type Page struct {
	data []byte

	header map[page.Header]interface{}
}

// Load loads the given bytes as a page. The page will operate on the given byte
// slice, and never copy or exchange it. Modifying the byte slice after loading
// a page from it will likely corrupt the page. External changes to the byte
// slice are not guaranteed to be immediately reflected in the page object.
func Load(data []byte) (page.Page, error) {
	return load(data)
}

// Header obtains the header field value of the given key. If the key is not
// supported by this implementation, it will return nil indicating an unknown
// header.
func (p *Page) Header(key page.Header) interface{} {
	val, ok := p.header[key]
	if !ok {
		return nil
	}
	return val
}

// Dirty returns the value of the header flag dirty, meaning that the page has
// been modified since ClearDirty() has been called the last time.
func (p *Page) Dirty() bool {
	return p.header[InternalHeaderDirty].(bool)
}

// MarkDirty marks this page as dirty. Call this when you have modified the
// page. This will
func (p *Page) MarkDirty() {
	p.data[InternalHeaderDirtyIndex] = converter.BoolToByte(true)
	p.header[InternalHeaderDirty] = true
}

// ClearDirty marks this page as NOT dirty (anymore). Call this when the page
// has been written to persistent storage.
func (p *Page) ClearDirty() {
	p.data[InternalHeaderDirtyIndex] = converter.BoolToByte(false)
	p.header[InternalHeaderDirty] = false
}

// StoreCell will store the given cell in this page. If there is not enough
// space, NO extension page will be allocated, but an error will be returned,
// indicating insufficient space.
func (p *Page) StoreCell(cell page.Cell) error {
	cellData := encodeCell(cell)
	insertionOffset, ok := p.findInsertionOffset(uint16(len(cellData)))
	if !ok {
		return page.ErrPageTooSmall
	}
	offsetInsertionOffset := p.findOffsetInsertionOffset(cell)
	p.insertOffset(insertionOffset, offsetInsertionOffset)
	copy(p.data[ContentOffset+insertionOffset.Location:], cellData)

	p.incrementCellCount()

	return nil
}

// Delete deletes the cell with the given key. If there is no such cell, this is
// a no-op. This never returns an error.
func (p *Page) Delete(key []byte) error {
	offsets := p.Offsets()

	for i, offset := range offsets {
		if bytes.Equal(key, p.getCell(offset).Key) {
			offsetData := p.data[ContentOffset : ContentOffset+(p.cellCount()*OffsetSize)]
			zero(offsetData)
			offsets = append(offsets[:i], offsets[i+1:]...)
			for j, off := range offsets {
				lo := uint16(j) * OffsetSize
				hi := lo + OffsetSize
				copy(offsetData[lo:hi], encodeOffset(off))
			}
			break
		}
	}

	p.decrementCellCount()

	return nil
}

// Cell searches for a cell with the given key, and returns a cell object
// representing all the found cell data. If no cell was found, false is
// returned.
func (p *Page) Cell(key []byte) (page.Cell, bool) {
	// TODO: binary search
	for _, cell := range p.Cells() {
		if bytes.Equal(key, cell.Key) {
			return cell, true
		}
	}
	return page.Cell{}, false
}

// Cells returns all cells that are stored in this page in sorted fashion,
// ordered ascending by key.
func (p *Page) Cells() (cells []page.Cell) {
	for _, offset := range p.Offsets() {
		cells = append(cells, p.getCell(offset))
	}
	return
}

// Offsets returns all offsets of this page sorted by key of the cell that they
// point to. Following the offsets in the order that they are returned, will
// result in a list of cells, that are sorted ascending by key.
func (p *Page) Offsets() (result []Offset) {
	cellCount := p.cellCount()
	offsetData := p.data[ContentOffset : ContentOffset+OffsetSize*cellCount]
	for i := 0; i < len(offsetData); i += int(OffsetSize) {
		result = append(result, decodeOffset(offsetData[i:i+int(OffsetSize)]))
	}
	return
}

func load(data []byte) (*Page, error) {
	if len(data) != int(PageSize) {
		return nil, page.ErrInvalidPageSize
	}
	p := &Page{
		data:   data,
		header: make(map[page.Header]interface{}),
	}
	p.loadHeader()
	return p, nil
}

// insertOffset inserts the given offset at the other given offset. A bit
// confusing, because we need to store an offset, and the location is given by
// another offset.
//
// This method takes care of inserting the offset, and moving other offsets to
// the right, instead of overwriting them.
func (p *Page) insertOffset(offset, at Offset) {
	cellCount := p.cellCount()
	offsetData := p.data[ContentOffset : ContentOffset+OffsetSize*(cellCount+1)]
	encOffset := encodeOffset(offset)
	buf := make([]byte, uint16(len(offsetData))-OffsetSize-at.Location)
	copy(buf, offsetData[at.Location:])
	copy(offsetData[at.Location+OffsetSize:], buf)
	copy(offsetData[at.Location:], encOffset)
}

// getCell returns the cell that an offset points to.
func (p *Page) getCell(offset Offset) page.Cell {
	return decodeCell(p.data[ContentOffset+offset.Location : ContentOffset+offset.Location+offset.Size])
}

// findInsertionOffset finds an offset for a data segment of the given size, or
// returns false if no space is available. The space is found by using a
// first-fit approach.
func (p *Page) findInsertionOffset(size uint16) (Offset, bool) {
	offsets := p.Offsets()
	sort.Slice(offsets, func(i int, j int) bool {
		return offsets[i].Location < offsets[j].Location
	})

	// TODO: best fit, this is currently first fit
	rightBound := ContentSize
	for i := len(offsets) - 1; i >= 0; i-- {
		current := offsets[i]
		if current.Location+current.Size >= rightBound-size {
			// doesn't fit
			rightBound = current.Location
		} else {
			break
		}
	}

	return Offset{
		Location: rightBound - size,
		Size:     size,
	}, true
}

// findOffsetInsertionOffset creates an offset to a location, where a new offset
// can be inserted. The new offset that should be inserted, is not part of this
// method. However, we need the cell that the new offset points to, in order to
// compare its key with other cells. This is, because we insert offsets in a
// way, so that all offsets from left to right point to cells, that are ordered
// ascending by key when following all offsets.
func (p *Page) findOffsetInsertionOffset(cell page.Cell) Offset {
	// TODO: binary search
	offsets := p.Offsets()
	for i, offset := range offsets {
		if bytes.Compare(p.getCell(offset).Key, cell.Key) > 0 {
			return Offset{
				Location: uint16(i) * OffsetSize,
				Size:     OffsetSize,
			}
		}
	}
	return Offset{
		Location: uint16(len(offsets)) * OffsetSize,
		Size:     OffsetSize,
	}
}

// loadHeader (re-)loads all header values from the page's data.
func (p *Page) loadHeader() {
	p.header[page.HeaderVersion] = converter.ByteSliceToUint16(p.data[HeaderVersionLo:HeaderVersionHi])
	p.header[page.HeaderID] = converter.ByteSliceToUint16(p.data[HeaderIDLo:HeaderIDHi])
	p.header[InternalHeaderCellCount] = converter.ByteSliceToUint16(p.data[InternalHeaderCellCountLo:InternalHeaderCellCountHi])
}

func (p *Page) setInternalHeaderCellCount(count uint16) {
	// set in memory cache
	p.header[InternalHeaderCellCount] = count
	// set in data
	copy(p.data[InternalHeaderCellCountLo:InternalHeaderCellCountHi], converter.Uint16ToByteSlice(count))
}

// incrementCellCount increments the header field InternalHeaderCellCount by one.
func (p *Page) incrementCellCount() {
	p.setInternalHeaderCellCount(p.header[InternalHeaderCellCount].(uint16) + 1)
}

// decrementCellCount decrements the header field InternalHeaderCellCount by one.
func (p *Page) decrementCellCount() {
	p.setInternalHeaderCellCount(p.header[InternalHeaderCellCount].(uint16) - 1)
}

// cellCount returns the amount of currently stored cells. This value is
// retrieved from the header field InternalHeaderCellCount.
func (p *Page) cellCount() uint16 {
	return p.header[InternalHeaderCellCount].(uint16)
}
