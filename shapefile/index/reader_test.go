package index_test

import (
	"encoding/binary"
	"testing"

	"github.com/myuksal/sshhpp/shapefile/index"
	"github.com/stretchr/testify/assert"
)

func TestCreateShapeIndexRecord(t *testing.T) {
	assert := assert.New(t)

	bytes := make([]byte, 8)
	binary.BigEndian.PutUint32(bytes[0:4], uint32(0))
	binary.BigEndian.PutUint32(bytes[4:8], uint32(10))

	index := index.CreateShapeIndexRecord(bytes)

	assert.Equal(index.RecordOffset, uint32(0))
	assert.Equal(index.ContentLength, uint32(10))
}

func TestShapeIndexGenerator(t *testing.T) {
	assert := assert.New(t)

	rowCount := 1234
	bytes := make([]byte, 8*rowCount)

	for i := 0; i < rowCount; i++ {
		binary.BigEndian.PutUint32(bytes[8*i:8*i+4], uint32(i))
		binary.BigEndian.PutUint32(bytes[8*i+4:8*i+8], uint32(10))
	}

	gen := index.ShapeIndexGenerator(bytes)

	for i := 0; i < rowCount; i++ {
		index := gen()
		assert.Equal(index.RecordOffset, uint32(i))
		assert.Equal(index.ContentLength, uint32(10))
	}
}
