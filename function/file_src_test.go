package function_test

import (
	"testing"

	fn "github.com/myuksal/sshhpp/function"
	"github.com/stretchr/testify/assert"
)

func TestFileName(t *testing.T) {
	assert := assert.New(t)

	testFile := "test_file_name"
	fileName := fn.CreateFileNames(testFile)

	assert.Equal(testFile+".shp", fileName.Shp)
	assert.Equal(testFile+".shx", fileName.Shx)
	assert.Equal(testFile+".dbf", fileName.Dbf)
}

func TestFileNameShpSuffix(t *testing.T) {
	assert := assert.New(t)

	testFile := "test_file_name"
	fileName := fn.CreateFileNames(testFile + ".shp")

	assert.Equal(testFile+".shp", fileName.Shp)
	assert.Equal(testFile+".shx", fileName.Shx)
	assert.Equal(testFile+".dbf", fileName.Dbf)
}

func TestFileNameShxSuffix(t *testing.T) {
	assert := assert.New(t)

	testFile := "test_file_name"
	fileName := fn.CreateFileNames(testFile + ".shx")

	assert.Equal(testFile+".shp", fileName.Shp)
	assert.Equal(testFile+".shx", fileName.Shx)
	assert.Equal(testFile+".dbf", fileName.Dbf)
}

func TestFileNameDbfSuffix(t *testing.T) {
	assert := assert.New(t)

	testFile := "test_file_name"
	fileName := fn.CreateFileNames(testFile + ".dbf")

	assert.Equal(testFile+".shp", fileName.Shp)
	assert.Equal(testFile+".shx", fileName.Shx)
	assert.Equal(testFile+".dbf", fileName.Dbf)
}
