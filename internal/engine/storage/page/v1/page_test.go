package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.EqualValues(OffsetSize, len(offsetData)/offsetCount)

	// inject the offset data
	p.incrementCellCount(3)               // set the cell count
	copy(p.data[HeaderSize:], offsetData) // copy the offset data

	// actual test can start

	offsets := p.Offsets()
	assert.Len(offsets, 3)
	assert.Equal(Offset{
		Offset: 0x0112,
		Size:   0x2334,
	}, offsets[0])
	assert.Equal(Offset{
		Offset: 0x4556,
		Size:   0x6778,
	}, offsets[1])
	assert.Equal(Offset{
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
