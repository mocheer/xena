package gm

type PointZ [3]float64

// SetXY 设置坐标
func (p *PointZ) SetXYZ(xyz [3]float64) (err error) {
	p[0] = xyz[0]
	p[1] = xyz[1]
	p[2] = xyz[2]
	return
}
