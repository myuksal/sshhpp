package function

import "strings"

type FileName struct {
	Shp string
	Shx string
	Dbf string
}

func CreateFileNames(src string) FileName {
	baseSrc := src

	if strings.HasSuffix(src, ".shp") {
		baseSrc = strings.TrimSuffix(src, ".shp")
	}
	if strings.HasSuffix(src, ".shx") {
		baseSrc = strings.TrimSuffix(src, ".shx")
	}
	if strings.HasSuffix(src, ".dbf") {
		baseSrc = strings.TrimSuffix(src, ".dbf")
	}

	return FileName{baseSrc + ".shp", baseSrc + ".shx", baseSrc + ".dbf"}
}
