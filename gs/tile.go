package gs

import (
	"math"

	"github.com/mocheer/xena/alg"
	"github.com/mocheer/xena/gm"
)

// Tile 地图瓦片
type Tile struct {
	X, Y, Z int
}

// GetTopLeft 获取瓦片左上角经纬度
func (m Tile) GetTopLeft() *gm.Point {
	x, y, z := m.X, m.Y, m.Z
	n := math.Pi - 2*math.Pi*float64(y)/math.Pow(float64(2), float64(z))
	return &gm.Point{float64(x)/math.Exp2(float64(z))*360 - 180, (alg.DEGREES_PER_RADIAN * math.Atan(0.5*(math.Exp(n)-math.Exp(-1.0*n))))}
}
