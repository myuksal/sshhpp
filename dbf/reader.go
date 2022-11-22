package dbf

import (
	"io"
	"os"
	"strings"

	fn "github.com/myuksal/sshhpp/function"
)

func ReadDbfFile(src string) (func() map[string]string, error) {
	fileName := fn.CreateFileNames(src)

	file, err := os.Open(fileName.Dbf)
	if err != nil {
		return nil, err
	}

	headBuffer := make([]byte, 32+32*255)
	file.Read(headBuffer)
	header := CreateDBFHeader(headBuffer)

	seekStart := int64(header.BytesOfHeader) + 1
	file.Seek(seekStart, io.SeekStart)

	return func() map[string]string {
		row := make(map[string]string)
		for _, field := range header.Fields {
			fieldBuffer := make([]byte, field.Size)
			file.Read(fieldBuffer)
			row[field.Name] = strings.Trim(string(fieldBuffer[:]), " ")
		}

		return row
	}, nil
}
