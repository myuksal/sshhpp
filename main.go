package main

import (
	"fmt"

	shapeFile "github.com/myuksal/sshhpp/shapefile"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	e, header, next := shapeFile.ReadShapeFile("./init_data/building/건물도형_전체분_세종특별자치시/Total.JUSURBALL.20220701.TL_SGCO_RNADR_ALL.36110")

	if e != nil {
		panic(e)
	}

	fmt.Println(header)

	for {
		data := next()
		if data == nil {
			break
		}
	}

}
