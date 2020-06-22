package storage

import (
	"encoding/binary"
	"fmt"
	"sync/atomic"
	"unsafe"

	"github.com/rs/zerolog"
	"github.com/spf13/afero"
	"github.com/tomarrell/lbadd/internal/engine/storage/cache"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

const (
	// DefaultCacheSize is the default amount of pages, that the cache can hold
	// at most. Current limit is 256, which, regarding the current page size of
	// 16K, means, that the maximum size that a full cache will occupy, is 4M
	// (CacheSize * page.Size).
	DefaultCacheSize = 1 << 8

	HeaderPageID page.ID = 0

	HeaderTables    = "tables"
	HeaderPageCount = "pageCount"
)

var (
	byteOrder = binary.BigEndian
)

type DBFile struct {
	closed uint32

	log       zerolog.Logger
	cacheSize int

	file        afero.File
	pageManager *PageManager
	cache       cache.Cache

	headerPage *page.Page
}

func Create(file afero.File, opts ...Option) (*DBFile, error) {
	if _, err := file.Stat(); err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	}

	mgr, err := NewPageManager(file)
	if err != nil {
		return nil, fmt.Errorf("new page manager: %w", err)
	}

	headerPage, err := mgr.AllocateNew()
	if err != nil {
		return nil, fmt.Errorf("allocate header page: %w", err)
	}
	tablesPage, err := mgr.AllocateNew()
	if err != nil {
		return nil, fmt.Errorf("allocate tables page: %w", err)
	}

	headerPage.StoreRecordCell(page.RecordCell{
		Key:    []byte(HeaderPageCount),
		Record: encodeUint64(2), // header and tables page
	})
	headerPage.StorePointerCell(page.PointerCell{
		Key:     []byte(HeaderTables),
		Pointer: tablesPage.ID(),
	})

	err = mgr.WritePage(headerPage) // immediately flush
	if err != nil {
		return nil, fmt.Errorf("write header page: %w", err)
	}
	err = mgr.WritePage(tablesPage) // immediately flush
	if err != nil {
		return nil, fmt.Errorf("write tables page: %w", err)
	}

	return newDB(file, mgr, headerPage, opts...), nil
}

func Open(file afero.File, opts ...Option) (*DBFile, error) {
	if err := NewValidator(file).Validate(); err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	mgr, err := NewPageManager(file)
	if err != nil {
		return nil, fmt.Errorf("new page manager: %w", err)
	}

	headerPage, err := mgr.ReadPage(HeaderPageID)
	if err != nil {
		return nil, fmt.Errorf("read header page: %w", err)
	}

	return newDB(file, mgr, headerPage, opts...), nil
}

func (db *DBFile) AllocateNewPage() (*page.Page, error) {
	if db.Closed() {
		return nil, ErrClosed
	}

	page, err := db.pageManager.AllocateNew()
	if err != nil {
		return nil, fmt.Errorf("allocate new: %w", err)
	}
	if err := db.incrementHeaderPageCount(); err != nil {
		return nil, fmt.Errorf("increment header page count: %w", err)
	}
	if err := db.pageManager.WritePage(db.headerPage); err != nil {
		return nil, fmt.Errorf("write header page: %w", err)
	}
	return page, nil
}

func (db *DBFile) Cache() cache.Cache {
	if db.Closed() {
		return nil
	}
	return db.cache
}

func (db *DBFile) Close() error {
	_ = db.cache.Close()
	_ = db.file.Close()
	atomic.StoreUint32(&db.closed, 1)
	return nil
}

func (db *DBFile) Closed() bool {
	return atomic.LoadUint32(&db.closed) == 1
}

func newDB(file afero.File, mgr *PageManager, headerPage *page.Page, opts ...Option) *DBFile {
	db := &DBFile{
		log:       zerolog.Nop(),
		cacheSize: DefaultCacheSize,

		file:        file,
		pageManager: mgr,
		headerPage:  headerPage,
	}
	for _, opt := range opts {
		opt(db)
	}

	db.cache = cache.NewLRUCache(db.cacheSize, mgr)

	return db
}

func (db *DBFile) incrementHeaderPageCount() error {
	val, ok := db.headerPage.Cell([]byte(HeaderPageCount))
	if !ok {
		return fmt.Errorf("no page count header field")
	}
	cell := val.(page.RecordCell)
	byteOrder.PutUint64(cell.Record, byteOrder.Uint64(cell.Record)+1)
	return nil
}

func encodeUint64(v uint64) []byte {
	buf := make([]byte, unsafe.Sizeof(v))
	byteOrder.PutUint64(buf, v)
	return buf
}
