package gm

import (
	"math"

	"github.com/mocheer/xena/pkg/alg"
)

type LonLat [2]float64

// Lon 获取经度
func (m LonLat) Lon() float64 {
	return m[0]
}

// Lat 获取纬度
func (m LonLat) Lat() float64 {
	return m[1]
}

// Point 获取点（将经纬度转换为Point对象）
func (m LonLat) Point() Point {
	return Point(m)
}

// ConvertToPoint 将经纬度转换成平面坐标
func (m LonLat) ConvertToPoint() *Point {
	return &Point{}
}

// GetTileAndOffset 获取经纬度对应的瓦片信息，这里是基于3857投影
func (m LonLat) GetTileAndOffset(z float64) (*Tile, *Point) {
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
	offsetPoint := &Point{
		float64(pixelX),
		float64(pixelY),
	}
	return &Tile{X: int(tileX), Y: int(tileY), Z: int(z)}, offsetPoint
}

// GetPointPX 获取经纬度对应的像素坐标，这里是基于3857投影，投影坐标系
func (m LonLat) GetPointPX(z float64) *Point {
	return m.GetPointPXByScale(math.Exp2(z) * 256)
}

// GetPointPXByScale 获取经纬度对应的像素坐标，这里是基于3857投影，投影坐标系
func (m LonLat) GetPointPXByScale(scale float64) *Point {
	lon := m.Lon()
	lat := m.Lat()
	tileTempX := (lon + 180.0) / 360.0 * scale
	tileTempY := math.Log(math.Tan(lat*alg.RADIANS_PER_DEGREE*0.5+alg.PI_OVER_FOUR)) / alg.PI2
	//
	pixelX := tileTempX
	pixelY := (1.0 - tileTempY) * scale
	//
	offsetPoint := &Point{
		float64(pixelX),
		float64(pixelY),
	}
	return offsetPoint
}
