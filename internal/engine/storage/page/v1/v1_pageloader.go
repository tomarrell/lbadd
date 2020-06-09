package v1

//go:generate mockery -inpkg -testonly -case=snake -name=SecondaryStorage

// SecondaryStorage descries a secondary storage component, such as a disk. It
// has to manage pages by ID and must ensure that pages are read and written
// from the underlying storage without any caching.
type SecondaryStorage interface {
	LoadPage(uint32) (*Page, error)
	WritePage(uint32, *Page) error
}
