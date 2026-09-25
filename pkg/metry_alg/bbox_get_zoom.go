package gm_alg

import (
	"math"

	"github.com/mocheer/pluto/pkg/fn"
	gm "github.com/mocheer/xena/pkg/metry"
)

// GetZoom 从经纬度范围和像素大小获取对应的zoom值
func GetZoom(bbox gm.BBox, size [2]float64) float64 {
	scale := size[0] / 256 / (bbox.Width() / 360.0)
	return math.Log2(scale)
}

// GetZoomRound
func GetZoomRound(bbox gm.BBox, size [2]float64) int {
	return fn.RoundInt(GetZoom(bbox, size))
}
