package sshhpp

import (
	dbfFile "github.com/myuksal/sshhpp/dbf"
	shapeFile "github.com/myuksal/sshhpp/shapefile"
	"github.com/myuksal/sshhpp/shapefile/record"
)

type Shape[T shapeFile.ShapeRecord] struct {
	IsEnd   bool
	Content T
	Columns map[string]string
}

func Point(file string) (func() *Shape[*record.PointContent], error) {
	nextShape, err := shapeFile.ReadShapeFile(file, &record.PointContent{})
	if err != nil {
		return nil, err
	}
	nextRow, err := dbfFile.ReadDbfFile(file)
	if err != nil {
		return nil, err
	}

	return func() *Shape[*record.PointContent] {
		row := nextRow()
		isEnd, content := nextShape()
		return &Shape[*record.PointContent]{isEnd, content, row}
	}, nil
}

func MultiPoint(file string) (func() *Shape[*record.MultiPointContent], error) {
	nextShape, err := shapeFile.ReadShapeFile(file, &record.MultiPointContent{})
	if err != nil {
		return nil, err
	}
	nextRow, err := dbfFile.ReadDbfFile(file)
	if err != nil {
		return nil, err
	}
	return func() *Shape[*record.MultiPointContent] {
		row := nextRow()
		isEnd, content := nextShape()
		return &Shape[*record.MultiPointContent]{isEnd, content, row}
	}, nil
}

func PolyLine(file string) (func() *Shape[*record.PolyLineContent], error) {
	nextShape, err := shapeFile.ReadShapeFile(file, &record.PolyLineContent{})
	if err != nil {
		return nil, err
	}
	nextRow, err := dbfFile.ReadDbfFile(file)
	if err != nil {
		nextRow()
		return nil, err
	}
	return func() *Shape[*record.PolyLineContent] {
		row := nextRow()
		isEnd, content := nextShape()
		return &Shape[*record.PolyLineContent]{isEnd, content, row}
	}, nil
}

func Polygon(file string) (func() *Shape[*record.PolygonContent], error) {
	nextShape, err := shapeFile.ReadShapeFile(file, &record.PolygonContent{})
	if err != nil {
		return nil, err
	}
	nextRow, err := dbfFile.ReadDbfFile(file)
	if err != nil {
		return nil, err
	}
	return func() *Shape[*record.PolygonContent] {
		row := nextRow()
		isEnd, content := nextShape()
		return &Shape[*record.PolygonContent]{isEnd, content, row}
	}, nil
}
