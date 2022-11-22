package sshhpp_test

import (
	"testing"

	"github.com/myuksal/sshhpp"
)

func TestReadPoint(t *testing.T) {
	// Download from https://gis-pdx.opendata.arcgis.com/datasets/1e41db35866d41aa961cb96bd8f1ce7b/explore?location=45.472143%2C-122.663150%2C14.88
	next, err := sshhpp.Point("./test_file/South_Portland_Retired_Address_Points.shp")
	if err != nil {
		panic(err)
	}

	for {
		shape := next()
		if err != nil {
			panic(err)
		}
		if shape.IsEnd {
			break
		}
	}
}
