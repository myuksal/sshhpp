package record

import (
	fn "github.com/myuksal/sshhpp/shapefile/function"
)

type PolyLineContent struct {
	XMin  float64
	YMin  float64
	XMax  float64
	YMax  float64
	Parts [][]PointContent
}

/**
## Poly line (60~ bytes)

| Byte Position                   | Field            | Type    | Endianness |
|---------------------------------|------------------|---------|------------|
| 0                               | X min            | Double  | Little     |
| 8                               | Y min            | Double  | Little     |
| 16                              | Y max            | Double  | Little     |
| 24                              | Y max            | Double  | Little     |
| 32                              | Number of parts  | Integer | Little     |
| 36                              | Number of points | Integer | Little     |
| 40, 40+4*num                    | Parts            | Integer | Little     |
| 40+4*part, (40+4*part)+(16*num) | Points           | Point   | Little     |
*/

func CreatePolyLineContent(bytes []byte) PolyLineContent {
	numberOfParts := fn.LittleEndianUInt32(bytes[32:36])
	numberOfPoints := fn.LittleEndianUInt32(bytes[36:40])

	// Read points
	points := make([]PointContent, numberOfPoints)
	for i := 0; i < int(numberOfPoints); i++ {
		var readStart = (40 + 4*int(numberOfParts)) + (16 * i)

		points[i] = CreatePointContent(bytes[readStart : readStart+16])
	}

	// Split points by parts
	parts := make([][]PointContent, numberOfParts)
	for i := 0; i < int(numberOfParts); i++ {
		var readStart = 40 + 4*i

		pointStart := fn.LittleEndianUInt32(bytes[readStart : readStart+4])
		pointEnd := fn.LittleEndianUInt32(bytes[readStart+4 : readStart+8])
		if int(numberOfParts-1) == i {
			pointEnd = numberOfPoints
		}
		parts[i] = points[pointStart:pointEnd]
	}

	return PolyLineContent{
		fn.LittleEndianFloat64(bytes[0:8]),
		fn.LittleEndianFloat64(bytes[8:16]),
		fn.LittleEndianFloat64(bytes[16:24]),
		fn.LittleEndianFloat64(bytes[24:32]),
		parts,
	}
}
