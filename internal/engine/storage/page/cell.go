package page

import (
	"bytes"
)

var _ CellTyper = (*RecordCell)(nil)
var _ CellTyper = (*PointerCell)(nil)

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
	// CellTyper describes a component that has a cell type.
	CellTyper interface {
		Type() CellType
	}

	// RecordCell is a cell with CellTypeRecord. It holds a key and a variable
	// size record.
	RecordCell struct {
		Key    []byte
		Record []byte
	}

	// PointerCell is a cell with CellTypePointer. It holds a key and an uint32,
	// pointing to another page.
	PointerCell struct {
		Key     []byte
		Pointer ID
	}
)

// Type returns CellTypeRecord.
func (RecordCell) Type() CellType { return CellTypeRecord }

// Type returns CellTypePointer.
func (PointerCell) Type() CellType { return CellTypePointer }

func decodeCell(data []byte) CellTyper {
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
	key := frame(cell.Key)
	record := frame(cell.Record)

	var buf bytes.Buffer
	buf.WriteByte(byte(CellTypeRecord))
	buf.Write(key)
	buf.Write(record)

	return buf.Bytes()
}

func decodeRecordCell(data []byte) RecordCell {
	keySize := byteOrder.Uint32(data[1:5])
	key := data[5 : 5+keySize]
	recordSize := byteOrder.Uint32(data[5+keySize : 5+keySize+4])
	record := data[5+keySize+4 : 5+keySize+4+recordSize]
	return RecordCell{
		Key:    key,
		Record: record,
	}
}

func encodePointerCell(cell PointerCell) []byte {
	key := frame(cell.Key)
	pointer := make([]byte, 4)
	byteOrder.PutUint32(pointer, cell.Pointer)

	var buf bytes.Buffer
	buf.WriteByte(byte(CellTypePointer))
	buf.Write(key)
	buf.Write(pointer)

	return buf.Bytes()
}

func decodePointerCell(data []byte) PointerCell {
	keySize := byteOrder.Uint32(data[1:5])
	key := data[5 : 5+keySize]
	pointer := byteOrder.Uint32(data[5+keySize : 5+keySize+4])
	return PointerCell{
		Key:     key,
		Pointer: pointer,
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
