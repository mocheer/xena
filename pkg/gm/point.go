package gm

type Point [2]float64

// 转换成ToLonLat
func (m Point) ToLonLat() LonLat {
	return LonLat(m)
}

// SetXY 设置坐标
func (p *Point) SetXY(xy [2]float64) (err error) {
	p[0] = xy[0]
	p[1] = xy[1]
	return
}

// Sub
func (p *Point) Sub(xy [2]float64) Point {
	return Point{p[0] - xy[0], p[1] - xy[1]}
}
