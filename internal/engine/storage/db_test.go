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

	db, err := Create(f)
	assert.NoError(err)

	mustHaveSize(assert, f, 2*page.Size)

	headerPage := loadPageFromOffset(assert, f, 0)
	assert.EqualValues(2, headerPage.CellCount())
	assert.Equal(
		encodeUint64(2),
		mustCell(assert, headerPage, HeaderPageCount).(page.RecordCell).Record,
	)

	_, err = db.AllocateNewPage()
	assert.NoError(err)

	mustHaveSize(assert, f, 3*page.Size)

	headerPage = loadPageFromOffset(assert, f, 0)
	assert.EqualValues(2, headerPage.CellCount())
	assert.Equal(
		encodeUint64(3),
		mustCell(assert, headerPage, HeaderPageCount).(page.RecordCell).Record,
	)

	assert.NoError(db.Close())
}

func mustHaveSize(assert *assert.Assertions, file afero.File, expectedSize int64) {
	stat, err := file.Stat()
	assert.NoError(err)
	assert.EqualValues(expectedSize, stat.Size())
}

func mustCell(assert *assert.Assertions, p *page.Page, key string) interface{} {
	val, ok := p.Cell([]byte(key))
	assert.Truef(ok, "page must have cell with key %v", key)
	return val
}

func loadPageFromOffset(assert *assert.Assertions, rd io.ReaderAt, off int64) *page.Page {
	buf := make([]byte, page.Size)
	_, err := rd.ReadAt(buf, off)
	assert.NoError(err)
	p, err := page.Load(buf)
	assert.NoError(err)
	return p
}
