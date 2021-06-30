package gm

type BBox [4]float64

// MaxX
func (m BBox) MaxX() float64 {
	return m[2]
}

// MinX
func (m BBox) MinX() float64 {
	return m[0]
}

// MaxY
func (m BBox) MaxY() float64 {
	return m[3]
}

// MinY
func (m BBox) MinY() float64 {
	return m[1]
}

func (m BBox) ContainPoint(point [2]float64) bool {
	return point[0] < m.MaxX() && point[0] > m.MinX() && point[1] < m.MaxY() && point[1] > m.MinY()
}
