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
*/
func TestReadPoint(t *testing.T) {
	assert := assert.New(t)

	bytes := make([]byte, 16)
	binary.LittleEndian.PutUint64(bytes[0:8], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(bytes[8:16], math.Float64bits(1.0))
	// Box
	point := record.CreatePointContent(bytes)

	// Validation
	assert.Equal(point.X, float64(1.0))
	assert.Equal(point.Y, float64(1.0))
}
