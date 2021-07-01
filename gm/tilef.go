package gm

import (
	"math"

	"github.com/mocheer/xena/alg"
)

// TileF 地图瓦片
type TileF struct {
	X, Y, Z float64
}

/**
 * 返回一个包含当前坐标的新的 coordinate 对象。即：向下取整。
 */
func (m *TileF) Tile() *Tile {
	c := m.Container()
	return &Tile{int(c.X), int(m.Y), int(m.Z)}
}

/**
 * 返回一个包含当前坐标的新的 coordinate 对象。即：向下取整。
 */
func (m *TileF) Container() *TileF {
	return &TileF{X: math.Floor(m.X), Y: math.Floor(m.Y), Z: m.Z}
}

/**
 * 将当前坐标缩放至 destination 级别，返回缩放后的对象副本。该方法不修改原始对象。
 * @param	destination	缩放后的缩放级别
 * @return	缩放后对应的坐标（行、列、缩放级别）
 */
func (m *TileF) ZoomTo(destination float64) *TileF {
	return &TileF{m.Y * math.Pow(2, destination-m.Z), m.X * math.Pow(2, destination-m.Z), destination}
}

/**
 * 对当前坐标缩放 distance 级，返回缩放后对象副本。该方法不修改原始对象。
 * @param	distance	要进行缩放的等级(正数为放大，负数为缩小)
 * @return	缩放后对应的坐标（行、列、缩放级别）
 */
func (m *TileF) ZoomBy(distance float64) *TileF {
	return &TileF{m.Y * math.Pow(2, distance), m.X * math.Pow(2, distance), m.Z + distance}
}

/**
 * IsRowEdge 当前坐标是否恰为某行（即不含小数）
 * @return	如果恰为某行，则返回true，否则返回false
 */
func (m *TileF) IsRowEdge() bool {
	return math.Floor(m.Y) == m.Y
}

/**
 * IsColumnEdge 当前坐标是否恰为某列（即不含小数）
 * @return	如果恰为某列，则返回true，否则返回false
 */
func (m *TileF) IsColumnEdge() bool {
	return math.Floor(m.X) == m.X
}

/**
 * 当前坐标是否恰为某行列（即不含小数）
 * @return	如果恰为某行某列，则返回true，否则返回false
 */
func (m *TileF) IsEdge() bool {
	return m.IsRowEdge() && m.IsColumnEdge()
}

// GetLonLat 获取瓦片所在的经纬度
func (m TileF) GetLonLat() *LonLat {
	x, y, z := m.X, m.Y, m.Z
	n := math.Pi - 2*math.Pi*y/math.Pow(2, z)
	return &LonLat{x/math.Exp2(z)*360 - 180, (alg.DEGREES_PER_RADIAN * math.Atan(0.5*(math.Exp(n)-math.Exp(-1.0*n))))}
}

// GetLonLat 获取瓦片左上角经纬度
func (m TileF) GetTopLeftLonLat() *LonLat {
	return m.Container().GetLonLat()
}
