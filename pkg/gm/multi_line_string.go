package gm

type MultiLineString [][][2]float64

func (m MultiLineString) Polygon() Polygon {
	return Polygon(m)
}
