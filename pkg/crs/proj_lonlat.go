package crs

import (
	"github.com/mocheer/xena/pkg/gm"
)

type ProjLonLat struct{}

// Transform
func (m ProjLonLat) Transform() Transformation {
	return Transformation{1 / 180.0, 1, -1 / 180.0, 0.5}
}

// Proj
func (m ProjLonLat) Proj(lonlat gm.LonLat) gm.Point {
	return lonlat.ToPoint()
}

// UnProj
func (m ProjLonLat) UnProj(point gm.Point) gm.LonLat {
	return point.ToLonLat()
}
