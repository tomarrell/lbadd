package cache

import "github.com/tomarrell/lbadd/internal/engine/storage/page"

//go:generate mockery -inpkg -case=snake -testonly -name SecondaryStorage

// SecondaryStorage is the abstraction that a cache uses, to synchronize dirty
// pages with secondary storage.
type SecondaryStorage interface {
	ReadPage(page.ID) (*page.Page, error)
	WritePage(*page.Page) error
}
