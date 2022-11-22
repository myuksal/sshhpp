package shapefile

/**
## Main file header (100 bytes)

| Byte Position  | Field        | Value       | Type    | Endianness |
|----------------|--------------|-------------|---------|------------|
| 0              | File Code    | 9994        | Integer | Big        |
| 4              | Unused       | 0           | Integer | Big        |
| 8              | Unused       | 0           | Integer | Big        |
| 12             | Unused       | 0           | Integer | Big        |
| 16             | Unused       | 0           | Integer | Big        |
| 20             | Unused       | 0           | Integer | Big        |
| 24             | File Length  | File Length | Integer | Big        |
| 28             | Version      | 1000        | Integer | Little     |
| 32             | Shape Type   | Shape Type  | Integer | Little     |
| 36             | Bounding Box | X min       | Double  | Little     |
| 44             | Bounding Box | Y min       | Double  | Little     |
| 52             | Bounding Box | X max       | Double  | Little     |
| 60             | Bounding Box | Y max       | Double  | Little     |
| ~~68~~         | Bounding Box | Z min       | Double  | Little     |
| ~~76~~         | Bounding Box | Z max       | Double  | Little     |
| ~~84~~         | Bounding Box | M min       | Double  | Little     |
| ~~92~~         | Bounding Box | M max       | Double  | Little     |

> **Byte Position** ~~number~~ is Unused with value 0.0. if not Measured or Z type.
*/

import (
	"fmt"

	fn "github.com/myuksal/sshhpp/function"
	types "github.com/myuksal/sshhpp/shapefile/type"
)

type ShapeHeader struct {
	Code       uint32
	FileLength uint32
	Version    uint32
	Shape      types.ShapeType
	XMin       float64
	YMin       float64
	XMax       float64
	YMax       float64
	zMin       float64
	zMax       float64
	mMin       float64
	mMax       float64
}

func (header ShapeHeader) String() string {
	title := fmt.Sprintf("|%-12s|%-8s|%-6s|%-10s|\n", "code", "version", "shape", "length")
	divider := "|------------|--------|------|----------|\n"
	data := fmt.Sprintf("|%-12d|%-8d|%-6d|%-10d|\n", header.Code, header.Version, header.Shape, header.FileLength)
	return title + divider + data
}

func CreateShapeHeader(bytes []byte) ShapeHeader {
	return ShapeHeader{
		fn.LittleEndianUInt32(bytes[0:4]),                    // file code
		fn.LittleEndianUInt32(bytes[24:28]),                  // file length
		fn.LittleEndianUInt32(bytes[28:32]),                  // version
		types.ShapeType(fn.LittleEndianUInt32(bytes[32:36])), // shape
		fn.LittleEndianFloat64(bytes[36:44]),                 // X min
		fn.LittleEndianFloat64(bytes[44:52]),                 // Y min
		fn.LittleEndianFloat64(bytes[52:60]),                 // X max
		fn.LittleEndianFloat64(bytes[60:68]),                 // Y max
		fn.LittleEndianFloat64(bytes[68:76]),                 // Z min
		fn.LittleEndianFloat64(bytes[76:84]),                 // Z max
		fn.LittleEndianFloat64(bytes[84:92]),                 // M min
		fn.LittleEndianFloat64(bytes[92:100]),                // M max
	}
}
