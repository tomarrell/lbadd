package cache

import (
	"io"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

// Cache describes a page cache that caches pages from a secondary storage.
type Cache interface {
	// FetchAndPin fetches a page from this cache. If it does not exist in the
	// cache, it must be loaded from a configured source. After the page was
	// fetched, it is pinned, meaning that it's guaranteed that the page is not
	// evicted. After working with the page, it must be released again, in order
	// for the cache to be able to free memory. If a page with the given id does
	// not exist, an error will be returned.
	FetchAndPin(id page.ID) (*page.Page, error)
	// Unpin tells the cache that the page with the given id is no longer
	// required directly, and that it can be evicted. Unpin is not a guarantee
	// that the page will be evicted. The cache determines, when to evict a
	// page. If a page with that id does not exist, this call is a no-op.
	Unpin(id page.ID)
	// Flush writes the contents of the page with the given id to the configured
	// source. Before a page is evicted, it is always flushed. Use this method
	// to tell the cache that the page must be flushed immediately. If a page
	// with the given id does not exist, an error will be returned.
	Flush(id page.ID) error
	io.Closer
}
