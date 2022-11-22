package record_test

import (
	"encoding/binary"
	"math"
	"testing"

	record "github.com/myuksal/sshhpp/shapefile/record"
	"github.com/stretchr/testify/assert"
)

/*
		(1,1)
		  *


	          (8,5)
	          	*
*/
func TestReadMultiPoint(t *testing.T) {
	assert := assert.New(t)

	bytes := make([]byte, 36+16*4)
	// Put bounds
	binary.LittleEndian.PutUint64(bytes[0:8], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(bytes[8:16], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(bytes[16:24], math.Float64bits(8.0))
	binary.LittleEndian.PutUint64(bytes[24:32], math.Float64bits(5.0))
	// Put number of points
	binary.LittleEndian.PutUint32(bytes[32:36], 2)
	// Put points
	binary.LittleEndian.PutUint64(bytes[36+0:36+8], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(bytes[36+8:36+16], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(bytes[36+16:36+24], math.Float64bits(8.0))
	binary.LittleEndian.PutUint64(bytes[36+24:36+32], math.Float64bits(5.0))

	// Box
	multiPoint := record.CreateMultiPointContent(bytes)

	// Validation
	assert.Equal(multiPoint.Points[0].X, float64(1.0))
	assert.Equal(multiPoint.Points[0].Y, float64(1.0))
	assert.Equal(multiPoint.Points[1].X, float64(8.0))
	assert.Equal(multiPoint.Points[1].Y, float64(5.0))
}
