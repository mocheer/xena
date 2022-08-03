package gm_alg

import (
	"math"

	"github.com/mocheer/xena/pkg/gm"
)

// GetZoom 从经纬度范围和像素大小获取对应的zoom值
func GetZoom(bbox gm.BBox, size [2]float64) float64 {
	scale := size[0] / 256 / (bbox.Width() / 360.0)
	return math.Log2(scale)
}
