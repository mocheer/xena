package backup

import (
	"math"

	"github.com/mocheer/xena/pkg/alg"
	"github.com/mocheer/xena/pkg/gm"
)

// 弃用
// GetTileAndOffset 获取经纬度对应的瓦片信息，这里是基于3857投影
func GetTileAndOffset(m gm.LonLat, z float64) (*gm.Tile, *gm.Point) {
	lon := m.Lon()
	lat := m.Lat()
	scaleZ := math.Exp2(z)
	tileTempX := (lon + 180.0) / 360.0 * scaleZ
	tileTempY := math.Log(math.Tan(lat*alg.RADIANS_PER_DEGREE*0.5+alg.PI_OVER_FOUR)) / alg.PI2
	// 瓦片坐标取整，返回小于或者等于的最大整数
	tileX := math.Floor(tileTempX)
	tileY := math.Floor((0.5 - tileTempY) * scaleZ)
	//
	pixelX := int(tileTempX*256.0) % 256
	pixelY := int((1.0-tileTempY)*scaleZ*256.0) % 256
	//
	offsetPoint := &gm.Point{
		float64(pixelX),
		float64(pixelY),
	}
	return &gm.Tile{X: int(tileX), Y: int(tileY), Z: int(z)}, offsetPoint
}

// 弃用
// GetPointPX 获取经纬度对应的像素坐标，这里是基于3857投影，投影坐标系
func GetPointPX(m gm.LonLat, z float64) *gm.Point {
	return GetPointPXByScale(m, math.Exp2(z)*256)
}

// 弃用
// GetPointPXByScale 获取经纬度对应的像素坐标，这里是基于3857投影，投影坐标系
func GetPointPXByScale(m gm.LonLat, scale float64) *gm.Point {
	lon := m.Lon()
	lat := m.Lat()
	tileTempX := (lon + 180.0) / 360.0 * scale
	tileTempY := math.Log(math.Tan(lat*alg.RADIANS_PER_DEGREE*0.5+alg.PI_OVER_FOUR)) / alg.PI2
	//
	pixelX := tileTempX
	pixelY := (1.0 - tileTempY) * scale
	//
	offsetPoint := &gm.Point{
		float64(pixelX),
		float64(pixelY),
	}
	return offsetPoint
}
