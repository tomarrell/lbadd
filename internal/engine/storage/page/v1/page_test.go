package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

func TestPage_StoreRecordCell(t *testing.T) {
	assert := assert.New(t)

	p, err := load(make([]byte, 36))
	assert.NoError(err)

	c := RecordCell{
		cell: cell{
			key: []byte{0xAB},
		},
		record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
	}

	err = p.StoreRecordCell(c)
	assert.NoError(err)
	assert.Equal([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // header
		0x00, 0x16, 0x00, 0x0E, // offset
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
		Offset: 18,
		Size:   4,
	}, freeSlots[0])

	pageData := make([]byte, len(p.data))
	copy(pageData, p.data)

	anotherCell := RecordCell{
		cell: cell{
			key: []byte("large key"),
		},
		record: []byte("way too large record"),
	}
	err = p.StoreRecordCell(anotherCell)
	assert.Equal(page.ErrPageFull, err)
	assert.Equal(pageData, p.data) // page must not have been modified
}

func TestPage_StoreRecordCell_Multiple(t *testing.T) {
	assert := assert.New(t)

	p, err := load(make([]byte, 64))
	assert.NoError(err)

	cells := []RecordCell{
		{
			cell: cell{
				key: []byte{0x11},
			},
			record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
		},
		{
			cell: cell{
				key: []byte{0x33},
			},
			record: []byte{0xD1, 0xCE},
		},
		{
			cell: cell{
				key: []byte{0x22},
			},
			record: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},
	}
	assert.NoError(p.storeRecordCell(cells[0]))
	assert.NoError(p.storeRecordCell(cells[1]))
	assert.NoError(p.storeRecordCell(cells[2]))
	assert.Equal([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // header
		0x00, 0x3C, 0x00, 0x04, // offset #0
		0x00, 0x3A, 0x00, 0x10, // offset #2
		0x00, 0x3A, 0x00, 0x02, // offset #1
	}, p.data)
}

func TestPage_offsets(t *testing.T) {
	assert := assert.New(t)

	// create a completely empty page
	p, err := load(make([]byte, PageSize))
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
	tests := []struct {
		name    string
		offsets []Slot
		want    Slot
		want1   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
