package shapefile

// Shape Type
type ShapeType uint32

const (
	Null_Shape ShapeType = iota + 0
	Point
	_
	PolyLine
	_
	Polygon
	_
	_
	MultiPoint
	_
	_
	PointZ
	_
	PolyLineZ
	_
	PolygonZ
	_
	_
	MultiPointZ
	_
	_
	PointM
	_
	PolyLineM
	_
	PolygonM
	_
	_
	MultiPointM
	_
	_
	MultiPatch
)

func (t ShapeType) String() string {
	switch t {
	case Null_Shape:
		return "Null Shape"
	case Point:
		return "Point"
	case PolyLine:
		return "PolyLine"
	case Polygon:
		return "Polygon"
	case MultiPoint:
		return "MultiPoint"
	case PointZ:
		return "PointZ"
	case PolyLineZ:
		return "PolyLineZ"
	case PolygonZ:
		return "PolygonZ"
	case MultiPointZ:
		return "MultiPointZ"
	case PointM:
		return "PointM"
	case PolyLineM:
		return "PolyLineM"
	case PolygonM:
		return "PolygonM"
	case MultiPointM:
		return "MultiPointM"
	case MultiPatch:
		return "MultiPatch"
	default:
		return ""
	}
}
