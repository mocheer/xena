package gm

import "math"

// @see https://github.com/go-spatial/geom/blob/master/bbox.go
// @see https://github.com/spatial-go/geoos/blob/main/geojson/bbox.go
// @see https://github.com/ctessum/geom/blob/master/bounds.go
// @see https://github.com/spatial-go/geoos/blob/main/space/bound.go
// Bbox
// 左下右上，先经度后纬度
type BBox [4]float64

// MinX
func (m BBox) MinX() float64 {
	return m[0]
}

// MinY
func (m BBox) MinY() float64 {
	return m[1]
}

// MaxX
func (m BBox) MaxX() float64 {
	return m[2]
}

// MaxY
func (m BBox) MaxY() float64 {
	return m[3]
}

// Polygon
func (m BBox) Polygon() Polygon {
	p1 := [2]float64{m.MinX(), m.MinY()}
	p2 := [2]float64{m.MinX(), m.MaxY()}
	p3 := [2]float64{m.MaxX(), m.MaxY()}
	p4 := [2]float64{m.MaxX(), m.MinY()}
	return Polygon{{p1, p2, p3, p4}}
}

// Width
func (m BBox) Width() float64 {
	return m.MaxX() - m.MinX()
}

// Height
func (m BBox) Height() float64 {
	return m.MaxY() - m.MinY()
}

// ContainsPoint
func (m BBox) ContainsPoint(p [2]float64) bool {
	return p[0] < m.MaxX() && p[0] > m.MinX() && p[1] < m.MaxY() && p[1] > m.MinY()
}

// Extend
func (m *BBox) Extend(p [2]float64) {
	m[0] = math.Min(p[0], m[0])
	m[1] = math.Min(p[1], m[1])
	m[2] = math.Max(p[0], m[2])
	m[3] = math.Max(p[1], m[3])
}

// Center
func (m BBox) Center() [2]float64 {
	return [2]float64{(m.MinX() + m.MaxX()) / 2, (m.MinY() + m.MaxY()) / 2}
}

// ExtendBySizeScale
// ExtendBySizeScale(8.0/256.0)
func (m BBox) ExtendBySizeScale(scale float64) BBox {
	w := m.Width() * scale
	h := m.Height() * scale
	return BBox{
		m.MinX() - w,
		m.MinY() - h,
		m.MaxX() + w,
		m.MaxY() + h,
	}
}
