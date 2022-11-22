package index

import (
	fn "github.com/myuksal/sshhpp/function"
)

/**
## Record header (8 bytes)

| Byte Position  | Field          | Value          | Type    | Endianness |
|----------------|----------------|----------------|---------|------------|
| 0              | Record Number  | Record Number  | Integer | Big        |
| 4              | Content Length | Content Length | Integer | Big        |
*/

type ShapeIndexRecord struct {
	RecordOffset  uint32
	ContentLength uint32
}

func CreateShapeIndexRecord(bytes []byte) ShapeIndexRecord {
	return ShapeIndexRecord{
		fn.BigEndianUInt32(bytes[0:4]), // record offset
		fn.BigEndianUInt32(bytes[4:8]), // content length
	}
}

func ShapeIndexGenerator(bytes []byte) func() ShapeIndexRecord {
	cursor := 0
	return func() ShapeIndexRecord {
		cursor += 8
		return CreateShapeIndexRecord(bytes[cursor-8 : cursor])
	}
}
