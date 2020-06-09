package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCacheQueue(t *testing.T) {
	assert := assert.New(t)
	must := func(p *Page, err error) *Page { assert.NoError(err); return p }
	pages := map[uint32]*Page{
		0: must(load(make([]byte, PageSize))),
		1: must(load(make([]byte, PageSize))),
		2: must(load(make([]byte, PageSize))),
		3: must(load(make([]byte, PageSize))),
		4: must(load(make([]byte, PageSize))),
	}

	secondaryStorage := new(MockSecondaryStorage)
	defer secondaryStorage.AssertExpectations(t)

	c := newCache(secondaryStorage)
	c.size = 2

	// first page - unpin after use
	secondaryStorage.
		On("LoadPage", mock.IsType(uint32(0))).
		Return(pages[0], nil).
		Once()
	p, err := c.FetchAndPin(0)
	secondaryStorage.AssertCalled(t, "LoadPage", uint32(0))
	assert.NoError(err)
	assert.Same(pages[0], p)
	assert.Equal([]uint32{0}, c.lru)
	c.Unpin(0)

	// second page - keep pinned
	secondaryStorage.
		On("LoadPage", mock.IsType(uint32(0))).
		Return(pages[1], nil).
		Once()
	p, err = c.FetchAndPin(1)
	secondaryStorage.AssertCalled(t, "LoadPage", uint32(1))
	assert.NoError(err)
	assert.Same(pages[1], p)
	assert.Equal([]uint32{1, 0}, c.lru)

	// third page - pages[0] must be evicted
	secondaryStorage.
		On("LoadPage", mock.IsType(uint32(0))).
		Return(pages[2], nil).
		Once()
	secondaryStorage.
		On("WritePage", mock.IsType(uint32(0)), mock.IsType(&Page{})).
		Return(nil).
		Once()
	p, err = c.FetchAndPin(2)
	secondaryStorage.AssertCalled(t, "LoadPage", uint32(2))
	secondaryStorage.AssertCalled(t, "WritePage", uint32(0), pages[0])
	assert.NoError(err)
	assert.Same(pages[2], p)
	assert.Equal([]uint32{2, 1}, c.lru)

	// fourth page - can't fetch because cache is full
	p, err = c.FetchAndPin(3)
	assert.Error(err)
	assert.Nil(p)
	assert.Equal([]uint32{2, 1}, c.lru)
}
