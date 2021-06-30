package gm_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mocheer/xena/gm"
	"github.com/stretchr/testify/assert"
)

func TestPointInPolygon(t *testing.T) {
	rectangle := gm.Polygon{
		{
			{1.0, 1.0},
			{1.0, 2.0},
			{2.0, 2.0},
			{2.0, 1.0},
		},
	}

	pt1 := gm.Point{1.1, 1.1}     // Should be true
	pt2 := gm.Point{1.2, 1.2}     // Should be true
	pt3 := gm.Point{1.3, 1.3}     // Should be true
	pt4 := gm.Point{1.4, 1.4}     // Should be true
	pt5 := gm.Point{1.5, 1.5}     // Should be true
	pt6 := gm.Point{1.6, 1.6}     // Should be true
	pt7 := gm.Point{1.7, 1.7}     // Should be true
	pt8 := gm.Point{1.8, 1.8}     // Should be true
	pt9 := gm.Point{-4.9, 1.2}    // Should be false
	pt10 := gm.Point{10.0, 10.0}  // Should be false
	pt11 := gm.Point{-5.0, -6.0}  // Should be false
	pt12 := gm.Point{-13.0, 1.0}  // Should be false
	pt13 := gm.Point{4.9, -1.2}   // Should be false
	pt14 := gm.Point{10.0, -10.0} // Should be false
	pt15 := gm.Point{5.0, 6.0}    // Should be false
	pt16 := gm.Point{-13.0, 1.0}  // Should be false

	assert.Equal(t, rectangle.ContainPoint(pt1), true)
	assert.Equal(t, rectangle.ContainPoint(pt2), true)
	assert.Equal(t, rectangle.ContainPoint(pt3), true)
	assert.Equal(t, rectangle.ContainPoint(pt4), true)
	assert.Equal(t, rectangle.ContainPoint(pt5), true)
	assert.Equal(t, rectangle.ContainPoint(pt6), true)
	assert.Equal(t, rectangle.ContainPoint(pt7), true)
	assert.Equal(t, rectangle.ContainPoint(pt8), true)
	//
	assert.Equal(t, rectangle.ContainPoint(pt9), false)
	assert.Equal(t, rectangle.ContainPoint(pt10), false)
	assert.Equal(t, rectangle.ContainPoint(pt11), false)
	assert.Equal(t, rectangle.ContainPoint(pt12), false)
	assert.Equal(t, rectangle.ContainPoint(pt13), false)
	assert.Equal(t, rectangle.ContainPoint(pt14), false)
	assert.Equal(t, rectangle.ContainPoint(pt15), false)
	assert.Equal(t, rectangle.ContainPoint(pt16), false)

	t.Log("Finished")
}

// TestBezierCurve 测试贝塞尔曲线
func TestBezierCurve(t *testing.T) {
	n := 5
	var data []gm.Point
	for index := 0; index < n; index++ {
		data = append(data, gm.Point{rand.Float64() * 800, rand.Float64() * 500})
	}
	bezierCurve := gm.NewBezierCurve(data)
	points := bezierCurve.GetPoints(0.01)
	for index, point := range points {
		fmt.Println(index, point, point)
	}
}
