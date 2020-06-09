package v1

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

const (
	// CacheSize is the default amount of pages that a cache can hold.
	CacheSize int = 1 << 8
)

// Cache is a simple LRU implementation of a page cache. Lookups in the cache
// are performed with amortized O(1) complexity, however, moving elements in the
// LRU list is more expensive.
type Cache struct {
	secondaryStorage SecondaryStorage

	pages  map[uint32]*Page
	pinned map[uint32]struct{}

	size int
	lru  []uint32
}

// NewCache creates a new cache with the given secondary storage.
func NewCache(secondaryStorage SecondaryStorage) page.Cache {
	return newCache(secondaryStorage)
}

// FetchAndPin fetches the page with the given ID from the cache or the
// secondary storage and pins it in the cache. Pinning prevents the page from
// being evicted.
func (c *Cache) FetchAndPin(id uint32) (page.Page, error) {
	return c.fetchAndPin(id)
}

// Unpin marks the given ID as ok to be evicted. This is a no-op if the ID
// doesn't exist or is not pinned.
func (c *Cache) Unpin(id uint32) {
	c.unpin(id)
}

// Flush writes the contents of the page with the given ID to the secondary
// storage.
func (c *Cache) Flush(id uint32) error {
	return c.flush(id)
}

// Close does nothing.
func (c *Cache) Close() error {
	return nil
}

func newCache(secondaryStorage SecondaryStorage) *Cache {
	return &Cache{
		secondaryStorage: secondaryStorage,

		pages:  make(map[uint32]*Page),
		pinned: make(map[uint32]struct{}),

		size: CacheSize,
		lru:  make([]uint32, 0, CacheSize),
	}
}

func (c *Cache) fetchAndPin(id uint32) (*Page, error) {
	// pin id first in order to avoid potential concurrent eviction at this
	// point
	c.pin(id)
	p, err := c.fetch(id)
	if err != nil {
		// unpin if a page with the given id cannot be loaded
		c.unpin(id)
		return nil, err
	}
	return p, nil
}

func (c *Cache) fetch(id uint32) (*Page, error) {
	// check if page is already cached
	if p, ok := c.pages[id]; ok {
		moveToFront(id, c.lru)
		return p, nil
	}

	// check if we have to evict pages
	if err := c.freeMemoryIfNeeded(); err != nil {
		return nil, fmt.Errorf("free mem: %w", err)
	}

	// fetch page from secondary storage
	p, err := c.secondaryStorage.LoadPage(id)
	if err != nil {
		return nil, fmt.Errorf("load page: %w", err)
	}
	// store in cache
	c.pages[id] = p
	// append in front
	c.lru = append([]uint32{id}, c.lru...)
	return p, nil
}

func (c *Cache) pin(id uint32) {
	c.pinned[id] = struct{}{}
}

func (c *Cache) unpin(id uint32) {
	delete(c.pinned, id)
}

func (c *Cache) evict(id uint32) error {
	if err := c.flush(id); err != nil {
		return fmt.Errorf("flush: %w", err)
	}
	delete(c.pages, id)
	return nil
}

func (c *Cache) flush(id uint32) error {
	if err := c.secondaryStorage.WritePage(id, c.pages[id]); err != nil {
		return fmt.Errorf("write page: %w", err)
	}
	return nil
}

func (c *Cache) freeMemoryIfNeeded() error {
	if len(c.lru) < c.size {
		return nil
	}
	for i := len(c.lru) - 1; i >= 0; i-- {
		id := c.lru[i]
		if _, ok := c.pinned[id]; ok {
			// can't evict current page, pinned
			continue
		}
		c.lru = c.lru[:len(c.lru)-1]
		return c.evict(id)
	}
	return fmt.Errorf("all pages pinned, cache is full")
}

func moveToFront(needle uint32, haystack []uint32) {
	if len(haystack) == 0 || haystack[0] == needle {
		return
	}
	var prev uint32
	for i, elem := range haystack {
		switch {
		case i == 0:
			haystack[0] = needle
			prev = elem
		case elem == needle:
			haystack[i] = prev
			return
		default:
			haystack[i] = prev
			prev = elem
		}
	}
}
