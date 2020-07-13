package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

func TestLRUCache(t *testing.T) {
	assert := assert.New(t)

	pages := make([]*page.Page, 6)
	for i := range pages {
		pages[i] = page.New(uint32(i))
	}

	secondaryStorage := new(MockSecondaryStorage)
	secondaryStorage.
		On("ReadPage", mock.IsType(page.ID(0))).
		Return(func(id uint32) *page.Page {
			return pages[id]
		}, nil)
	secondaryStorage.
		On("WritePage", mock.IsType((*page.Page)(nil))).
		Return(nil)

	c := NewLRUCache(5, secondaryStorage)
	defer func() { _ = c.Close() }()

	// load 5 pages, fill cache up
	p, err := c.FetchAndPin(0)
	assert.NoError(err)
	assert.Same(pages[0], p)
	secondaryStorage.AssertCalled(t, "ReadPage", uint32(0))
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything)

	p, err = c.FetchAndPin(1)
	assert.NoError(err)
	assert.Same(pages[1], p)
	secondaryStorage.AssertCalled(t, "ReadPage", uint32(1))
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything)

	p, err = c.FetchAndPin(2)
	assert.NoError(err)
	assert.Same(pages[2], p)
	secondaryStorage.AssertCalled(t, "ReadPage", uint32(2))
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything)

	p, err = c.FetchAndPin(3)
	assert.NoError(err)
	assert.Same(pages[3], p)
	secondaryStorage.AssertCalled(t, "ReadPage", uint32(3))
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything)

	p, err = c.FetchAndPin(4)
	assert.NoError(err)
	assert.Same(pages[4], p)
	secondaryStorage.AssertCalled(t, "ReadPage", uint32(4))
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything)

	// all pages are fetched and locked now, cache can not evict any pages
	p, err = c.FetchAndPin(5)
	assert.Error(err)
	assert.Nil(p)
	secondaryStorage.AssertNotCalled(t, "ReadPage", uint32(5))
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything)

	// must release a page
	c.Unpin(0) // unpin first page
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything)

	// load another page
	p, err = c.FetchAndPin(5) // page[5] can now be loaded
	assert.NoError(err)
	assert.Same(pages[5], p)
	secondaryStorage.AssertNotCalled(t, "WritePage", mock.Anything) // page 0 evicted, but no writes, because page 0 was not dirty
	secondaryStorage.AssertCalled(t, "ReadPage", uint32(5))         // page 5 loaded

	// mark page 1 as dirty
	pages[1].MarkDirty()
	// release page 1
	c.Unpin(1)

	// load another page again
	p, err = c.FetchAndPin(0) // page[0] can now be loaded
	assert.NoError(err)
	assert.Same(pages[0], p)
	secondaryStorage.AssertCalled(t, "WritePage", pages[1]) // page 1 evicted
	secondaryStorage.AssertCalled(t, "ReadPage", uint32(0)) // page 0 loaded

	secondaryStorage.AssertNumberOfCalls(t, "ReadPage", 7)
}
