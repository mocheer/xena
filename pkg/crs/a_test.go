package crs_test

import (
	"math"
	"testing"

	"github.com/mocheer/xena/pkg/crs"
	"github.com/mocheer/xena/pkg/crs/backup"
	"github.com/mocheer/xena/pkg/gm"
)

func TestCRS(t *testing.T) {
	lonlat := gm.LonLat{120, 30}

	t.Log(backup.GetTileAndOffset(lonlat, 10))
	t.Log(backup.GetPointPXByScale(lonlat, math.Exp2(10)*256))

	EPSG3857 := crs.EPSG3857
	t.Log(EPSG3857.LonlatToPoint(lonlat, 10))
	t.Log(EPSG3857.LonlatToTile(lonlat, 10))
	t.Log(EPSG3857.LonlatToTileAndOffset(lonlat, 10))
	p1 := EPSG3857.Transform(EPSG3857.Proj(lonlat), math.Exp2(10)*256)
	t.Log(p1)
}
