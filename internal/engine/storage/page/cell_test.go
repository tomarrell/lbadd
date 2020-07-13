package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_frame(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want []byte
	}{
		{
			"empty",
			[]byte{},
			[]byte{
				0x00, 0x00, 0x00, 0x00, // frame
				// no data
			},
		},
		{
			"single",
			[]byte{0xD1},
			[]byte{
				0x00, 0x00, 0x00, 0x01, // frame
				0xD1, // data
			},
		},
		{
			"double",
			[]byte{0xD1, 0xCE},
			[]byte{
				0x00, 0x00, 0x00, 0x02, // frame
				0xD1, 0xCE, // data
			},
		},
		{
			"large",
			make([]byte, 1<<16), // 64KB
			append(
				[]byte{0x00, 0x01, 0x00, 0x00}, // frame
				make([]byte, 1<<16)...,         // data
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got := frame(tt.data)
			assert.Equal(tt.want, got)
		})
	}
}

func Test_encodeRecordCell(t *testing.T) {
	tests := []struct {
		name string
		cell RecordCell
		want []byte
	}{
		{
			"empty",
			RecordCell{},
			[]byte{
				byte(CellTypeRecord),   // cell type
				0x00, 0x00, 0x00, 0x00, // key frame
				// no key
				0x00, 0x00, 0x00, 0x00, // record frame
				// no record
			},
		},
		{
			"small",
			RecordCell{
				Key:    []byte{0xD1, 0xCE},
				Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
			},
			[]byte{
				byte(CellTypeRecord),   // cell type
				0x00, 0x00, 0x00, 0x02, // key frame
				0xD1, 0xCE, // key
				0x00, 0x00, 0x00, 0x04, // record frame
				0xCA, 0xFE, 0xBA, 0xBE, // record
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got := encodeRecordCell(tt.cell)
			assert.Equal(tt.want, got)
		})
	}
}

func Test_encodePointerCell(t *testing.T) {
	tests := []struct {
		name string
		cell PointerCell
		want []byte
	}{
		{
			"empty",
			PointerCell{},
			[]byte{
				byte(CellTypePointer),  // cell type
				0x00, 0x00, 0x00, 0x00, // key frame
				// no key
				0x00, 0x00, 0x00, 0x00, // pointer
			},
		},
		{
			"simple",
			PointerCell{
				Key:     []byte{0xD1, 0xCE},
				Pointer: 0xCAFEBABE,
			},
			[]byte{
				byte(CellTypePointer),  // cell type
				0x00, 0x00, 0x00, 0x02, // key frame
				0xD1, 0xCE, // key
				0xCA, 0xFE, 0xBA, 0xBE, // pointer
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got := encodePointerCell(tt.cell)
			assert.Equal(tt.want, got)
		})
	}
}

func TestAnyCell_Type(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(CellTypeRecord, RecordCell{}.Type())
	assert.Equal(CellTypePointer, PointerCell{}.Type())
}

func Test_decodeRecordCell(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want RecordCell
	}{
		{
			"zero value",
			[]byte{
				byte(CellTypeRecord),   // cell type
				0x00, 0x00, 0x00, 0x00, // key frame
				// no key
				0x00, 0x00, 0x00, 0x00, // record frame
				// no record
			},
			RecordCell{
				Key:    []byte{},
				Record: []byte{},
			},
		},
		{
			"simple",
			[]byte{
				byte(CellTypeRecord),   // cell type
				0x00, 0x00, 0x00, 0x02, // key frame
				0xD1, 0xCE, // key
				0x00, 0x00, 0x00, 0x04, // record frame
				0xCA, 0xFE, 0xBA, 0xBE, // record
			},
			RecordCell{
				Key:    []byte{0xD1, 0xCE},
				Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got := decodeRecordCell(tt.data)
			assert.Equal(tt.want, got)
		})
	}
}

func Test_decodePointerCell(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want PointerCell
	}{
		{
			"zero value",
			[]byte{
				byte(CellTypePointer),  // cell type
				0x00, 0x00, 0x00, 0x00, // key frame
				// no key
				0x00, 0x00, 0x00, 0x00, // pointer
			},
			PointerCell{
				Key:     []byte{},
				Pointer: 0,
			},
		},
		{
			"simple",
			[]byte{
				byte(CellTypePointer),  // cell type
				0x00, 0x00, 0x00, 0x02, // key frame
				0xD1, 0xCE, // key
				0xCA, 0xFE, 0xBA, 0xBE, // pointer
			},
			PointerCell{
				Key:     []byte{0xD1, 0xCE},
				Pointer: 0xCAFEBABE,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got := decodePointerCell(tt.data)
			assert.Equal(tt.want, got)
		})
	}
}

func Test_decodeCell(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want CellTyper
	}{
		{
			"zero record cell",
			[]byte{
				byte(CellTypeRecord),   // cell type
				0x00, 0x00, 0x00, 0x00, // key frame
				// no key
				0x00, 0x00, 0x00, 0x00, // record frame
				// no record
			},
			RecordCell{
				Key:    []byte{},
				Record: []byte{},
			},
		},
		{
			"zero pointer cell",
			[]byte{
				byte(CellTypePointer),  // cell type
				0x00, 0x00, 0x00, 0x00, // key frame
				// no key
				0x00, 0x00, 0x00, 0x00, // pointer
			},
			PointerCell{
				Key:     []byte{},
				Pointer: 0,
			},
		},
		{
			"invalid cell type",
			[]byte{
				byte(CellType(123)),    // cell type
				0x00, 0x00, 0x00, 0x00, // key frame
				// no key
				0x00, 0x00, 0x00, 0x00, // pointer
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			got := decodeCell(tt.data)
			assert.Equal(tt.want, got)
		})
	}
}
