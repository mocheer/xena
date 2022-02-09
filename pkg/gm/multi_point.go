package gm

type MultiPoint [][2]float64

func (m MultiPoint) LineString() LineString {
	return LineString(m)
}
