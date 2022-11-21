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
	   \
	    \
	     *
		  (4,4)
*/
func TestReadPolyLine1Line(t *testing.T) {
	assert := assert.New(t)

	// Box
	box := make([]byte, 8*4)
	binary.LittleEndian.PutUint64(box[0:8], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(box[8:16], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(box[16:24], math.Float64bits(4.0))
	binary.LittleEndian.PutUint64(box[24:32], math.Float64bits(4.0))
	// Number of parts
	numberOfParts := make([]byte, 4)
	binary.LittleEndian.PutUint32(numberOfParts, 1)
	// Number of points
	numberOfPoints := make([]byte, 4)
	binary.LittleEndian.PutUint32(numberOfPoints, 2)
	// Parts
	parts := make([]byte, 4)
	binary.LittleEndian.PutUint32(parts, 0)
	// Points
	points := make([]byte, 16*2)
	binary.LittleEndian.PutUint64(points[0:8], math.Float64bits(1))
	binary.LittleEndian.PutUint64(points[8:16], math.Float64bits(1))
	binary.LittleEndian.PutUint64(points[16:24], math.Float64bits(4))
	binary.LittleEndian.PutUint64(points[24:32], math.Float64bits(4))

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
	assert.Equal(polyLine.XMax, float64(4.0))
	assert.Equal(polyLine.YMax, float64(4.0))
	resultPoints := [][]record.PointContent{{
		record.CreatePointContent(points[0:16]),
		record.CreatePointContent(points[16:32]),
	}}
	assert.Equal(polyLine.Parts, resultPoints)
}

/*
	   (1,1)   (8,0)  (13,0)
		           *----*
			  *     /      \
			   \   /        *
			    \ /        (15,2)
			     *
				  (4,4)
*/
func TestReadPolyLineMultiPoint(t *testing.T) {
	assert := assert.New(t)

	// Box
	box := make([]byte, 8*4)
	binary.LittleEndian.PutUint64(box[0:8], math.Float64bits(1.0))
	binary.LittleEndian.PutUint64(box[8:16], math.Float64bits(0.0))
	binary.LittleEndian.PutUint64(box[16:24], math.Float64bits(15.0))
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
	pointXYs := [5][2]float64{{1, 1}, {4, 4}, {8, 0}, {13, 0}, {15, 2}}
	for i := 0; i < 5; i++ {
		binary.LittleEndian.PutUint64(points[4*i:4*i+8], math.Float64bits(pointXYs[i][0]))
		binary.LittleEndian.PutUint64(points[4*i+8:4*i+16], math.Float64bits(pointXYs[i][1]))
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
	assert.Equal(polyLine.YMin, float64(0.0))
	assert.Equal(polyLine.XMax, float64(15.0))
	assert.Equal(polyLine.YMax, float64(4.0))
	resultPoints := [][]record.PointContent{
		{
			record.CreatePointContent(points[0:16]),
			record.CreatePointContent(points[16:32]),
			record.CreatePointContent(points[32:48]),
			record.CreatePointContent(points[48:64]),
			record.CreatePointContent(points[64:80]),
		},
	}
	assert.Equal(polyLine.Parts, resultPoints)
}
