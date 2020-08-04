package storage

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

type PageManagerSuite struct {
	suite.Suite

	file afero.File
}

func (suite *PageManagerSuite) SetupTest() {
	fs := afero.NewMemMapFs()
	// empty file
	file, err := fs.Create("mydbfile")
	suite.NoError(err)

	setupMgr, err := NewPageManager(file)
	suite.NoError(err)

	// allocate three pages with id 0 through 2
	_, err = setupMgr.AllocateNew()
	suite.NoError(err)
	_, err = setupMgr.AllocateNew()
	suite.NoError(err)
	_, err = setupMgr.AllocateNew()
	suite.NoError(err)

	// modify id of page with id 1
	_, err = file.WriteAt([]byte{0x00, 0x00, 0x00, 0x03}, page.Size)
	suite.NoError(err)

	suite.file = file
}

func (suite *PageManagerSuite) TestPageManager_loadPageIDs() {
	testMgr, err := NewPageManager(suite.file)
	suite.NoError(err)

	// test the page manager's loaded page offsets
	suite.EqualValues(0*page.Size, testMgr.pageOffsets[0])
	suite.EqualValues(1*page.Size, testMgr.pageOffsets[3])
	suite.EqualValues(2*page.Size, testMgr.pageOffsets[2])
	suite.Len(testMgr.pageOffsets, 3)

	var p *page.Page
	p, err = testMgr.ReadPage(0)
	suite.NoError(err)
	suite.EqualValues(0, p.ID())
	p, err = testMgr.ReadPage(2)
	suite.NoError(err)
	suite.EqualValues(2, p.ID())
	p, err = testMgr.ReadPage(3)
	suite.NoError(err)
	suite.EqualValues(3, p.ID())
}

func (suite *PageManagerSuite) TestPageManager_AllocateNew() {
	fs := afero.NewMemMapFs()
	// empty file
	file, err := fs.Create("mydbfile")
	suite.NoError(err)

	mgr, err := NewPageManager(file)
	suite.NoError(err)
	defer func() { _ = mgr.Close() }()

	var p []*page.Page
	for i := 0; i < 5; i++ {
		newPage, err := mgr.AllocateNew()
		suite.NoError(err)
		p = append(p, newPage)
	}

	for i := 0; i < len(p); i++ {
		suite.EqualValues(i, p[i].ID())
	}
}

func TestPageManagerSuite(t *testing.T) {
	suite.Run(t, new(PageManagerSuite))
}
