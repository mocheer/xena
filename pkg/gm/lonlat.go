package gm

import (
	"math"
)

const RADIANS_PER_DEGREE = math.Pi / 180 // 每一个角度单位对应的弧度值
const DEGREE_PER_RADIANS = 180 / math.Pi

type LonLat [2]float64

// Lon 获取经度
func (m LonLat) Lon() float64 {
	return m[0]
}

// Lat 获取纬度
func (m LonLat) Lat() float64 {
	return m[1]
}

// ToPoint 获取点（将经纬度转换为Point对象）
func (m LonLat) ToPoint() Point {
	return Point{m[0], m[1]}
}

// ToPointZ
func (m LonLat) ToPointZ() PointZ {
	return PointZ{m[0], m[1], 0}
}

// ToCartographic
func (m LonLat) ToCartographic() Cartographic {
	return Cartographic{
		m[0] * RADIANS_PER_DEGREE,
		m[1] * RADIANS_PER_DEGREE,
	}
}

// ConvertToPoint 将经纬度转换成平面坐标
func (m LonLat) ConvertToPoint() *Point {
	return &Point{}
}
