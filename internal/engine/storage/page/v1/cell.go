package v1

import (
	"bytes"

	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

var _ page.Cell = (*RecordCell)(nil)
var _ page.Cell = (*PointerCell)(nil)

//go:generate stringer -type=CellType

// CellType is the type of a page.
type CellType uint8

const (
	// CellTypeUnknown indicates a corrupted page or an incorrectly decoded
	// cell.
	CellTypeUnknown CellType = iota
	// CellTypeRecord indicates a RecordCell, which stores a key and a variable
	// size record.
	CellTypeRecord
	// CellTypePointer indicates a PointerCell, which stores a key and an
	// uint32, which points to another page.
	CellTypePointer
)

type (
	// Cell is a cell that has a type and a key.
	Cell interface {
		page.Cell
		Type() CellType
	}

	cell struct {
		key []byte
	}

	// RecordCell is a cell with CellTypeRecord. It holds a key and a variable
	// size record.
	RecordCell struct {
		cell
		record []byte
	}

	// PointerCell is a cell with CellTypePointer. It holds a key and an uint32,
	// pointing to another page.
	PointerCell struct {
		cell
		pointer page.ID
	}
)

// Key returns the key of this cell.
func (c cell) Key() []byte { return c.key }

// Record returns the record data of this cell.
func (c RecordCell) Record() []byte { return c.record }

// Pointer returns the pointer of this page, that points to another page.
func (c PointerCell) Pointer() page.ID { return c.pointer }

// Type returns CellTypeRecord.
func (c RecordCell) Type() CellType { return CellTypeRecord }

// Type returns CellTypePointer.
func (c PointerCell) Type() CellType { return CellTypePointer }

func decodeCell(data []byte) Cell {
	switch t := CellType(data[0]); t {
	case CellTypePointer:
		return decodePointerCell(data)
	case CellTypeRecord:
		return decodeRecordCell(data)
	default:
		return nil
	}
}

func encodeRecordCell(cell RecordCell) []byte {
	key := frame(cell.key)
	record := frame(cell.record)

	var buf bytes.Buffer
	buf.WriteByte(byte(CellTypeRecord))
	buf.Write(key)
	buf.Write(record)

	return buf.Bytes()
}

func decodeRecordCell(data []byte) RecordCell {
	cp := copySlice(data)

	keySize := byteOrder.Uint32(cp[1:5])
	key := cp[5 : 5+keySize]
	recordSize := byteOrder.Uint32(cp[5+keySize : 5+keySize+4])
	record := cp[5+keySize+4 : 5+keySize+4+recordSize]
	return RecordCell{
		cell: cell{
			key: key,
		},
		record: record,
	}
}

func encodePointerCell(cell PointerCell) []byte {
	key := frame(cell.key)
	pointer := make([]byte, 4)
	byteOrder.PutUint32(pointer, cell.pointer)

	var buf bytes.Buffer
	buf.WriteByte(byte(CellTypePointer))
	buf.Write(key)
	buf.Write(pointer)

	return buf.Bytes()
}

func decodePointerCell(data []byte) PointerCell {
	cp := copySlice(data)

	keySize := byteOrder.Uint32(cp[1:5])
	key := cp[5 : 5+keySize]
	pointer := byteOrder.Uint32(cp[5+keySize : 5+keySize+4])
	return PointerCell{
		cell: cell{
			key: key,
		},
		pointer: pointer,
	}
}

func frame(data []byte) []byte {
	// this allocation can be optimized, however, it would mess up the API, but
	// it should be considered in the future
	result := make([]byte, 4+len(data))
	copy(result[4:], data)
	byteOrder.PutUint32(result, uint32(len(data)))
	return result
}

func copySlice(original []byte) []byte {
	copied := make([]byte, len(original))
	copy(copied, original)
	return copied
}
