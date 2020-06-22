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

	// HeaderPageID is the page ID of the header page of the database file.
	HeaderPageID page.ID = 0

	// HeaderTables is the string key for the header page's cell "tables"
	HeaderTables = "tables"
	// HeaderPageCount is the string key for the header page's cell "pageCount"
	HeaderPageCount = "pageCount"
)

var (
	byteOrder = binary.BigEndian
)

// DBFile is a database file that can be opened or created. From this file, you
// can obtain a page cache, which you must use for reading pages.
type DBFile struct {
	closed uint32

	log       zerolog.Logger
	cacheSize int

	file        afero.File
	pageManager *PageManager
	cache       cache.Cache

	headerPage *page.Page
}

// Create creates a new database in the given file with the given options. The
// file must exist, but be empty and must be a regular file.
func Create(file afero.File, opts ...Option) (*DBFile, error) {
	if info, err := file.Stat(); err != nil {
		return nil, fmt.Errorf("stat: %w", err)
	} else if info.IsDir() {
		return nil, fmt.Errorf("file is directory")
	} else if size := info.Size(); size != 0 {
		return nil, fmt.Errorf("file is not empty, has %v bytes", size)
	} else if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("file is not a regular file")
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

	if err := headerPage.StoreRecordCell(page.RecordCell{
		Key:    []byte(HeaderPageCount),
		Record: encodeUint64(2), // header and tables page
	}); err != nil {
		return nil, fmt.Errorf("store record cell: %w", err)
	}
	if err := headerPage.StorePointerCell(page.PointerCell{
		Key:     []byte(HeaderTables),
		Pointer: tablesPage.ID(),
	}); err != nil {
		return nil, fmt.Errorf("store pointer cell: %w", err)
	}

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

// Open opens and validates a given file and creates a (*DBFile) with the given
// options. If the validation fails, an error explaining the failure will be
// returned.
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

// AllocateNewPage allocates and immediately persists a new page in secondary
// storage. This will fail if the DBFile is closed. After this method returns,
// the allocated page can immediately be found by the cache, and you can use the
// returned page ID to load the page through the cache.
func (db *DBFile) AllocateNewPage() (page.ID, error) {
	if db.Closed() {
		return 0, ErrClosed
	}

	page, err := db.pageManager.AllocateNew()
	if err != nil {
		return 0, fmt.Errorf("allocate new: %w", err)
	}
	if err := db.incrementHeaderPageCount(); err != nil {
		return 0, fmt.Errorf("increment header page count: %w", err)
	}
	if err := db.pageManager.WritePage(db.headerPage); err != nil {
		return 0, fmt.Errorf("write header page: %w", err)
	}
	return page.ID(), nil
}

// Cache returns the cache implementation, that you must use to obtain pages.
// This will fail if the DBFile is closed.
func (db *DBFile) Cache() cache.Cache {
	if db.Closed() {
		return nil
	}
	return db.cache
}

// Close will close the underlying cache, as well as page manager, as well as
// file.
func (db *DBFile) Close() error {
	_ = db.cache.Close()
	_ = db.pageManager.Close()
	_ = db.file.Close()
	atomic.StoreUint32(&db.closed, 1)
	return nil
}

// Closed indicates, whether this file was closed.
func (db *DBFile) Closed() bool {
	return atomic.LoadUint32(&db.closed) == 1
}

// newDB creates a new DBFile from the given objects, and applies all options.
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

// incrementHeaderPageCount will increment the 8 byte uint64 in the
// HeaderPageCount cell by 1.
func (db *DBFile) incrementHeaderPageCount() error {
	val, ok := db.headerPage.Cell([]byte(HeaderPageCount))
	if !ok {
		return fmt.Errorf("no page count header field")
	}
	cell := val.(page.RecordCell)
	byteOrder.PutUint64(cell.Record, byteOrder.Uint64(cell.Record)+1)
	return nil
}

// encodeUint64 will allocate 8 bytes to encode the given uint64 into. This
// newly allocated byte-slice is then returned.
func encodeUint64(v uint64) []byte {
	buf := make([]byte, unsafe.Sizeof(v)) // #nosec
	byteOrder.PutUint64(buf, v)
	return buf
}
