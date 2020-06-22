package cache

import (
	"fmt"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

type SecondaryStorage interface {
	ReadPage(page.ID) (*page.Page, error)
	WritePage(*page.Page) error
}

var _ Cache = (*LRUCache)(nil)

type LRUCache struct {
	store  SecondaryStorage
	pages  map[uint32]*page.Page
	pinned map[uint32]struct{}
	size   int
	lru    []uint32
}

func NewLRUCache(size int, store SecondaryStorage) *LRUCache {
	return &LRUCache{
		store:  store,
		pages:  make(map[uint32]*page.Page),
		pinned: make(map[uint32]struct{}),
		size:   size,
		lru:    make([]uint32, size),
	}
}

func (c *LRUCache) FetchAndPin(id page.ID) (*page.Page, error) {
	return c.fetchAndPin(id)
}

func (c *LRUCache) Unpin(id page.ID) {
	c.unpin(id)
}

func (c *LRUCache) Flush(id page.ID) error {
	return c.flush(id)
}

func (c *LRUCache) Close() error {
	return nil
}

func (c *LRUCache) fetchAndPin(id page.ID) (*page.Page, error) {
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

func (c *LRUCache) pin(id uint32) {
	c.pinned[id] = struct{}{}
}

func (c *LRUCache) unpin(id uint32) {
	delete(c.pinned, id)
}

func (c *LRUCache) fetch(id uint32) (*page.Page, error) {
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
	p, err := c.store.ReadPage(id)
	if err != nil {
		return nil, fmt.Errorf("load page: %w", err)
	}
	// store in cache
	c.pages[id] = p
	// append in front
	c.lru = append([]uint32{id}, c.lru...)
	return p, nil
}

func (c *LRUCache) evict(id uint32) error {
	if err := c.flush(id); err != nil {
		return fmt.Errorf("flush: %w", err)
	}
	delete(c.pages, id)
	return nil
}

func (c *LRUCache) flush(id page.ID) error {
	if err := c.store.WritePage(c.pages[id]); err != nil {
		return fmt.Errorf("write page: %w", err)
	}
	return nil
}

func (c *LRUCache) freeMemoryIfNeeded() error {
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

func moveToFront(needle page.ID, haystack []page.ID) {
	if len(haystack) == 0 || haystack[0] == needle {
		return
	}
	var prev page.ID
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
