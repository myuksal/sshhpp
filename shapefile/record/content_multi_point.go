package record

import (
	fn "github.com/myuksal/sshhpp/function"
)

type MultiPointContent struct {
	XMin   float64
	YMin   float64
	XMax   float64
	YMax   float64
	Points []PointContent
}

/**
## Multi point (52~ bytes)

| Byte Position  | Field            | Type    | Endianness |
|----------------|------------------|---------|------------|
| 0              | X min            | Double  | Little     |
| 8              | Y min            | Double  | Little     |
| 16             | Y max            | Double  | Little     |
| 24             | Y max            | Double  | Little     |
| 32             | Number of points | Integer | Little     |
| 36             | Points           | Point   | Little     |
*/

func CreateMultiPointContent(bytes []byte) MultiPointContent {
	numberOfPoints := fn.LittleEndianUInt32(bytes[32:36])
	points := make([]PointContent, numberOfPoints)

	for i := 0; i < int(numberOfPoints); i++ {
		var start = 36 + 16*i
		points[i] = CreatePointContent(bytes[start : start+16])
	}

	return MultiPointContent{
		fn.LittleEndianFloat64(bytes[0:8]),
		fn.LittleEndianFloat64(bytes[8:16]),
		fn.LittleEndianFloat64(bytes[16:24]),
		fn.LittleEndianFloat64(bytes[24:32]),
		points,
	}
}

func (multiPoint *MultiPointContent) Bind(bytes []byte) {
	newMultiPoint := CreateMultiPointContent(bytes)
	multiPoint.Points = newMultiPoint.Points
	multiPoint.XMax = newMultiPoint.XMax
	multiPoint.YMax = newMultiPoint.YMax
	multiPoint.XMin = newMultiPoint.XMin
	multiPoint.YMin = newMultiPoint.YMin
}
