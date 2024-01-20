package tile

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/console"
	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/ts/awt"
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/crs"
	"github.com/mocheer/xena/pkg/gm"
)

// LoadAndSave 下载瓦片服务中的所有瓦片数据
func (m LoadConfig) LoadAndSave() error {
	return m.Load(m.loadingAndSave)
}

func (m LoadConfig) Load(callback func(*gm.Tile) error) error {
	minZoom := m.MinZoom
	maxZoom := m.MaxZoom
	for i := minZoom; i < maxZoom; i++ {
		m.loadingByZoom(i, callback)
	}
	return nil
}

// loadingByZoom 下载tile
// loadingByZoom(z,m.loadingAndSave)
func (m *LoadConfig) loadingByZoom(z int, callback func(*gm.Tile) error) error {
	count := m.FiberCount
	if count == 0 {
		count = 10
	}
	at := awt.New(count, 100)
	srid := 3857
	startX, startY, endX, endY := GetChinaTileLimit(srid, z)
	//
	for x := startX; x < endX; x++ {
		for y := startY; y < endY; y++ {
			tile := &gm.Tile{X: x, Y: y, Z: z}
			at.Add(func(args ...any) {
				tile := args[0].(*gm.Tile)
				callback(tile)
			}, tile)
		}
	}
	at.Wait()
	return nil
}

// loadingAndSave
func (m *LoadConfig) loadingAndSave(tile *gm.Tile) error {
	rootUrl := m.URL
	dirName := m.DirName

	x, y, z := tile.X, tile.Y, tile.Z
	savePath := m.SavePath

	if savePath == "" {
		savePath = "{z}/{x}/{y}.png"
	}
	fpath := fn.FormatByMap(savePath, map[string]any{"z": z, "x": x, "y": y})
	url := fn.FormatByMap(rootUrl, map[string]any{"z": z, "x": x, "y": y, "s": m.GetRandSubdomains()})
	dirPath := filepath.Join(dirName, fpath)
	if !ds.IsExist(dirPath) {

		err := ctp.Save(url, dirPath)
		if err == nil {
			msg := fmt.Sprintf("下载瓦片成功：%s", url)
			console.Log(msg)
			// 休眠一小会，防止被认为是爬虫抓取数据
			// time.Sleep(time.Millisecond * 100)
		} else {
			msg := fmt.Sprintf("下载瓦片失败：%s , %s ", url, err)
			console.Warn(msg)
		}
	}
	return nil
}

// GetChinaTileLimit
func GetChinaTileLimit(srid int, z int) (int, int, int, int) {
	return GetTileLimitByBbox(srid, z, gm.BBox{74, 4, 135, 54})
}

// GetChinaTileLimit
func GetTileLimitByBbox(srid int, z int, bbox gm.BBox) (int, int, int, int) {
	prj := crs.FromSRID(srid)
	startX := 0
	startY := 0
	endX := 0
	endY := 0
	//
	if z >= 10 {
		p1 := gm.LonLat{bbox.MinX(), bbox.MaxY()}
		p2 := gm.LonLat{bbox.MaxX(), bbox.MinY()}
		t1, _ := prj.LonlatToTileAndOffset(p1, float64(z))
		t2, _ := prj.LonlatToTileAndOffset(p2, float64(z))
		startX = t1.X
		startY = t1.Y
		endX = t2.X
		endY = t2.Y
	} else {
		switch srid {
		case 3857:
			size := int(math.Pow(2, float64(z)))
			endX = size
			endY = size
		case 4326:

			size := int(math.Pow(2, float64(z)))
			endX = size * 2
			endY = size
		}
	}
	return startX, startY, endX, endY
}
