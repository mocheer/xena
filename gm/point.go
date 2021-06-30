package gm

type Point [2]float64

// SetXY 设置坐标
func (p *Point) SetXY(xy [2]float64) (err error) {
	p[0] = xy[0]
	p[1] = xy[1]
	return
}
