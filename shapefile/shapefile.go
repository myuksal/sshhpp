package shapefile

import (
	"os"

	fn "github.com/myuksal/sshhpp/function"
	header "github.com/myuksal/sshhpp/shapefile/header"
	index "github.com/myuksal/sshhpp/shapefile/index"
	record "github.com/myuksal/sshhpp/shapefile/record"
)

type ShapeRecord interface {
	Bind([]byte)
}

func ReadShapeFile[T ShapeRecord](src string, shape T) (func() (bool, T), error) {
	fileName := fn.CreateFileNames(src)

	indexFile, err := os.Open(fileName.Shx)
	if err != nil {
		return nil, err
	}
	shapeFile, err := os.Open(fileName.Shp)
	if err != nil {
		return nil, err
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

	// Record generator
	return func() (bool, T) {
		indexRecord := getOffset()

		// End of index
		if indexRecord.RecordOffset == 0 {
			return true, shape
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

		shape.Bind(content.Content)
		return false, shape
	}, nil

}
