package record

import (
	fn "github.com/myuksal/sshhpp/shapefile/function"
	types "github.com/myuksal/sshhpp/shapefile/type"
)

/**
## Record header (8 bytes)

| Byte Position  | Field          | Value          | Type    | Endianness |
|----------------|----------------|----------------|---------|------------|
| 0              | Record Number  | Record Number  | Integer | Big        |
| 4              | Content Length | Content Length | Integer | Big        |
*/

type ShapeRecordHeader struct {
	RecordNumber  uint32
	ContentLength uint32
}

func CreateRecordHeader(bytes []byte) ShapeRecordHeader {
	return ShapeRecordHeader{
		fn.BigEndianUInt32(bytes[0:4]), // record number
		fn.BigEndianUInt32(bytes[4:8]), // content length
	}
}

/**
## Record content (4~ bytes)

| Byte Position  | Field      | Value      | Type    | Endianness |
|----------------|------------|------------|---------|------------|
| 0              | Shape Type | Shape Type | Integer | Little     |
| 4              | Content    | Content    | Integer |            |
*/

type ShapeRecordContent struct {
	Shape   types.ShapeType
	Content *[]byte
}

func CreateRecordContent(bytes []byte) ShapeRecordContent {
	shapeType := types.ShapeType(fn.LittleEndianUInt32(bytes[0:4]))
	content := bytes[4:]

	return ShapeRecordContent{shapeType, &content}
}
