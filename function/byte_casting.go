package function

import (
	"bytes"
	"encoding/binary"
	"math"
)

func LittleEndianFloat64(bytes []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(bytes))
}

func BigEndianFloat64(bytes []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(bytes))
}

func LittleEndianUInt32(bytes []byte) uint32 {
	return binary.LittleEndian.Uint32(bytes)
}

func LittleEndianUInt16(bytes []byte) uint16 {
	return binary.LittleEndian.Uint16(bytes)
}

func UInt8(bytes byte) uint8 {
	return uint8(bytes)
}

func BigEndianUInt32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}

func Ascii(d []byte) string {
	return string(bytes.Trim(d, "\x00"))
}
