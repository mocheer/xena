package gm

import (
	"math"

	"github.com/mocheer/xena/pkg/alg"
)

// Tile 地图瓦片
type Tile struct {
	X, Y, Z int
}

/**
 * 返回当前坐标上移 distance 行对应的坐标，即：Row - distance
 * @param	distance	往上的行数
 * @return			对应坐标
 */
func (m Tile) Up(distance int) *Tile {
	return &Tile{m.Y - distance, m.X, m.Z}
}

/**
 * 返回当前坐标右移 distance 列对应的坐标，即：Column + distance
 * @param	distance	往右的列数
 * @return			对应坐标
 */
func (m Tile) Right(distance int) *Tile { // = 1
	return &Tile{m.Y, m.X + distance, m.Z}
}

/**
 * 返回当前坐标下移 distance 行对应的坐标，即：Row + distance
 * @param	distance	下移的行数
 * @return			对应坐标
 */
func (m Tile) Down(distance int) *Tile { //
	return &Tile{m.Y + distance, m.X, m.Z}
}

/**
 * 返回当前坐标左移 distance 列对应的坐标，即：Column - distance
 * @param	distance	往左的列数
 * @return			对应坐标
 */
func (m Tile) Left(distance int) *Tile {
	return &Tile{m.Y, m.X - distance, m.Z}
}

// GetLonLat 获取瓦片所在的经纬度
func (m Tile) GetLonLat() *LonLat {
	x, y, z := float64(m.X), float64(m.Y), float64(m.Z)
	n := math.Pi - 2*math.Pi*y/math.Pow(2, z)
	return &LonLat{x/math.Exp2(z)*360 - 180, (alg.DEGREES_PER_RADIAN * math.Atan(0.5*(math.Exp(n)-math.Exp(-1.0*n))))}
}
