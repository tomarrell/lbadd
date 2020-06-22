package storage

import (
	"fmt"
	"sync/atomic"

	"github.com/spf13/afero"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

type PageManager struct {
	file      afero.File
	largestID page.ID
}

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

func (m *PageManager) WritePage(p *page.Page) error {
	data := p.RawData() // TODO: avoid copying in RawData()
	_, err := m.file.WriteAt(data, int64(p.ID())*page.Size)
	if err != nil {
		return fmt.Errorf("write at: %w", err)
	}
	if err := m.file.Sync(); err != nil {
		return fmt.Errorf("sync: %w", err)
	}
	return nil
}

func (m *PageManager) AllocateNew() (*page.Page, error) {
	id := atomic.AddUint32(&m.largestID, 1) - 1

	p := page.New(id)
	if err := m.WritePage(p); err != nil {
		return nil, fmt.Errorf("write new page: %w", err)
	}
	return p, nil
}
