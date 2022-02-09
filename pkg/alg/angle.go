package alg

import (
	"math"
)

// 获取线段与x轴夹角的弧度值，顺时针，凸多边形上的点集合，可通过这种方式排序生成多边形(只适合凸多边形)
func GetLineRadinX(p1, p2 [2]float64) float64 {
	return math.Atan2((p2[1] - p1[1]), (p2[0] - p1[0]))
}

// 获取线段与x轴夹角的角度值，顺时针
func GetLineAngleX(p1, p2 [2]float64) float64 {
	return GetLineRadinX(p1, p2) * DEGREES_PER_RADIAN
}

// 获取线段与y轴夹角的弧度值，顺时针
func GetLineRadinY(p1, p2 [2]float64) float64 {
	return math.Atan2((p2[0] - p1[0]), (p2[1] - p1[1]))
}

// 获取线段与y轴夹角的角度值，顺时针
func GetLineAngleY(p1, p2 [2]float64) float64 {
	return GetLineRadinY(p1, p2) * DEGREES_PER_RADIAN
}
