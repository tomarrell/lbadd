package page_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
	v1 "github.com/tomarrell/lbadd/internal/engine/storage/page/v1"
)

var (
	Loaders = []struct {
		Name     string
		Loader   page.Loader
		PageSize int
	}{
		{"v1", v1.Load, 1 << 14},
	}

	Cells = []page.Cell{
		{
			Key:    []byte("key[0]"),
			Record: []byte("data[0]"),
		},
		{
			Key:    []byte("key[1]"),
			Record: []byte("data[1]"),
		},
		{
			Key:    []byte("key[2]"),
			Record: []byte("data[2]"),
		},
		{
			Key:    []byte("key[3]"),
			Record: []byte("data[3]"),
		},
	}
)

func TestImplementations(t *testing.T) {
	for _, loader := range Loaders {
		t.Run("loader="+loader.Name, func(t *testing.T) { _TestPageOperations(t, loader.Loader, loader.PageSize) })
	}
}

func _TestPageOperations(t *testing.T, loader page.Loader, pageSize int) {
	assert := assert.New(t)
	rand := rand.New(rand.NewSource(87234562678)) // reproducible random
	cells := make([]page.Cell, len(Cells))
	copy(cells, Cells)
	rand.Shuffle(len(cells), func(i, j int) { cells[i], cells[j] = cells[j], cells[i] })

	data := make([]byte, pageSize)
	p, err := loader(data)
	assert.NoError(err)
	assert.NotNil(p)

	for _, cell := range cells {
		err = p.StoreCell(cell)
		assert.NoError(err)
	}

	// after insertion, all cells must be returned sorted, as the original Cells
	// slice
	assert.Equal(Cells, p.Cells())
	for _, cell := range cells {
		obtained, ok := p.Cell(cell.Key)
		assert.True(ok)
		assert.Equal(cell, obtained)
	}

	// delete one cell
	deleteIndex := 2
	err = p.Delete(Cells[deleteIndex].Key)
	assert.NoError(err)

	afterDeletionCells := make([]page.Cell, len(Cells))
	copy(afterDeletionCells, Cells)
	afterDeletionCells = append(afterDeletionCells[:deleteIndex], afterDeletionCells[deleteIndex+1:]...)
	assert.Equal(afterDeletionCells, p.Cells())
	obtained, ok := p.Cell(Cells[deleteIndex].Key)
	assert.False(ok, "deleted cell must not be obtainable anymore")
	assert.Zero(obtained)
}
