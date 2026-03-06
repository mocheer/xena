package gm

// Line 一条线
type Line [2][2]float64

// Mid 获取线段中点
func (m Line) Mid() [2]float64 {
	return [2]float64{
		(m[0][0] + m[1][0]) / 2,
		(m[0][1] + m[1][1]) / 2,
	}
}
