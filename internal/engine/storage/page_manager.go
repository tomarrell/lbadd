package storage

import (
	"fmt"
	"sync/atomic"

	"github.com/spf13/afero"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

// PageManager is a manager that is responsible for reading pages from and
// writing pages to secondary storage. It also can allocate new pages, which
// will immediately be written to secondary storage.
type PageManager struct {
	file      afero.File
	largestID page.ID
}

// NewPageManager creates a new page manager over the given file. It is assumed,
// that the file passed validation by a (*storage.Validator).
func NewPageManager(file afero.File) (*PageManager, error) {
	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	mgr := &PageManager{
		file:      file,
		largestID: uint32(stat.Size() / page.Size),
	}

	return mgr, nil
}

// ReadPage returns the page with the given ID, or an error if reading is
// impossible.
func (m *PageManager) ReadPage(id page.ID) (*page.Page, error) {
	data := make([]byte, page.Size)
	_, err := m.file.ReadAt(data, int64(id)*page.Size)
	if err != nil {
		return nil, fmt.Errorf("read at: %w", err)
	}
	p, err := page.Load(data)
	if err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}
	return p, nil
}

// WritePage will write the given page into secondary storage. It is guaranteed,
// that after this call returns, the page is present on disk.
func (m *PageManager) WritePage(p *page.Page) error {
	data := p.RawData()
	_, err := m.file.WriteAt(data, int64(p.ID())*page.Size)
	if err != nil {
		return fmt.Errorf("write at: %w", err)
	}
	if err := m.file.Sync(); err != nil {
		return fmt.Errorf("sync: %w", err)
	}
	return nil
}

// AllocateNew will allocate a new page and immediately persist it in secondary
// storage. It is guaranteed, that after this call returns, the page is present
// on disk.
func (m *PageManager) AllocateNew() (*page.Page, error) {
	id := atomic.AddUint32(&m.largestID, 1) - 1

	p := page.New(id)
	if err := m.WritePage(p); err != nil {
		return nil, fmt.Errorf("write new page: %w", err)
	}
	return p, nil
}

// Close will sync the file with secondary storage and then close it. If syncing
// fails, the file will not be closed, and an error will be returned.
func (m *PageManager) Close() error {
	if err := m.file.Sync(); err != nil {
		return fmt.Errorf("sync: %w", err)
	}
	_ = m.file.Close()
	return nil
}
