package page

// Header is an enum type for header fields.
type Header uint8

const (
	// HeaderVersion is the version number, that a page is in. This is a uint16.
	HeaderVersion Header = iota
	// HeaderID is the page ID, which may be used outside of the page for
	// housekeeping. This is a uint16.
	HeaderID
)

// ID is the type of a page ID. This is mainly to avoid any confusion.
// Changing this will break existing database files, so only change during major
// version upgrades.
type ID = uint32

// Page describes a memory page that stores (page.Cell)s. A page consists of
// header fields and cells, and is a plain store. Obtained cells are always
// ordered ascending by the cell key. A page supports variable size cell keys
// and records. A page is generally NOT safe for concurrent writes.
type Page interface {
	// Version returns the version of the page layout. Use this for choosing the
	// page implementation to use to decode the page.
	Version() uint32

	// ID returns the page ID, as it is used by any page loader. It is unique in
	// the scope of one database.
	ID() ID

	// Dirty determines whether this page has been modified since the last time
	// Page.ClearDirty was called.
	Dirty() bool
	// MarkDirty marks this page as dirty.
	MarkDirty()
	// ClearDirty unsets the dirty flag from this page.
	ClearDirty()

	// StorePointerCell stores the given pointer cell in the page.
	//
	// If a cell with the same key as the given pointer already exists in the
	// page, it will be overwritten.
	//
	// If a cell with the same key as the given cell does NOT already exist, it
	// will be created.
	//
	// To change the type of a cell, delete it and store a new cell.
	StorePointerCell(PointerCell) error
	// StoreRecordCell stores the given record cell in the page.
	//
	// If a cell with the same key as the given cell already exists in the page,
	// it will be overwritten.
	//
	// If a cell with the same key as the given pointer does NOT already exist,
	// it will be created.
	//
	// To change the type of a cell, delete it and store a new cell.
	StoreRecordCell(RecordCell) error
	// Delete deletes the cell with the given bytes as key. If the key couldn't
	// be found, false is returned. If an error occured during deletion, the
	// error is returned.
	DeleteCell([]byte) (bool, error)
	// Cell returns a cell with the given key, together with a bool indicating
	// whether any cell in the page has that key. Use a switch statement to
	// determine which type of cell you just obtained (pointer, record).
	Cell([]byte) (Cell, bool)
	// Cells returns all cells in this page as a slice. Cells are ordered
	// ascending by key. Calling this method can be expensive since all cells
	// have to be decoded.
	Cells() []Cell
}

// Cell describes a generic page cell that holds a key. Use a switch statement
// to determine the type of the cell.
type Cell interface {
	// Key returns the key of the cell.
	Key() []byte
}

// PointerCell describes a cell that points to another page in memory.
type PointerCell interface {
	Cell
	// Pointer returns the page ID of the child page that this cell points to.
	Pointer() ID
}

// RecordCell describes a cell that holds some kind of value. What value format
// was used is none of the cells concern, just use it as what you put in.
type RecordCell interface {
	Cell
	// Record is the data record in this cell, returned as a byte slice.
	Record() []byte
}
