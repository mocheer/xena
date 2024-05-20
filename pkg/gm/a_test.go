package gm_test

import (
	"math/rand/v2"
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/stretchr/testify/assert"
)

func TestPolygonConctains(t *testing.T) {
	rectangle := gm.Polygon{
		{
			{1.0, 1.0},
			{1.0, 2.0},
			{2.0, 2.0},
			{2.0, 1.0},
		},
	}

	pt1 := gm.Point{1.1, 1.1}
	pt2 := gm.Point{1.2, 1.2}
	pt3 := gm.Point{1.3, 1.3}
	pt4 := gm.Point{1.4, 1.4}
	pt5 := gm.Point{1.5, 1.5}
	pt6 := gm.Point{1.6, 1.6}
	pt7 := gm.Point{1.7, 1.7}
	pt8 := gm.Point{1.8, 1.8}
	pt9 := gm.Point{-4.9, 1.2}
	pt10 := gm.Point{10.0, 10.0}
	pt11 := gm.Point{-5.0, -6.0}
	pt12 := gm.Point{-13.0, 1.0}
	pt13 := gm.Point{4.9, -1.2}
	pt14 := gm.Point{10.0, -10.0}
	pt15 := gm.Point{5.0, 6.0}
	pt16 := gm.Point{-13.0, 1.0}

	assert.Equal(t, rectangle.ContainsPoint(pt1), true)
	assert.Equal(t, rectangle.ContainsPoint(pt2), true)
	assert.Equal(t, rectangle.ContainsPoint(pt3), true)
	assert.Equal(t, rectangle.ContainsPoint(pt4), true)
	assert.Equal(t, rectangle.ContainsPoint(pt5), true)
	assert.Equal(t, rectangle.ContainsPoint(pt6), true)
	assert.Equal(t, rectangle.ContainsPoint(pt7), true)
	assert.Equal(t, rectangle.ContainsPoint(pt8), true)
	//
	assert.Equal(t, rectangle.ContainsPoint(pt9), false)
	assert.Equal(t, rectangle.ContainsPoint(pt10), false)
	assert.Equal(t, rectangle.ContainsPoint(pt11), false)
	assert.Equal(t, rectangle.ContainsPoint(pt12), false)
	assert.Equal(t, rectangle.ContainsPoint(pt13), false)
	assert.Equal(t, rectangle.ContainsPoint(pt14), false)
	assert.Equal(t, rectangle.ContainsPoint(pt15), false)
	assert.Equal(t, rectangle.ContainsPoint(pt16), false)
}

func TestPolygonGrids(t *testing.T) {
	poly := gm.NewPolygonFromJSON("./testdata/polygon.json")
	data := poly.Grids(0.00001)
	ds_json.Save("./testdata/polygon-grids.json", data)
	t.Log(len(data))
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
		t.Log(index, point, point)
	}
}

// func TestPolygonGrids2(t *testing.T) {
// 	fileName := "./testdata/ASTGTMV003_N43E093/ASTGTMV003_N43E093_dem.tif"
// 	tfs := gtif.Read(fileName)

// 	size := 0.0001
// 	// 源头
// 	data := []float64{93.87096088876093, 43.07181417965543}
// 	water := 1000 //假设水位10m，网格点为0.1m

// 	for _, grid := range data {
// 		lon, lat := grid[0], grid[1]
// 		height := tfs.GetAltByLonLat(lon, lat)
// 		//
// 		grid1 := []float64{lon + size, lat}
// 		grid2 := []float64{lon, lat + size}
// 		grid3 := []float64{lon + size, lat + size}
// 		grid4 := []float64{lon - size, lat}
// 		grid5 := []float64{lon, lat - size}
// 		grid6 := []float64{lon - size, lat - size}
// 		grid7 := []float64{lon + size, lat - size}
// 		grid8 := []float64{lon - size, lat + size}

// 		if tfs.GetAltByLonLat(grid1) < height {

// 		}

// 	}

// 	t.Log(len(data))
// }
