package types

import "encoding/binary"

var (
	byteOrder = binary.BigEndian
)

// Serializer wraps two basic methods for two-way serializing values. Which
// values can be serialized and deserialized, is up to the implementing object.
// It must be documented, what can and can not be serialized and deserialized.
type Serializer interface {
	Serialize(Value) ([]byte, error)
	Deserialize([]byte) (Value, error)
}

// frame prepends an 4 byte big endian encoded unsigned integer to the given
// data, which represents the length of the data.
func frame(data []byte) []byte {
	size := len(data)
	result := make([]byte, 4+len(data))
	byteOrder.PutUint32(result, uint32(size))
	copy(result[4:], data)
	return result
}
