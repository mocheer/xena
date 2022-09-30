package crs

import "github.com/mocheer/xena/pkg/gm"

type Transformation [4]float64

func (m Transformation) Transform(point gm.Point, scale float64) gm.Point {
	x := scale * (m[0]*point[0] + m[1])
	y := scale * (m[2]*point[1] + m[3])
	return gm.Point{x, y}
}

func (m Transformation) UnTransform(point gm.Point, scale float64) gm.Point {
	x := (point[0]/scale - m[1])/m[0]
	y := (point[1]/scale - m[3])/m[2]
	return gm.Point{x, y}
}
