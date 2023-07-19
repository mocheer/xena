package crs

import (
	"math"

	"github.com/mocheer/xena/pkg/gm"
)

type CRS struct {
	SRID int
	Projection
	Transformation
}

// LonlatToPoint 经纬度转投影坐标
func (m CRS) LonlatToPoint(lonlat gm.LonLat, zoom float64) gm.Point {
	return m.Transform(m.Proj(lonlat), math.Exp2(zoom))
}

// PointToLonlat 投影坐标转经纬度
func (m CRS) PointToLonlat(point gm.Point, zoom float64) gm.LonLat {
	return m.UnProj(m.UnTransform(point, math.Exp2(zoom)))
}

// LonlatToTile 经纬度转瓦片坐标
func (m CRS) LonlatToTile(lonlat gm.LonLat, zoom float64) gm.Tile {
	p1 := m.LonlatToPoint(lonlat, zoom)
	return gm.Tile{X: int(p1[0]), Y: int(p1[1]), Z: int(zoom)}
}

// LonlatToTileAndOffset 经纬度转成瓦片坐标和瓦片上的位置坐标
func (m CRS) LonlatToTileAndOffset(lonlat gm.LonLat, zoom float64) (gm.Tile, gm.Point) {
	p1 := m.LonlatToPoint(lonlat, zoom)
	tx := int(p1[0])
	ty := int(p1[1])
	x := (p1[0] - float64(tx)) * 256
	y := (p1[1] - float64(ty)) * 256
	return gm.Tile{X: tx, Y: ty, Z: int(zoom)}, gm.Point{x, y}
}

type Projection interface {
	Proj(lonlat gm.LonLat) gm.Point
	UnProj(point gm.Point) gm.LonLat
}

func FromSRID(srid int) CRS {
	switch srid {
	case 3857:
		p := ProjSphericalMercator{}
		return CRS{
			SRID:           srid,
			Projection:     p,
			Transformation: p.Transform(),
		}
	case 4326:
		p := ProjLonLat{}
		return CRS{
			SRID:           srid,
			Projection:     p,
			Transformation: p.Transform(),
		}
	}
	return CRS{}
}

var EPSG3857 = FromSRID(3857)
