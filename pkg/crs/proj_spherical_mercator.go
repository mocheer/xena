package crs

import (
	"math"

	"github.com/mocheer/xena/pkg/alg"
	"github.com/mocheer/xena/pkg/gm"
)

// 墨卡托投影算法，对坐标进行墨卡托投影变换
type ProjSphericalMercator struct{}

func (m ProjSphericalMercator) MaxLat() float64 {
	return 85.0511287798
}

func (m ProjSphericalMercator) radius() float64 {
	return 6378137.0
}

func (m ProjSphericalMercator) Transform() Transformation {
	scale := 0.5 / (math.Pi * m.radius())
	return Transformation{scale, 0.5, -scale, 0.5}
}

func (m ProjSphericalMercator) Proj(lonlat gm.LonLat) gm.Point {
	r := m.radius()
	d := alg.RADIANS_PER_DEGREE
	maxLat := m.MaxLat()
	lat := math.Max(math.Min(maxLat, lonlat.Lat()), -maxLat)
	sin := math.Sin(lat * d)
	return gm.Point{r * lonlat.Lon() * d, r * math.Log((1+sin)/(1-sin)) / 2}
}

func (m ProjSphericalMercator) UnProj(point gm.Point) gm.LonLat {
	r := m.radius()
	d := alg.DEGREES_PER_RADIAN
	return gm.LonLat{point[0] * d / r, (2*math.Atan(math.Exp(point[1]/r)) - (math.Pi / 2)) * d}
}
