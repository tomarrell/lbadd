package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage_StoreRecordCell(t *testing.T) {
	assert := assert.New(t)

	p, err := load(make([]byte, 32))
	assert.NoError(err)

	c := RecordCell{
		Key:    []byte{0xAB},
		Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	}

	err = p.StoreRecordCell(c)
	assert.NoError(err)
	assert.Equal([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // header
		0x00, 0x12, 0x00, 0x0E, // offset
		0x00, 0x00, 0x00, 0x00, // reserved for next offset
		0x00, 0x00, 0x00, 0x00, // free slot #0
		0x01,                   // cell type
		0x00, 0x00, 0x00, 0x01, // key frame
		0xAB,                   // key
		0x00, 0x00, 0x00, 0x04, // record frame
		0xCA, 0xFE, 0xBA, 0xBE, // record
	}, p.data)

	freeSlots := p.FreeSlots()
	assert.Len(freeSlots, 1)
	// offset must skipt reserved space for offset, as the offset is not free
	// space
	assert.Equal(Slot{
		Offset: 14,
		Size:   4,
	}, freeSlots[0])

	pageData := make([]byte, len(p.data))
	copy(pageData, p.data)

	anotherCell := RecordCell{
		Key:    []byte("large key"),
		Record: []byte("way too large record"),
	}
	err = p.StoreRecordCell(anotherCell)
	assert.Equal(ErrPageFull, err)
	assert.Equal(pageData, p.data) // page must not have been modified
}

func TestPage_StoreRecordCell_Multiple(t *testing.T) {
	assert := assert.New(t)

	p, err := load(make([]byte, 60))
	assert.NoError(err)

	cells := []RecordCell{
		{
			Key:    []byte{0x11},
			Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
		},
		{
			Key:    []byte{0x33},
			Record: []byte{0xD1, 0xCE},
		},
		{
			Key:    []byte{0x22},
			Record: []byte{0xFF},
		},
	}
	assert.NoError(p.storeRecordCell(cells[0]))
	assert.NoError(p.storeRecordCell(cells[1]))
	assert.NoError(p.storeRecordCell(cells[2]))
	assert.Equal([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, // header
		0x00, 0x2E, 0x00, 0x0E, // offset #0
		0x00, 0x17, 0x00, 0x0B, // offset #2
		0x00, 0x22, 0x00, 0x0C, // offset #1
		0x00, 0x00, 0x00, 0x00, 0x00, // free space
		// cell #3
		0x01,                   // cell #3 type
		0x00, 0x00, 0x00, 0x01, // cell #3 key frame
		0x22,                   // cell #3 key
		0x00, 0x00, 0x00, 0x01, // cell #3 record frame
		0xFF, // cell #3 record
		// cell #2
		0x01,                   // cell #2 type
		0x00, 0x00, 0x00, 0x01, // cell #2 key frame
		0x33,                   // cell #2 key
		0x00, 0x00, 0x00, 0x02, // cell #2 record frame
		0xD1, 0xCE, // cell #2 record
		// cell #1
		0x01,                   // cell #1 type
		0x00, 0x00, 0x00, 0x01, // cell #1 key frame
		0x11,                   // cell #1 key
		0x00, 0x00, 0x00, 0x04, // cell #1 record frame
		0xCA, 0xFE, 0xBA, 0xBE, // cell #1 record
	}, p.data)
}

func TestPage_OccupiedSlots(t *testing.T) {
	assert := assert.New(t)

	// create a completely empty page
	p, err := load(make([]byte, Size))
	assert.NoError(err)

	// create the offset source data
	offsetCount := 3
	offsetData := []byte{
		// offset[0]
		0x01, 0x12, // offset
		0x23, 0x34, // size
		// offset[1]
		0x45, 0x56, // offset
		0x67, 0x78, // size
		// offset[2]
		0x89, 0x9A, // offset
		0xAB, 0xBC, // size
	}
	// quick check if we made a mistake in the test
	assert.EqualValues(SlotByteSize, len(offsetData)/offsetCount)

	// inject the offset data
	p.incrementCellCount(3)               // set the cell count
	copy(p.data[HeaderSize:], offsetData) // copy the offset data

	// actual test can start

	offsets := p.OccupiedSlots()
	assert.Len(offsets, 3)
	assert.Equal(Slot{
		Offset: 0x0112,
		Size:   0x2334,
	}, offsets[0])
	assert.Equal(Slot{
		Offset: 0x4556,
		Size:   0x6778,
	}, offsets[1])
	assert.Equal(Slot{
		Offset: 0x899A,
		Size:   0xABBC,
	}, offsets[2])
}

func TestPage_moveAndZero(t *testing.T) {
	type args struct {
		offset uint16
		size   uint16
		target uint16
	}
	tests := []struct {
		name string
		data []byte
		args args
		want []byte
	}{
		{
			"same position",
			[]byte{1, 1, 2, 2, 2, 2, 1, 1, 1, 1},
			args{2, 4, 2},
			[]byte{1, 1, 2, 2, 2, 2, 1, 1, 1, 1},
		},
		{
			"single no overlap to right",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{0, 1, 2},
			[]byte{0, 2, 1, 4, 5, 6, 7, 8, 9},
		},
		{
			"double no overlap to right",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{0, 2, 3},
			[]byte{0, 0, 3, 1, 2, 6, 7, 8, 9},
		},
		{
			"many no overlap to right",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{0, 4, 5},
			[]byte{0, 0, 0, 0, 5, 1, 2, 3, 4},
		},
		{
			"single no overlap to left",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{8, 1, 2},
			[]byte{1, 2, 9, 4, 5, 6, 7, 8, 0},
		},
		{
			"double no overlap to left",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{7, 2, 3},
			[]byte{1, 2, 3, 8, 9, 6, 7, 0, 0},
		},
		{
			"many no overlap to left",
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			args{5, 4, 0},
			[]byte{6, 7, 8, 9, 5, 0, 0, 0, 0},
		},
		{
			"double 1 overlap to right",
			[]byte{1, 1, 2, 2, 1, 1, 1, 1, 1, 1},
			args{2, 2, 3},
			[]byte{1, 1, 0, 2, 2, 1, 1, 1, 1, 1},
		},
		{
			"double 1 overlap to left",
			[]byte{1, 1, 1, 2, 2, 1, 1, 1, 1, 1},
			args{3, 2, 2},
			[]byte{1, 1, 2, 2, 0, 1, 1, 1, 1, 1},
		},
		{
			"triple 1 overlap to right",
			[]byte{1, 1, 2, 2, 2, 1, 1, 1, 1, 1},
			args{2, 3, 4},
			[]byte{1, 1, 0, 0, 2, 2, 2, 1, 1, 1},
		},
		{
			"triple 2 overlap to right",
			[]byte{1, 1, 2, 2, 2, 1, 1, 1, 1, 1},
			args{2, 3, 3},
			[]byte{1, 1, 0, 2, 2, 2, 1, 1, 1, 1},
		},
		{
			"triple 1 overlap to left",
			[]byte{1, 1, 1, 1, 2, 2, 2, 1, 1, 1},
			args{4, 3, 2},
			[]byte{1, 1, 2, 2, 2, 0, 0, 1, 1, 1},
		},
		{
			"triple 2 overlap to left",
			[]byte{1, 1, 1, 2, 2, 2, 1, 1, 1, 1},
			args{3, 3, 2},
			[]byte{1, 1, 2, 2, 2, 0, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			p := &Page{
				data: tt.data,
			}
			p.moveAndZero(tt.args.offset, tt.args.size, tt.args.target)
			assert.Equal(tt.want, p.data)
		})
	}
}

func TestPage_FindFreeSlotForSize(t *testing.T) {
	assert := assert.New(t)

	p, err := load(make([]byte, 100))
	assert.NoError(err)

	occupiedSlots := []Slot{
		{90, 10},
		// 1 byte
		{80, 9},
		// 25 bytes
		{50, 5},
		// 10 bytes
		{30, 10},
	}

	for i, slot := range occupiedSlots {
		slot.encodeInto(p.data[HeaderSize+i*int(SlotByteSize):])
	}
	p.incrementCellCount(uint16(len(occupiedSlots)))

	slot, ok := p.FindFreeSlotForSize(1)
	assert.True(ok)
	assert.Equal(Slot{89, 1}, slot)

	slot, ok = p.FindFreeSlotForSize(15)
	assert.True(ok)
	assert.Equal(Slot{55, 25}, slot)

	slot, ok = p.FindFreeSlotForSize(25)
	assert.True(ok)
	assert.Equal(Slot{55, 25}, slot)

	slot, ok = p.FindFreeSlotForSize(10)
	assert.True(ok)
	assert.Equal(Slot{40, 10}, slot)

	slot, ok = p.FindFreeSlotForSize(5)
	assert.True(ok)
	assert.Equal(Slot{40, 10}, slot)
}

func TestPage_FreeSlots(t *testing.T) {
	assert := assert.New(t)

	p, err := load(make([]byte, 100))
	assert.NoError(err)

	occupiedSlots := []Slot{
		// 2 bytes
		{28, 8},
		// 10 bytes
		{46, 5},
		// 25 bytes
		{76, 9},
		// 1 byte
		{86, 10},
	}

	for i, slot := range occupiedSlots {
		slot.encodeInto(p.data[HeaderSize+i*int(SlotByteSize):])
	}
	p.incrementCellCount(uint16(len(occupiedSlots)))

	assert.EqualValues([]Slot{
		{26, 2},
		{36, 10},
		{51, 25},
		{85, 1},
	}, p.FreeSlots())
}

func TestPage_Defragment(t *testing.T) {
	tests := []struct {
		name   string
		before []byte
		after  []byte
	}{
		{
			"small 2 cells",
			[]byte{
				/* 0x00 */ 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				/* 0x06 */ 0x00, 0x12, 0x00, 0x04, // offset #0
				/* 0x0a */ 0x00, 0x1A, 0x00, 0x04, // offset #1
				/* 0x0e */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x12 */ 0x01, 0x01, 0x01, 0x01, // cell #0
				/* 0x16 */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x1a */ 0x02, 0x02, 0x02, 0x02, // cell #1
			},
			[]byte{
				/* 0x00 */ 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				/* 0x06 */ 0x00, 0x16, 0x00, 0x04, // offset #0
				/* 0x0a */ 0x00, 0x1A, 0x00, 0x04, // offset #1
				/* 0x0e */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x12 */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x16 */ 0x01, 0x01, 0x01, 0x01, // cell #0
				/* 0x1a */ 0x02, 0x02, 0x02, 0x02, // cell #1
			},
		},
		{
			"small 2 cells free end",
			[]byte{
				/* 0x00 */ 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				/* 0x06 */ 0x00, 0x12, 0x00, 0x04, // offset #0
				/* 0x0a */ 0x00, 0x1A, 0x00, 0x04, // offset #1
				/* 0x0e */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x12 */ 0x01, 0x01, 0x01, 0x01, // cell #0
				/* 0x16 */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x1a */ 0x02, 0x02, 0x02, 0x02, // cell #1
				/* 0x1e */ 0x00, 0x00, 0x00, 0x00, // free space
			},
			[]byte{
				/* 0x00 */ 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				/* 0x06 */ 0x00, 0x1A, 0x00, 0x04, // offset #0
				/* 0x0a */ 0x00, 0x1E, 0x00, 0x04, // offset #1
				/* 0x0e */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x12 */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x16 */ 0x00, 0x00, 0x00, 0x00, // free space
				/* 0x1a */ 0x01, 0x01, 0x01, 0x01, // cell #0
				/* 0x1e */ 0x02, 0x02, 0x02, 0x02, // cell #1
			},
		},
		{
			"full page",
			[]byte{
				/* 0x00 */ 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				/* 0x06 */ 0x00, 0x0E, 0x00, 0x04, // offset #0
				/* 0x0a */ 0x00, 0x12, 0x00, 0x04, // offset #1
				/* 0x0e */ 0x01, 0x01, 0x01, 0x01, // cell #0
				/* 0x12 */ 0x02, 0x02, 0x02, 0x02, // cell #1
			},
			[]byte{
				/* 0x00 */ 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				/* 0x06 */ 0x00, 0x0E, 0x00, 0x04, // offset #0
				/* 0x0a */ 0x00, 0x12, 0x00, 0x04, // offset #1
				/* 0x0e */ 0x01, 0x01, 0x01, 0x01, // cell #0
				/* 0x12 */ 0x02, 0x02, 0x02, 0x02, // cell #1
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			p, err := load(tt.before)
			assert.NoError(err)

			assert.Equal(tt.before, p.data)
			p.Defragment()
			assert.Equal(tt.after, p.data)
		})
	}
}

func TestPage_DeleteCell(t *testing.T) {
	tests := []struct {
		name      string
		before    []byte
		deleteKey []byte
		after     []byte
		ok        bool
	}{
		{
			"small 2 cells delete front",
			[]byte{
				0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				0x00, 0x12, 0x00, 0x0A, // offset #0
				0x00, 0x20, 0x00, 0x0A, // offset #1
				0x00, 0x00, 0x00, 0x00, // free space
				0x01, 0x00, 0x00, 0x00, 0x01, 0x0A, 0x00, 0x00, 0x00, 0x00, // cell #0
				0x00, 0x00, 0x00, 0x00, // free space
				0x01, 0x00, 0x00, 0x00, 0x01, 0x0B, 0x00, 0x00, 0x00, 0x00, // cell #1
			},
			[]byte{0x0A},
			[]byte{
				0x00, 0x00, 0x00, 0x01, 0x00, 0x01, // header
				0x00, 0x20, 0x00, 0x0A, // offset #1
				0x00, 0x00, 0x00, 0x00, // free space
				0x00, 0x00, 0x00, 0x00, // free space
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // free space
				0x00, 0x00, 0x00, 0x00, // free space
				0x01, 0x00, 0x00, 0x00, 0x01, 0x0B, 0x00, 0x00, 0x00, 0x00, // cell #1
			},
			true,
		},
		{
			"small 2 cells delete end",
			[]byte{
				0x00, 0x00, 0x00, 0x01, 0x00, 0x02, // header
				0x00, 0x12, 0x00, 0x0A, // offset #0
				0x00, 0x20, 0x00, 0x0A, // offset #1
				0x00, 0x00, 0x00, 0x00, // free space
				0x01, 0x00, 0x00, 0x00, 0x01, 0x0A, 0x00, 0x00, 0x00, 0x00, // cell #0
				0x00, 0x00, 0x00, 0x00, // free space
				0x01, 0x00, 0x00, 0x00, 0x01, 0x0B, 0x00, 0x00, 0x00, 0x00, // cell #1
			},
			[]byte{0x0B},
			[]byte{
				0x00, 0x00, 0x00, 0x01, 0x00, 0x01, // header
				0x00, 0x12, 0x00, 0x0A, // offset #0
				0x00, 0x00, 0x00, 0x00, // free space
				0x00, 0x00, 0x00, 0x00, // free space
				0x01, 0x00, 0x00, 0x00, 0x01, 0x0A, 0x00, 0x00, 0x00, 0x00, // cell #0
				0x00, 0x00, 0x00, 0x00, // free space
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // free space
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			p, err := load(tt.before)
			assert.NoError(err)

			assert.Equal(tt.before, p.data)
			ok, err := p.DeleteCell(tt.deleteKey)
			assert.Equal(tt.ok, ok)
			assert.NoError(err)
			assert.Equal(tt.after, p.data)
		})
	}
}

func TestPage_findCell(t *testing.T) {
	pageID := ID(0)
	p := New(pageID)
	cells := []CellTyper{
		// these cells should remain sorted, as sorted insertion is tested
		// somewhere else, and by being sorted, the tests are more readable
		// regarding the offset indexes
		PointerCell{
			Key:     []byte("001 first"),
			Pointer: ID(1),
		},
		PointerCell{
			Key:     []byte("002 second"),
			Pointer: ID(2),
		},
		PointerCell{
			Key:     []byte("003 third"),
			Pointer: ID(3),
		},
		PointerCell{
			Key:     []byte("004 fourth"),
			Pointer: ID(4),
		},
	}
	for _, cell := range cells {
		switch c := cell.(type) {
		case RecordCell:
			assert.NoError(t, p.StoreRecordCell(c))
		case PointerCell:
			assert.NoError(t, p.StorePointerCell(c))
		default:
			assert.FailNow(t, "unknown cell type")
		}
	}

	// actual tests

	tests := []struct {
		name            string
		p               *Page
		key             string
		wantOffsetIndex uint16
		wantCellSlot    Slot
		wantCell        CellTyper
		wantFound       bool
	}{
		{
			name:            "first",
			p:               p,
			key:             "001 first",
			wantOffsetIndex: 6,
			wantCellSlot:    Slot{Offset: 16366, Size: 18},
			wantCell:        cells[0],
			wantFound:       true,
		},
		{
			name:            "second",
			p:               p,
			key:             "002 second",
			wantOffsetIndex: 10,
			wantCellSlot:    Slot{Offset: 16347, Size: 19},
			wantCell:        cells[1],
			wantFound:       true,
		},
		{
			name:            "third",
			p:               p,
			key:             "003 third",
			wantOffsetIndex: 14,
			wantCellSlot:    Slot{Offset: 16329, Size: 18},
			wantCell:        cells[2],
			wantFound:       true,
		},
		{
			name:            "fourth",
			p:               p,
			key:             "004 fourth",
			wantOffsetIndex: 18,
			wantCellSlot:    Slot{Offset: 16310, Size: 19},
			wantCell:        cells[3],
			wantFound:       true,
		},
		{
			name:            "missing cell",
			p:               p,
			key:             "some key that doesn't exist",
			wantOffsetIndex: 0,
			wantCellSlot:    Slot{Offset: 0, Size: 0},
			wantCell:        nil,
			wantFound:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			offsetIndex, cellSlot, cell, found := p.findCell([]byte(tt.key))
			assert.Equal(tt.wantOffsetIndex, offsetIndex, "offset indexes don't match")
			assert.Equal(tt.wantCellSlot, cellSlot, "cell slot don't match")
			assert.Equal(tt.wantCell, cell, "cell don't match")
			assert.Equal(tt.wantFound, found, "found don't match")
		})
	}
}
