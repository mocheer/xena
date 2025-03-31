package gm

import (
	"math"
)

// @see https://github.com/go-spatial/geom/blob/master/bbox.go
// @see https://github.com/spatial-go/geoos/blob/main/geojson/bbox.go
// @see https://github.com/ctessum/geom/blob/master/bounds.go
// @see https://github.com/spatial-go/geoos/blob/main/space/bound.go
// Bbox
// 左下右上，先经度后纬度
// 有些程序直接用 BBox([4]float64{}) 这个是危险操作，可能需要重新修正
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

func NewBBox() *BBox {
	maxNum := math.Inf(+1) //正无穷大的浮点数
	minNum := math.Inf(-1) //负无穷大的浮点数
	// 任何有限的浮点数都大于负无穷大。
	// 任何有限的浮点数都小于正无穷大。
	return &BBox{maxNum, maxNum, minNum, minNum}
}

// ToPolygon
func (m BBox) ToPolygon() Polygon {
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

// WestSouth
// LeftBottom
func (m BBox) WestSouth() LonLat {
	return LonLat{m.MinX(), m.MinY()}
}

// EastNorth
// RightTop
func (m BBox) EastNorth() LonLat {
	return LonLat{m.MaxX(), m.MaxY()}
}

// ContainsPoint
func (m BBox) ContainsPoint(p [2]float64) bool {
	return p[0] < m.MaxX() && p[0] > m.MinX() && p[1] < m.MaxY() && p[1] > m.MinY()
}

// FromSlices
func (m *BBox) FromSlices(values []float64) {
	m[0] = values[0]
	m[1] = values[1]
	m[2] = values[2]
	m[3] = values[3]
}

// Extend
func (m *BBox) Extend(p [2]float64) {
	m[0] = math.Min(p[0], m[0])
	m[1] = math.Min(p[1], m[1])
	m[2] = math.Max(p[0], m[2])
	m[3] = math.Max(p[1], m[3])
}

// ExtendPoints
func (m *BBox) ExtendPoints(coords []Point) {
	for _, p := range coords {
		if m[0] > p[0] {
			m[0] = p[0]
		}
		if m[1] > p[1] {
			m[1] = p[1]
		}
		if m[2] < p[0] {
			m[2] = p[0]
		}
		if m[3] < p[1] {
			m[3] = p[1]
		}
	}
}

// ExtendPoints
func (m *BBox) ExtendPointZs(coords []PointZ) {
	for _, p := range coords {
		if m[0] > p[0] {
			m[0] = p[0]
		}
		if m[1] > p[1] {
			m[1] = p[1]
		}
		if m[2] < p[0] {
			m[2] = p[0]
		}
		if m[3] < p[1] {
			m[3] = p[1]
		}
	}
}

// Center
func (m BBox) Center() [2]float64 {
	return [2]float64{(m.MinX() + m.MaxX()) / 2, (m.MinY() + m.MaxY()) / 2}
}

// ExtendBySizeScale
// ExtendBySizeScale(8.0/256.0)
func (m *BBox) ExtendBySizeScale(scale float64) {
	w := m.Width() * scale
	h := m.Height() * scale
	m.ExtendBySize(w, h)
}

func (m *BBox) ExtendBySize(width float64, height float64) {
	m[0] -= width
	m[1] -= height
	m[2] += width
	m[3] += height
}

func (m *BBox) Clone() *BBox {
	return &BBox{m[0], m[1], m[2], m[3]}
}
