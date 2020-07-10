package cache

import (
	"fmt"
	"sync"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

var _ Cache = (*LRUCache)(nil)

// LRUCache is a simple implementation of an LRU cache.
type LRUCache struct {
	store     SecondaryStorage
	pages     map[page.ID]*page.Page
	pageLocks map[page.ID]*sync.Mutex
	pinned    map[page.ID]struct{}
	size      int
	lru       []page.ID
}

// NewLRUCache creates a new LRU cache with the given size and secondary storage
// to write dirty pages to. The size is the maximum amount of pages, that can be
// held by this cache. If more pages than the size are requested, old pages will
// be evicted from the cache. If all pages are pinned (in use and not released
// yet), requesting a new page will fail.
func NewLRUCache(size int, store SecondaryStorage) *LRUCache {
	return &LRUCache{
		store:     store,
		pages:     make(map[uint32]*page.Page),
		pageLocks: make(map[uint32]*sync.Mutex),
		pinned:    make(map[uint32]struct{}),
		size:      size,
		lru:       make([]uint32, 0),
	}
}

// FetchAndPin will return the page with the given ID. This will fail, if the
// page with the given ID is not in the cache, but the cache is full and all
// pages are pinned. After obtaining a page with this method, you MUST release
// it, once you are done using it, with Unpin(ID).
func (c *LRUCache) FetchAndPin(id page.ID) (*page.Page, error) {
	return c.fetchAndPin(id)
}

// Unpin unpins the page with the given ID. If the page with the given ID is not
// pinned, then this is a no-op.
func (c *LRUCache) Unpin(id page.ID) {
	c.unpin(id)
}

// Flush writes the contents of the page with the given ID to secondary storage.
// This fails, if the page with the given ID is not in the cache anymore. Only
// call this if you know what you are doing. Pages will always be flushed before
// being evicted. If you really do need to use this, call it before unpinning
// the page, to guarantee that the page will not be evicted between unpinning
// and flushing.
func (c *LRUCache) Flush(id page.ID) error {
	return c.flush(id)
}

// Close will flush all dirty pages and then close this cache.
func (c *LRUCache) Close() error {
	for id := range c.pages {
		_ = c.flush(id)
	}
	return nil
}

func (c *LRUCache) fetchAndPin(id page.ID) (*page.Page, error) {
	// first lock page for others
	lock := c.obtainPageLock(id)
	lock.Lock() // unpin unlocks this lock

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

func (c *LRUCache) obtainPageLock(id page.ID) *sync.Mutex {
	lock, ok := c.pageLocks[id]
	if !ok {
		lock = new(sync.Mutex)
		c.pageLocks[id] = lock
	}
	return lock
}

func (c *LRUCache) pin(id uint32) {
	c.pinned[id] = struct{}{}
}

func (c *LRUCache) unpin(id uint32) {
	delete(c.pinned, id)
	c.pageLocks[id].Unlock() // unlock page lock after page is released by user
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
	page, ok := c.pages[id]
	if !ok || !page.Dirty() {
		return nil
	}
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
