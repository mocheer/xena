package gs

import (
	"math"

	"github.com/mocheer/xena/alg"
	"github.com/mocheer/xena/gm"
)

type LonLat [2]float64

func (m LonLat) Lon() float64 {
	return m[0]
}

func (m LonLat) Lat() float64 {
	return m[1]
}

// GetTileAndOffset 获取经纬度对应的瓦片信息，这里是基于3857投影
func (m LonLat) GetTileAndOffset(z float64) (*Tile, *gm.Point) {
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
	return &Tile{X: int(tileX), Y: int(tileY), Z: int(z)}, offsetPoint
}
