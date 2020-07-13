package storage

import (
	"io"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	fs := afero.NewMemMapFs()

	f, err := fs.Create("mydbfile")
	assert.NoError(err)

	// actual tests

	// create the database file contents in file f
	db, err := Create(f)
	assert.NoError(err)

	// f must have the size of 3 pages, header, tables and config page
	mustHaveSize(assert, f, 3*page.Size)

	// load the header page, which is the first page (offset 0) in the file
	headerPage := loadPageFromOffset(assert, f, 0)
	assert.EqualValues(3, headerPage.CellCount())
	assert.Equal(
		encodeUint64(3),
		mustCell(assert, headerPage, HeaderPageCount).(page.RecordCell).Record,
	)

	// Allocating a new page must persist it is in the created database file. This
	// check ensures, that the file is writable.
	_, err = db.AllocateNewPage()
	assert.NoError(err)

	// after allocating a new page, the file must have grown to 4 times the size
	// of a single page
	mustHaveSize(assert, f, 4*page.Size)

	// check the header page again, which must have the same amount of cells,
	// but the page count cell value must have been incremented by 1
	headerPage = loadPageFromOffset(assert, f, 0)
	assert.EqualValues(3, headerPage.CellCount())
	assert.Equal(
		encodeUint64(4),
		mustCell(assert, headerPage, HeaderPageCount).(page.RecordCell).Record,
	)

	assert.NoError(db.Close())
}

func mustHaveSize(assert *assert.Assertions, file afero.File, expectedSize int64) {
	stat, err := file.Stat()
	assert.NoError(err)
	assert.EqualValues(expectedSize, stat.Size())
}

func mustCell(assert *assert.Assertions, p *page.Page, key string) page.CellTyper {
	val, ok := p.Cell([]byte(key))
	assert.Truef(ok, "page must have cell with key %v", key)
	return val
}

func loadPageFromOffset(assert *assert.Assertions, rd io.ReaderAt, off int64) *page.Page {
	buf := make([]byte, page.Size)
	n, err := rd.ReadAt(buf, off)
	assert.Equal(len(buf), n)
	assert.NoError(err)
	p, err := page.Load(buf)
	assert.NoError(err)
	return p
}
