package record

import (
	fn "github.com/myuksal/sshhpp/shapefile/function"
)

type PointContent struct {
	X float64
	Y float64
}

/**
## Point (16 bytes)

| Byte Position  | Field | Value | Type   | Endianness |
|----------------|-------|-------|--------|------------|
| 0              | X     | X     | Double | Little     |
| 8              | Y     | Y     | Double | Little     |
*/

func CreatePointContent(bytes []byte) PointContent {
	return PointContent{
		fn.LittleEndianFloat64(bytes[0:8]),
		fn.LittleEndianFloat64(bytes[8:16]),
	}
}
