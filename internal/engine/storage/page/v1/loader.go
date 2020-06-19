package v1

import (
	"github.com/spf13/afero"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

var _ page.Loader = (*Loader)(nil)

// Loader is the v1 implementation of a page.Loader, used to retrieve pages from
// secondary storage.
type Loader struct {
	fs afero.Fs
}

// NewLoader creates a new, ready to use loader. If during initialization, an
// error occurs, the error will be returned. It may be wrapped.
func NewLoader(fs afero.Fs) (*Loader, error) {
	l := &Loader{
		fs: fs,
	}
	// TODO: add initialization if needed
	return l, nil
}

// Load loads the page with the given ID from the database files.
func (l *Loader) Load(id page.ID) (page.Page, error) {
	return l.load(id)
}

func (l *Loader) load(id page.ID) (*Page, error) {
	return nil, nil
}
