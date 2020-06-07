package v1

import (
	"bytes"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine/converter"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

func Test_Page_StoreCell(t *testing.T) {
	t.Run("single cell", func(t *testing.T) {
		assert := assert.New(t)

		p, err := load(make([]byte, PageSize)) // empty page: version=0,id=0,cellcount=0
		assert.NoError(err)
		assert.NotNil(p)

		assert.NoError(
			p.StoreCell(page.Cell{
				Key:    []byte{0xCA},
				Record: []byte{0xFE, 0xBA, 0xBE},
			}),
		)

		assert.Equal([]byte{
			0x00, 0x01, // key frame
			0xCA,       // key
			0x00, 0x03, // record frame
			0xFE, 0xBA, 0xBE, // record
		}, p.data[PageSize-8:])

		assert.EqualValues(1, p.header[InternalHeaderCellCount])
		assert.Equal([]byte{
			0x3D, 0xF8, // location of our cell
			0x00, 0x08, // size of our cell
		}, p.data[ContentOffset:ContentOffset+4])

		allCells := p.Cells()
		assert.Len(allCells, 1)
	})
	t.Run("multiple cells", func(t *testing.T) {
		assert := assert.New(t)

		p, err := load(make([]byte, PageSize)) // empty page: version=0,id=0,cellcount=0
		assert.NoError(err)
		assert.NotNil(p)

		// first cell

		assert.NoError(
			p.StoreCell(page.Cell{
				Key:    []byte{0xAA},
				Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
			}),
		)

		assert.Equal([]byte{
			0x00, 0x01, // key frame
			0xAA,       // key
			0x00, 0x04, // record frame
			0xCA, 0xFE, 0xBA, 0xBE, // record
		}, p.data[PageSize-9:])

		// second cell

		assert.NoError(
			p.StoreCell(page.Cell{
				Key:    []byte{0xFF},
				Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
			}),
		)

		assert.Equal([]byte{
			0x00, 0x01, // key frame
			0xFF,       // key
			0x00, 0x04, // record frame
			0xCA, 0xFE, 0xBA, 0xBE, // record
		}, p.data[PageSize-18:PageSize-9])

		// third cell

		assert.NoError(
			p.StoreCell(page.Cell{
				Key:    []byte{0x11},
				Record: []byte{0xCA, 0xFE, 0xBA, 0xBE},
			}),
		)

		assert.Equal([]byte{
			0x00, 0x01, // key frame
			0x11,       // key
			0x00, 0x04, // record frame
			0xCA, 0xFE, 0xBA, 0xBE, // record
		}, p.data[PageSize-27:PageSize-18])

		// check that all the offsets at the beginning of the content area are
		// sorted by the key of the cell they point to

		x11Offset := converter.Uint16ToByteArray(ContentSize - 27)
		xAAOffset := converter.Uint16ToByteArray(ContentSize - 9)
		xFFOffset := converter.Uint16ToByteArray(ContentSize - 18)

		assert.Equal([]byte{
			// first offset
			x11Offset[0], x11Offset[1], // location
			0x00, 0x09,
			// second offset
			xAAOffset[0], xAAOffset[1], // location
			0x00, 0x09,
			// third offset
			xFFOffset[0], xFFOffset[1], // location
			0x00, 0x09,
		}, p.data[ContentOffset:ContentOffset+OffsetSize*3])

		allCells := p.Cells()
		assert.Len(allCells, 3)
		assert.True(sort.SliceIsSorted(allCells, func(i, j int) bool {
			return bytes.Compare(allCells[i].Key, allCells[j].Key) < 0
		}), "p.Cells() must return all cells ordered by key")
	})
}

func TestInvalidPageSize(t *testing.T) {
	tf := func(t *testing.T, data []byte) {
		assert := assert.New(t)
		p, err := Load(data)
		assert.Equal(page.ErrInvalidPageSize, err)
		assert.Nil(p)
	}
	t.Run("invalid=nil", func(t *testing.T) {
		tf(t, nil)
	})
	t.Run("invalid=smaller", func(t *testing.T) {
		data := make([]byte, PageSize/2)
		tf(t, data)
	})
	t.Run("invalid=larger", func(t *testing.T) {
		data := make([]byte, int(PageSize)*2)
		tf(t, data)
	})
}

func TestHeaderVersion(t *testing.T) {
	assert := assert.New(t)

	versionBytes := []byte{0xCA, 0xFE}

	data := make([]byte, PageSize)
	copy(data[:2], versionBytes)

	p, err := load(data)
	assert.NoError(err)

	version := p.Header(page.HeaderVersion)
	assert.IsType(uint16(0), version)
	assert.EqualValues(0xCAFE, version)

	assert.Equal(versionBytes, p.data[:2])
	for _, b := range p.data[2:] {
		assert.EqualValues(0, b)
	}
}

func TestHeaderID(t *testing.T) {
	assert := assert.New(t)

	idBytes := []byte{0xCA, 0xFE, 0xBA, 0xBE}

	data := make([]byte, PageSize)
	copy(data[2:6], idBytes)

	p, err := load(data)
	assert.NoError(err)

	id := p.Header(page.HeaderID)
	assert.IsType(uint16(0), id)
	assert.EqualValues(0xCAFE, id)

	for _, b := range p.data[:2] {
		assert.EqualValues(0, b)
	}
	assert.Equal(idBytes, p.data[2:6])
	for _, b := range p.data[6:] {
		assert.EqualValues(0, b)
	}
}

func TestHeaderCellCount(t *testing.T) {
	assert := assert.New(t)

	cellCountBytes := []byte{0xCA, 0xFE}

	data := make([]byte, PageSize)
	copy(data[6:8], cellCountBytes)

	p, err := load(data)
	assert.NoError(err)

	cellCount := p.Header(InternalHeaderCellCount)
	assert.IsType(uint16(0), cellCount)
	assert.EqualValues(0xCAFE, cellCount)

	for _, b := range p.data[:6] {
		assert.EqualValues(0, b)
	}
	assert.Equal(cellCountBytes, p.data[6:8])
	for _, b := range p.data[8:] {
		assert.EqualValues(0, b)
	}
}
