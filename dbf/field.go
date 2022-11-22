package dbf

import "github.com/myuksal/sshhpp/function"

/*
https://en.wikipedia.org/wiki/.dbf

| Byte Position | Contents | Meaning                                                                                     |
|---------------|----------|---------------------------------------------------------------------------------------------|
| 0–10          | 11 bytes | Field name in ASCII (zero-filled)                                                           |
| 11            | 1 byte   | Field type. Allowed values: C, D, F, L, M, or N (see next table for meanings)               |
| 12–15         | 4 bytes  | Reserved                                                                                    |
| 16            | 1 byte   | Field length in binary (maximum 254 (0xFE))                                                 |
| 17            | 1 byte   | Field decimal count in binary                                                               |
| 18–19         | 2 bytes  | Work area ID                                                                                |
| 20            | 1 byte   | Example                                                                                     |
| 21–30         | 10 bytes | Reserved                                                                                    |
| 31            | 1 byte   | Production MDX field flag; 1 if field has an index tag in the production MDX file, 0 if not |
*/

type DBFField struct {
	Name            string
	Type            byte
	Size            uint8
	NumberOfDecimal uint8
}

func CreateDBFField(bytes []byte) DBFField {
	return DBFField{
		function.Ascii(bytes[:11]),
		bytes[11],
		function.UInt8(bytes[16]),
		function.UInt8(bytes[17]),
	}
}
