package gm

import (
	"bytes"
	"math"
	"strconv"

	"github.com/mocheer/xena/pkg/alg"
)

// Tile 地图瓦片
type Tile struct {
	X, Y, Z int
}

// NewTileFromQuadKey 从quadKey实例化Tile对象
func NewTileFromQuadKey(quadKey string) Tile {
	x := 0
	y := 0
	z := len(quadKey)
	for i := z; i > 0; i-- {
		mask := 1 << (i - 1)
		switch string(quadKey[z-i]) {
		case "0":
		case "1":
			x |= mask
		case "2":
			y |= mask
		case "3":
			x |= mask
			y |= mask
		default:
			panic("无效的QuadKey")
		}
	}
	return Tile{z, y, z}
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

// ToQuadKey 必应地图瓦片id的算法
func (m Tile) ToQuadKey() string {
	x, y, z := m.X, m.Y, m.Z
	var buffer bytes.Buffer
	for i := z; i > 0; i-- {
		digit := 0
		mask := 1 << (i - 1)
		if (x & mask) != 0 {
			digit++
		}
		if (y & mask) != 0 {
			digit++
			digit++
		}
		buffer.WriteString(strconv.Itoa(digit))
	}
	return buffer.String()
}
