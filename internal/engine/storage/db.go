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
	// HeaderConfig is the string key for the header page's cell "config"
	HeaderConfig = "config"
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
	configPage *page.Page
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
	configPage, err := mgr.AllocateNew()
	if err != nil {
		return nil, fmt.Errorf("allocate config page: %w", err)
	}
	tablesPage, err := mgr.AllocateNew()
	if err != nil {
		return nil, fmt.Errorf("allocate tables page: %w", err)
	}

	// store page count
	if err := headerPage.StoreRecordCell(page.RecordCell{
		Key:    []byte(HeaderPageCount),
		Record: encodeUint64(3), // header, config and tables page
	}); err != nil {
		return nil, fmt.Errorf("store page count: %w", err)
	}
	// store pointer to config page
	if err := headerPage.StorePointerCell(page.PointerCell{
		Key:     []byte(HeaderConfig),
		Pointer: configPage.ID(),
	}); err != nil {
		return nil, fmt.Errorf("store config pointer: %w", err)
	}
	// store pointer to tables page
	if err := headerPage.StorePointerCell(page.PointerCell{
		Key:     []byte(HeaderTables),
		Pointer: tablesPage.ID(),
	}); err != nil {
		return nil, fmt.Errorf("store tables pointer: %w", err)
	}

	err = mgr.WritePage(headerPage) // immediately flush
	if err != nil {
		return nil, fmt.Errorf("write header page: %w", err)
	}
	err = mgr.WritePage(configPage) // immediately flush
	if err != nil {
		return nil, fmt.Errorf("write config page: %w", err)
	}
	err = mgr.WritePage(tablesPage) // immediately flush
	if err != nil {
		return nil, fmt.Errorf("write tables page: %w", err)
	}

	return newDB(file, mgr, headerPage, opts...)
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

	return newDB(file, mgr, headerPage, opts...)
}

// AllocateNewPage allocates and immediately persists a new page in secondary
// storage. This will fail if the DBFile is closed. After this method returns,
// the allocated page can immediately be found by the cache (it is not loaded
// yet), and you can use the returned page ID to load the page through the
// cache.
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
// the file. Everything will be closed after writing the config and header page.
func (db *DBFile) Close() error {
	if err := db.pageManager.WritePage(db.headerPage); err != nil {
		return fmt.Errorf("write header page: %w", err)
	}
	if err := db.pageManager.WritePage(db.configPage); err != nil {
		return fmt.Errorf("write config page: %w", err)
	}
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
func newDB(file afero.File, mgr *PageManager, headerPage *page.Page, opts ...Option) (*DBFile, error) {
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

	if err := db.initialize(); err != nil {
		return nil, fmt.Errorf("initialize: %w", err)
	}

	db.cache = cache.NewLRUCache(db.cacheSize, mgr)

	return db, nil
}

func (db *DBFile) initialize() error {
	// get config page id
	cfgPageID, err := pointerCellValue(db.headerPage, HeaderConfig)
	if err != nil {
		return err
	}

	// read config page
	cfgPage, err := db.pageManager.ReadPage(cfgPageID)
	if err != nil {
		return fmt.Errorf("can't read config page: %w", err)
	}
	db.configPage = cfgPage

	return nil
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

func pointerCellValue(p *page.Page, cellKey string) (page.ID, error) {
	cell, ok := p.Cell([]byte(cellKey))
	if !ok {
		return 0, ErrNoSuchCell(cellKey)
	}
	if cell.Type() != page.CellTypePointer {
		return 0, fmt.Errorf("cell '%v' is %v, which is not a pointer cell", cellKey, cell.Type())
	}
	return cell.(page.PointerCell).Pointer, nil
}
