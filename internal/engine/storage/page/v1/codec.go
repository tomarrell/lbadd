// This file contains encoding and decoding functions.

package v1

import (
	"github.com/tomarrell/lbadd/internal/engine/converter"
	"github.com/tomarrell/lbadd/internal/engine/storage/page"
)

// encodeCell frames the cell key and record, and concatenates them. The
// concatenation is NOT framed itself.
//
//  encoded = | frame | key | frame | record |
func encodeCell(cell page.Cell) []byte {
	keyFrame := frameData(cell.Key)
	recordFrame := frameData(cell.Record)
	return append(keyFrame, recordFrame...)
}

// decodeCell takes data of the format that encodeCell produced, and convert it
// back into a (page.Cell).
func decodeCell(data []byte) page.Cell {
	keyDataSize := converter.ByteSliceToUint16(data[:FrameSizeSize])
	keyLo := FrameSizeSize
	keyHi := FrameSizeSize + keyDataSize

	dataLo := keyHi + FrameSizeSize
	dataSize := converter.ByteSliceToUint16(data[keyHi:dataLo])
	dataHi := dataLo + dataSize

	return page.Cell{
		Key:    data[keyLo:keyHi],
		Record: data[dataLo:dataHi],
	}
}

// frameData frames the given data and returns the framed bytes. The frame
// contains the length of the data, as uint16.
//
//	framed = | length | data |
func frameData(data []byte) []byte {
	return append(converter.Uint16ToByteSlice(uint16(len(data))), data...)
}

// encodeOffset converts the offset to a byte slice of the following form.
//
//  encoded = | location | size |
//
// Both the location and size will be encoded as 2 byte uint16.
func encodeOffset(offset Offset) []byte {
	return append(converter.Uint16ToByteSlice(offset.Location), converter.Uint16ToByteSlice(offset.Size)...)
}

// decodeOffset takes bytes of the form produced by encodeOffset and converts it
// back into an Offset.
func decodeOffset(data []byte) Offset {
	return Offset{
		Location: converter.ByteSliceToUint16(data[:int(OffsetSize)/2]),
		Size:     converter.ByteSliceToUint16(data[int(OffsetSize)/2 : int(OffsetSize)]),
	}
}

func zero(b []byte) {
	for i := range b {
		b[i] = 0x00
	}
}
