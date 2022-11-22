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
		  ┌────────┐
		  │        │
		  │        │
		  └────────┘
				     (10,4)
*/
func TestReadPolygon1(t *testing.T) {
	assert := assert.New(t)

	// Box
	box := make([]byte, 8*4)
	binary.LittleEndian.PutUint64(box[0:8], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(box[8:16], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(box[16:24], math.Float64bits(10.0))
	binary.LittleEndian.PutUint64(box[24:32], math.Float64bits(4.0))
	// Number of parts
	numberOfParts := make([]byte, 4)
	binary.LittleEndian.PutUint32(numberOfParts, 1)
	// Number of points
	numberOfPoints := make([]byte, 4)
	binary.LittleEndian.PutUint32(numberOfPoints, 5)
	// Parts
	parts := make([]byte, 4)
	binary.LittleEndian.PutUint32(parts, 0)
	// Points
	points := make([]byte, 16*5)
	mockPoints := [][]float64{{1, 1}, {10, 1}, {10, 4}, {1, 4}, {1, 1}}
	for i, p := range mockPoints {
		startIndex := 16 * i
		binary.LittleEndian.PutUint64(points[startIndex:startIndex+8], math.Float64bits(p[0]))
		binary.LittleEndian.PutUint64(points[startIndex+8:startIndex+16], math.Float64bits(p[1]))
	}

	// Make bytes parameter
	var input []byte
	input = append(input, box...)
	input = append(input, numberOfParts...)
	input = append(input, numberOfPoints...)
	input = append(input, parts...)
	input = append(input, points...)

	// Run
	polyLine := record.CreatePolyLineContent(input)

	// Validation
	assert.Equal(polyLine.XMin, float64(1.0))
	assert.Equal(polyLine.YMin, float64(1.0))
	assert.Equal(polyLine.XMax, float64(10.0))
	assert.Equal(polyLine.YMax, float64(4.0))
	for i, p := range mockPoints {
		expectPoint := record.PointContent{X: p[0], Y: p[1]}
		assert.Equal(polyLine.Parts[0][i], expectPoint)
	}
}

/*
	  (1,1)
			┌──────────────┐
			│              │
			│ ┌─────┐      │
			│ │     │      │
			│ │     │      │
			│ └─────┘      │
			│              │
			│              │
			└──────────────┘
				           (17,9)
*/
func TestReadPolygon2(t *testing.T) {
	assert := assert.New(t)

	// Box
	box := make([]byte, 8*4)
	binary.LittleEndian.PutUint64(box[0:8], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(box[8:16], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(box[16:24], math.Float64bits(17.0))
	binary.LittleEndian.PutUint64(box[24:32], math.Float64bits(9.0))
	// Number of parts
	numberOfParts := make([]byte, 4)
	binary.LittleEndian.PutUint32(numberOfParts, 2)
	// Number of points
	numberOfPoints := make([]byte, 4)
	binary.LittleEndian.PutUint32(numberOfPoints, 10)
	// Parts
	parts := make([]byte, 8)
	binary.LittleEndian.PutUint32(parts[0:4], 0)
	binary.LittleEndian.PutUint32(parts[4:8], 5)
	// Points
	points := make([]byte, 16*10)
	mockPoints := [][]float64{
		{1, 1}, {17, 1}, {17, 9}, {1, 9}, {1, 1},
		{3, 3}, {9, 3}, {9, 6}, {3, 6}, {3, 3},
	}
	for i, p := range mockPoints {
		startIndex := 16 * i
		binary.LittleEndian.PutUint64(points[startIndex:startIndex+8], math.Float64bits(p[0]))
		binary.LittleEndian.PutUint64(points[startIndex+8:startIndex+16], math.Float64bits(p[1]))
	}

	// Make bytes parameter
	var input []byte
	input = append(input, box...)
	input = append(input, numberOfParts...)
	input = append(input, numberOfPoints...)
	input = append(input, parts...)
	input = append(input, points...)

	// Run
	polyLine := record.CreatePolyLineContent(input)

	// Validation
	assert.Equal(polyLine.XMin, float64(1.0))
	assert.Equal(polyLine.YMin, float64(1.0))
	assert.Equal(polyLine.XMax, float64(17.0))
	assert.Equal(polyLine.YMax, float64(9.0))
	for i, p := range mockPoints[0:5] {
		expectPoint := record.PointContent{X: p[0], Y: p[1]}
		assert.Equal(polyLine.Parts[0][i], expectPoint)
	}
	for i, p := range mockPoints[5:10] {
		expectPoint := record.PointContent{X: p[0], Y: p[1]}
		assert.Equal(polyLine.Parts[1][i], expectPoint)
	}
}
