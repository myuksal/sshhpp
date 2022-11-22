package dbf_test

import (
	"testing"

	dbf "github.com/myuksal/sshhpp/dbf"
	"github.com/stretchr/testify/assert"
)

func TestCreateField(t *testing.T) {
	assert := assert.New(t)

	fName := make([]byte, 11)
	fName[0] = 'A'
	fName[1] = 'b'
	fName[2] = 'C'
	fType := []byte{'C'}
	fLength := []byte{uint8(123)}
	fDecimal := []byte{uint8(0)}
	b4 := make([]byte, 4)
	b2 := make([]byte, 2)
	b1 := make([]byte, 1)
	b10 := make([]byte, 10)

	var buffer []byte
	for _, bt := range [][]byte{fName, fType, b4, fLength, fDecimal, b2, b1, b10, b1} {
		buffer = append(buffer, bt...)
	}

	field := dbf.CreateDBFField(buffer)
	expect := dbf.DBFField{"AbC", 'C', 123, 0}
	assert.Equal(expect, field)
}
