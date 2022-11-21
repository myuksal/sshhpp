package shapefile

import (
	"os"

	header "github.com/myuksal/sshhpp/shapefile/header"
	index "github.com/myuksal/sshhpp/shapefile/index"
	"github.com/myuksal/sshhpp/shapefile/record"
	types "github.com/myuksal/sshhpp/shapefile/type"
)

func ReadShapeFile(src string) (error, *header.ShapeHeader, func() any) {
	indexFile, err := os.Open(src + ".shx")
	if err != err {
		return err, nil, nil
	}
	shapeFile, err := os.Open(src + ".shp")
	if err != err {
		return err, nil, nil
	}

	// Read index header
	indexHeaderBuffer := make([]byte, 100)
	indexFile.Read(indexHeaderBuffer)
	indexHeader := header.CreateShapeHeader(indexHeaderBuffer[0:100])

	// Read index record
	indexRecordBuffer := make([]byte, indexHeader.FileLength/8)
	indexFile.Read(indexRecordBuffer)

	// Create index generator
	getOffset := index.ShapeIndexGenerator(indexRecordBuffer)

	// Read shape header
	shapeHeaderBuffer := make([]byte, 100)
	shapeFile.Read(shapeHeaderBuffer)
	shapeHeader := header.CreateShapeHeader(shapeHeaderBuffer)

	// Record generator
	return nil, &shapeHeader, func() any {
		indexRecord := getOffset()

		// End of index
		if indexRecord.RecordOffset == 0 {
			return nil
		}

		// Seek record offset
		shapeFile.Seek(int64(indexRecord.RecordOffset*2), 0)

		// Read shape record header
		recordHeaderBuffer := make([]byte, 8)
		shapeFile.Read(recordHeaderBuffer)
		recordHeader := record.CreateRecordHeader(recordHeaderBuffer)

		// Read shape record content
		recordContentBuffer := make([]byte, recordHeader.ContentLength*2)
		shapeFile.Read(recordContentBuffer)

		content := record.CreateRecordContent(recordContentBuffer)

		var result any
		switch content.Shape {
		case types.Point:
			result = record.CreatePointContent(*content.Content)
			break
		case types.MultiPoint:
			result = record.CreateMultiPointContent(*content.Content)
			break
		case types.PolyLine:
			result = record.CreatePolyLineContent(*content.Content)
			break
		case types.Polygon:
			result = record.CreatePolygonContent(*content.Content)
			break
		}

		return result
	}

}
