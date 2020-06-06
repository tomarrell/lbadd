package page

// Header is an enum type for header fields.
type Header uint8

const (
	// HeaderVersion is the version number, that a page is in. This is a uint16.
	HeaderVersion Header = iota
	// HeaderID is the page ID, which may be used outside of the page for
	// housekeeping. This is a uint16.
	HeaderID
	// HeaderCellCount is the amount of cells that are currently stored in the
	// page. This is a uint16.
	HeaderCellCount
)

// Loader is a function that can load a page from a given byte slice, and return
// errors if any occur.
type Loader func([]byte) (Page, error)

// Page describes a memory page that stores (page.Cell)s. A page consists of
// header fields and cells, and is a plain store. Obtained cells are always
// ordered ascending by the cell key. A page supports variable size cell keys
// and records.
type Page interface {
	// Header obtains a header field from the page's header. If the header is
	// not supported, a result=nil,error=ErrUnknownHeader is returned. The type
	// of the returned header field value is documented in the header key's
	// godoc.
	Header(Header) (interface{}, error)

	// StoreCell stores a cell on the page. If the page does not fit the cell
	// because of size or too much fragmentation, an error will be returned.
	StoreCell(Cell) error
	// Delete deletes the cell with the given bytes as key. If the key couldn't
	// be found, nil is returned. If an error occured during deletion, the error
	// is returned.
	Delete([]byte) error
	// Cell returns a cell with the given key, together with a bool indicating
	// whether any cell in the page has that key.
	Cell([]byte) (Cell, bool)
	// Cells returns all cells in this page as a slice. Cells are ordered
	// ascending by key. Calling this method is relatively cheap.
	Cells() []Cell
}

// Cell is a structure that represents a key-value cell. Both the key and the
// record can be of variable size.
type Cell struct {
	// Key is the key of this cell, used for ordering.
	Key []byte
	// Record is the data stored inside the cell.
	Record []byte
}
