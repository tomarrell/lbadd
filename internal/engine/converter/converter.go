package converter

import (
	"encoding/binary"
	"math"
)

const (
	falseByte byte = byte(0)
	trueByte  byte = ^falseByte
)

var (
	byteOrder = binary.BigEndian
)

// bool

// BoolToByte converts the given bool to a byte which is then returned.
func BoolToByte(v bool) byte {
	if v {
		return trueByte
	}
	return falseByte
}

// ByteToBool converts the given byte back to a bool.
func ByteToBool(v byte) bool {
	return v != falseByte
}

// BoolToByteArray converts the given bool to a [1]byte is then returned.
func BoolToByteArray(v bool) [1]byte {
	return [1]byte{BoolToByte(v)}
}

// ByteArrayToBool converts the given [1]byte back to a bool.
func ByteArrayToBool(v [1]byte) bool {
	return ByteToBool(v[0])
}

// BoolToByteSlice converts the given bool to a []byte which is then returned.
func BoolToByteSlice(v bool) []byte {
	arr := BoolToByteArray(v)
	return arr[:]
}

// ByteSliceToBool converts the given []byte back to a bool.
func ByteSliceToBool(v []byte) bool {
	return ByteToBool(v[0])
}

// integral

// Uint16ToByteArray converts the given uint16 to a [2]byte which is then
// returned.
func Uint16ToByteArray(v uint16) (result [2]byte) {
	byteOrder.PutUint16(result[:], v)
	return
}

// ByteArrayToUint16 converts the given [2]byte back to a uint16.
func ByteArrayToUint16(v [2]byte) uint16 {
	return byteOrder.Uint16(v[:])
}

// Uint16ToByteSlice converts the given uint16 to a []byte which is then
// returned.
func Uint16ToByteSlice(v uint16) (result []byte) {
	result = make([]byte, 2)
	byteOrder.PutUint16(result, v)
	return
}

// ByteSliceToUint16 converts the given []byte back to a uint16.
func ByteSliceToUint16(v []byte) uint16 {
	return byteOrder.Uint16(v)
}

// Uint32ToByteArray converts the given uint32 to a [4]byte which is then
// returned.
func Uint32ToByteArray(v uint32) (result [4]byte) {
	byteOrder.PutUint32(result[:], v)
	return
}

// ByteArrayToUint32 converts the given [4]byte back to a uint32.
func ByteArrayToUint32(v [4]byte) uint32 {
	return byteOrder.Uint32(v[:])
}

// Uint32ToByteSlice converts the given uint32 to a []byte which is then
// returned.
func Uint32ToByteSlice(v uint32) (result []byte) {
	result = make([]byte, 4)
	byteOrder.PutUint32(result, v)
	return
}

// ByteSliceToUint32 converts the given []byte back to a uint32.
func ByteSliceToUint32(v []byte) uint32 {
	return byteOrder.Uint32(v)
}

// Uint64ToByteArray converts the given uint64 to a [8]byte which is then
// returned.
func Uint64ToByteArray(v uint64) (result [8]byte) {
	byteOrder.PutUint64(result[:], v)
	return
}

// ByteArrayToUint64 converts the given [8]byte back to a uint64.
func ByteArrayToUint64(v [8]byte) uint64 {
	return byteOrder.Uint64(v[:])
}

// Uint64ToByteSlice converts the given uint64 to a []byte which is then
// returned.
func Uint64ToByteSlice(v uint64) (result []byte) {
	result = make([]byte, 8)
	byteOrder.PutUint64(result, v)
	return
}

// ByteSliceToUint64 converts the given []byte back to a uint64.
func ByteSliceToUint64(v []byte) uint64 {
	return byteOrder.Uint64(v)
}

// fractal

// Float32ToByteArray converts the given float32 to a [4]byte, which is then
// returned.
func Float32ToByteArray(v float32) (result [4]byte) {
	return Uint32ToByteArray(math.Float32bits(v))
}

// Float32ToByteSlice converts the given float32 to a []byte, which is then
// returned.
func Float32ToByteSlice(v float32) (result []byte) {
	return Uint32ToByteSlice(math.Float32bits(v))
}

// Float64ToByteArray converts the given float64 to a [8]byte, which is then
// returned.
func Float64ToByteArray(v float64) (result [8]byte) {
	return Uint64ToByteArray(math.Float64bits(v))
}

// Float64ToByteSlice converts the given float64 to a []byte, which is then
// returned.
func Float64ToByteSlice(v float64) (result []byte) {
	return Uint64ToByteSlice(math.Float64bits(v))
}

// complex

// Complex64ToByteArray converts the given complex64 to a [8]byte, which is then
// returned.
func Complex64ToByteArray(v complex64) (result [8]byte) {
	copy(result[:4], Float32ToByteSlice(real(v)))
	copy(result[4:], Float32ToByteSlice(imag(v)))
	return
}

// Complex64ToByteSlice converts the given complex64 to a []byte, which is then
// returned.
func Complex64ToByteSlice(v complex64) (result []byte) {
	result = make([]byte, 8)
	copy(result[:4], Float32ToByteSlice(real(v)))
	copy(result[4:], Float32ToByteSlice(imag(v)))
	return
}

// Complex128ToByteArray converts the given complex128 to a [16]byte, which is
// then returned.
func Complex128ToByteArray(v complex128) (result [16]byte) {
	copy(result[:8], Float64ToByteSlice(real(v)))
	copy(result[8:], Float64ToByteSlice(imag(v)))
	return
}

// Complex128ToByteSlice converts the given complex128 to a []byte, which is
// then returned.
func Complex128ToByteSlice(v complex128) (result []byte) {
	result = make([]byte, 16)
	copy(result[:8], Float64ToByteSlice(real(v)))
	copy(result[8:], Float64ToByteSlice(imag(v)))
	return
}

// variable-size

// StringToByteSlice converts the given string a []byte which is then returned.
func StringToByteSlice(v string) []byte {
	return []byte(v)
}

// ByteSliceToString converts the given []byte back to a string.
func ByteSliceToString(v []byte) string {
	return string(v)
}
