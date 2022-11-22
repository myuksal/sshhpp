package dbf

import (
	fn "github.com/myuksal/sshhpp/function"
)

/*
https://en.wikipedia.org/wiki/.dbf

| Byte Position | Contents      | Meaning                                                                                |
|---------------|---------------|----------------------------------------------------------------------------------------|
| 0             | 1 byte        | Valid dBASE for DOS file; bits 0–2 indicate version number, bit 3 indicates the presence of a dBASE for DOS memo file, bits 4–6 indicate the presence of a SQL table, bit 7 indicates the presence of any memo file (either dBASE m PLUS or dBASE for DOS)|
| 1–3           | 3 bytes       | Date of last update; formatted as YYMMDD (with YY being the number of years since 1900)|
| 4–7           | 32-bit number | Number of records in the database file                                                 |
| 8–9           | 16-bit number | Number of bytes in the header                                                          |
| 10–11         | 16-bit number | Number of bytes in the record                                                          |
| 12–13         | 2 bytes       | Reserved; fill with 0                                                                  |
| 14            | 1 byte        | Flag indicating incomplete transaction                                                 |
| 15            | 1 byte        | Encryption flag                                                                        |
| 16–27         | 12 bytes      | Reserved for dBASE for DOS in a multi-user environment                                 |
| 28            | 1 byte        | Production .mdx file flag; 1 if there is a production .mdx file, 0 if not              |
| 29            | 1 byte        | Language driver ID                                                                     |
| 30–31         | 2 bytes       | Reserved; fill with 0                                                                  |
| 32–n          | 32 bytes each | array of field descriptors (see below for layout of descriptors)                       |
| n + 1         | 1 byte        | 0x0D as the field descriptor array terminator                                          |
*/

type DBFHeader struct {
	NumberOfRecords  uint32
	BytesOfHeader    uint16
	BytesOfRecord    uint16
	IsMdxFileExist   bool
	LanguageDriverId uint8
	Fields           []DBFField
}

func CreateDBFHeader(bytes []byte) DBFHeader {
	numberOfBytesInHeader := fn.LittleEndianUInt16(bytes[8:10])
	fields := make([]DBFField, int((numberOfBytesInHeader-32)/32))

	fieldsTarget := bytes[32 : 32+numberOfBytesInHeader]
	for i := 0; i < int((numberOfBytesInHeader-32)/32); i++ {
		fields[i] = CreateDBFField(fieldsTarget[i*32 : (i*32)+32])
	}
	return DBFHeader{
		fn.LittleEndianUInt32(bytes[4:8]),   // number of records
		numberOfBytesInHeader,               // number of bytes in header
		fn.LittleEndianUInt16(bytes[10:12]), // number of bytes in record
		fn.UInt8(bytes[28]) == 1,            // is production .mdx file exist
		fn.UInt8(bytes[29]),                 // language driver id
		fields,                              // fields
	}
}
