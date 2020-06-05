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

// BoolToByte converts the given argument to a byte output, which is then
// returned.
func BoolToByte(v bool) byte {
	if v {
		return trueByte
	}
	return falseByte
}

// BoolToByteArray converts the given argument to a byte output, which is then
// returned.
func BoolToByteArray(v bool) [1]byte {
	return [1]byte{BoolToByte(v)}
}

// BoolToByteSlice converts the given argument to a byte output, which is then
// returned.
func BoolToByteSlice(v bool) []byte {
	arr := BoolToByteArray(v)
	return arr[:]
}

// integral

// Uint16ToByteArray converts the given argument to a byte output, which is then
// returned.
func Uint16ToByteArray(v uint16) (result [2]byte) {
	byteOrder.PutUint16(result[:], v)
	return
}

// Uint16ToByteSlice converts the given argument to a byte output, which is then
// returned.
func Uint16ToByteSlice(v uint16) (result []byte) {
	result = make([]byte, 2)
	byteOrder.PutUint16(result, v)
	return
}

// Uint32ToByteArray converts the given argument to a byte output, which is then
// returned.
func Uint32ToByteArray(v uint32) (result [4]byte) {
	byteOrder.PutUint32(result[:], v)
	return
}

// Uint32ToByteSlice converts the given argument to a byte output, which is then
// returned.
func Uint32ToByteSlice(v uint32) (result []byte) {
	result = make([]byte, 4)
	byteOrder.PutUint32(result, v)
	return
}

// Uint64ToByteArray converts the given argument to a byte output, which is then
// returned.
func Uint64ToByteArray(v uint64) (result [8]byte) {
	byteOrder.PutUint64(result[:], v)
	return
}

// Uint64ToByteSlice converts the given argument to a byte output, which is then
// returned.
func Uint64ToByteSlice(v uint64) (result []byte) {
	result = make([]byte, 8)
	byteOrder.PutUint64(result, v)
	return
}

// fractal

// Float32ToByteArray converts the given argument to a byte output, which is
// then returned.
func Float32ToByteArray(v float32) (result [4]byte) {
	return Uint32ToByteArray(math.Float32bits(v))
}

// Float32ToByteSlice converts the given argument to a byte output, which is
// then returned.
func Float32ToByteSlice(v float32) (result []byte) {
	return Uint32ToByteSlice(math.Float32bits(v))
}

// Float64ToByteArray converts the given argument to a byte output, which is
// then returned.
func Float64ToByteArray(v float64) (result [8]byte) {
	return Uint64ToByteArray(math.Float64bits(v))
}

// Float64ToByteSlice converts the given argument to a byte output, which is
// then returned.
func Float64ToByteSlice(v float64) (result []byte) {
	return Uint64ToByteSlice(math.Float64bits(v))
}

// complex

// Complex64ToByteArray converts the given argument to a byte output, which is
// then returned.
func Complex64ToByteArray(v complex64) (result [8]byte) {
	copy(result[:4], Float32ToByteSlice(real(v)))
	copy(result[4:], Float32ToByteSlice(imag(v)))
	return
}

// Complex64ToByteSlice converts the given argument to a byte output, which is
// then returned.
func Complex64ToByteSlice(v complex64) (result []byte) {
	result = make([]byte, 8)
	copy(result[:4], Float32ToByteSlice(real(v)))
	copy(result[4:], Float32ToByteSlice(imag(v)))
	return
}

// Complex128ToByteArray converts the given argument to a byte output, which is
// then returned.
func Complex128ToByteArray(v complex128) (result [16]byte) {
	copy(result[:8], Float64ToByteSlice(real(v)))
	copy(result[8:], Float64ToByteSlice(imag(v)))
	return
}

// Complex128ToByteSlice converts the given argument to a byte output, which is
// then returned.
func Complex128ToByteSlice(v complex128) (result []byte) {
	result = make([]byte, 16)
	copy(result[:8], Float64ToByteSlice(real(v)))
	copy(result[8:], Float64ToByteSlice(imag(v)))
	return
}

// variable-size

// StringToByteSlice converts the given argument to a byte output, which is then
// returned.
func StringToByteSlice(v string) []byte {
	return []byte(v)
}
