package tile_arcgis

import (
	"github.com/mocheer/xena/pkg/crs"
	"github.com/mocheer/xena/pkg/gm"
)

type ArcgisTileCDI struct {
	XMin float64
	YMin float64
	XMax float64
	YMax float64
	WKID int // 切片服务wkid
}

func (m *ArcgisTileCDI) GetCenter() [2]float64 {
	return m.GetLonLatBbox().Center()
}

func (m *ArcgisTileCDI) GetLonLatBbox() gm.BBox {
	switch m.WKID {
	case 3857, 102100:
		min := crs.EPSG3857.UnProj(gm.Point{m.XMin, m.YMin})
		max := crs.EPSG3857.UnProj(gm.Point{m.XMax, m.YMax})
		return gm.BBox{
			min.Lon(),
			min.Lat(),
			max.Lon(),
			max.Lat(),
		}
	default:
		return m.GetBbox()
	}
}

func (m *ArcgisTileCDI) GetBbox() gm.BBox {
	return gm.BBox{
		m.XMin,
		m.YMin,
		m.XMax,
		m.YMax,
	}
}
