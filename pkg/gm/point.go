package gm

import "math"

type Point [2]float64

// 转换成ToLonLat
func (m Point) ToLonLat() LonLat {
	return LonLat(m)
}

// SetXY 设置坐标
func (m *Point) Copy(xy Point) (err error) {
	m[0] = xy[0]
	m[1] = xy[1]
	return
}

// Sub
func (m *Point) Sub(xy [2]float64) Point {
	return Point{m[0] - xy[0], m[1] - xy[1]}
}

// 计算两个坐标点之间的距离
func (m Point) Distance(p Point) float64 {
	return math.Sqrt(math.Pow(p[0]-m[0], 2) + math.Pow(p[1]-m[1], 2))
}
