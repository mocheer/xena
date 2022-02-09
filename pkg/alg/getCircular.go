package alg

import "math"

// GetCircular 已知圆上任意两点求圆心和半径_已知圆上三点坐标求圆心和半径
func GetCircular(p1, p2, p3 [2]float64) (float64, [2]float64) {
	a := 2 * (p2[0] - p1[0])
	b := 2 * (p2[1] - p1[1])
	c := p2[0]*p2[0] + p2[1]*p2[1] - p1[0]*p1[0] - p1[1]*p1[1]
	d := 2 * (p3[0] - p2[0])
	e := 2 * (p3[1] - p2[1])
	f := p3[0]*p3[0] + p3[1]*p3[1] - p2[0]*p2[0] - p2[1]*p2[1]
	x := (b*f - e*c) / (b*d - e*a)
	y := (d*c - a*f) / (b*d - e*a)
	r := math.Sqrt(((x-p1[0])*(x-p1[0]) + (y-p1[1])*(y-p1[1])))
	return r, [2]float64{x, y}
}
