package function_test

import (
	"encoding/binary"
	"math"
	"testing"
	"unsafe"

	fn "github.com/myuksal/sshhpp/function"
	"github.com/stretchr/testify/assert"
)

func toBytes(i any) []byte {
	return *(*[]byte)(unsafe.Pointer(&i))
}

func TestLittleEndianFloat64(t *testing.T) {
	assert := assert.New(t)

	var buffer [8]byte
	binary.LittleEndian.PutUint64(buffer[:], math.Float64bits(123.4567))

	result := fn.LittleEndianFloat64(buffer[:])
	assert.Equal(result, float64(123.4567))
}
func TestBigEndianFloat64(t *testing.T) {
	assert := assert.New(t)

	var buffer [8]byte
	binary.BigEndian.PutUint64(buffer[:], math.Float64bits(123.4567))

	result := fn.BigEndianFloat64(buffer[:])
	assert.Equal(result, float64(123.4567))
}
func TestLittleEndianUInt32(t *testing.T) {
	assert := assert.New(t)

	var buffer [4]byte
	binary.LittleEndian.PutUint32(buffer[:], 123456)

	result := fn.LittleEndianUInt32(buffer[:])
	assert.Equal(result, uint32(123456))
}

func TestBigEndianUInt32(t *testing.T) {
	assert := assert.New(t)

	var buffer [4]byte
	binary.BigEndian.PutUint32(buffer[:], 123456)

	result := fn.BigEndianUInt32(buffer[:])
	assert.Equal(result, uint32(123456))
}
func TestLittleEndianUInt16(t *testing.T) {
	assert := assert.New(t)

	var buffer [4]byte
	binary.LittleEndian.PutUint16(buffer[:], 12345)

	result := fn.LittleEndianUInt16(buffer[:])
	assert.Equal(result, uint16(12345))
}
func TestUInt8(t *testing.T) {
	assert := assert.New(t)
	result := fn.UInt8(byte(123))
	assert.Equal(result, uint8(123))
}
func TestAscii(t *testing.T) {

	assert := assert.New(t)
	result := fn.Ascii([]byte{'A', 'B', 'C'})
	assert.Equal(result, "ABC")
}
