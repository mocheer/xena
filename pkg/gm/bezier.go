package gm

import (
	"github.com/mocheer/xena/pkg/alg"
)

//BezierCurve interface
type BezierCurve interface {
	GetPoints(step float64) []Point
	GetPoint(t float64) Point
}

// bezierCurve implements the BezierCurve interface.
type bezierCurve struct {
	data []Point
}

//NewBezierCurve creates a new bezier curve.
func NewBezierCurve(data []Point) BezierCurve {
	return &bezierCurve{data}
}

//GetPoints step [0,1]
func (m *bezierCurve) GetPoints(step float64) []Point {
	var points []Point
	for t := 0.0; t <= 1; t += step {
		points = append(points, m.GetPoint(t))
	}
	return points
}

//GetPoint t [0,1]
func (m *bezierCurve) GetPoint(t float64) Point {
	data := m.data
	var x float64
	var y float64
	n := len(data)
	for i, p := range data {
		b := alg.BezierFormula(n-1, i, t)
		x += p[0] * b
		y += p[1] * b
	}
	return Point{x, y}
}
