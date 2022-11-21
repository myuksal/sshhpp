package function

import (
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

func BigEndianUInt32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}
