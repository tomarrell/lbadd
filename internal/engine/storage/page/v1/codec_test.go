package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

func Test_encodeCell(t *testing.T) {
	tests := []struct {
		name string
		cell page.Cell
		want []byte
	}{
		{
			"empty cell",
			page.Cell{
				Key:    []byte{},
				Record: []byte{},
			},
			[]byte{0x00, 0x00, 0x00, 0x00},
		},
		{
			"cell",
			page.Cell{
				Key:    []byte{0x01, 0x02},
				Record: []byte{0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			},
			[]byte{
				0x00, 0x02, // key frame
				0x01, 0x02, // key
				0x00, 0x06, // record frame
				0x03, 0x04, 0x05, 0x06, 0x07, 0x08, // record
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.want, encodeCell(tt.cell))
		})
	}
}

func Test_decodeCell(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want page.Cell
	}{
		{
			"empty cell",
			[]byte{0x00, 0x00, 0x00, 0x00},
			page.Cell{
				Key:    []byte{},
				Record: []byte{},
			},
		},
		{
			"cell",
			[]byte{
				0x00, 0x02, // key frame
				0x01, 0x02, // key
				0x00, 0x06, // record frame
				0x03, 0x04, 0x05, 0x06, 0x07, 0x08, // record
			},
			page.Cell{
				Key:    []byte{0x01, 0x02},
				Record: []byte{0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tt.want, decodeCell(tt.data))
		})
	}
}

func Test_encodeOffset(t *testing.T) {
	tests := []struct {
		name   string
		offset Offset
		want   []byte
	}{
		{
			"empty offset",
			Offset{},
			[]byte{0x00, 0x00, 0x00, 0x00},
		},
		{
			"offset",
			Offset{
				Location: 0xCAFE,
				Size:     0xBABE,
			},
			[]byte{0xCA, 0xFE, 0xBA, 0xBE},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.want, encodeOffset(tt.offset))
		})
	}
}

func Test_decodeOffset(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want Offset
	}{
		{
			"empty offset",
			[]byte{0x00, 0x00, 0x00, 0x00},
			Offset{},
		},
		{
			"offset",
			[]byte{0xCA, 0xFE, 0xBA, 0xBE},
			Offset{
				Location: 0xCAFE,
				Size:     0xBABE,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.want, decodeOffset(tt.data))
		})
	}
}
