package tile

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/console"
	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/web_request"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/tile_arcgis"
)

type LoadConfig struct {
	URL     string
	DirName string
	MinZoom int
	MaxZoom int
	Origin  string
}

// LoadArcgis 下载arcgis本地图包中的所有瓦片数据
func LoadArcgis(localPath, dirName string) error {
	service, err := tile_arcgis.NewTileArcgis(filepath.Join(localPath, "conf.xml"))
	if err != nil {
		return err
	}

	for z := 0; z < 18; z++ {
		size := int(math.Pow(2, float64(z)))
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				zoomRowColumn := fmt.Sprintf("%d/%d/%d", z, i, j)
				dirPath := filepath.Join(dirName, zoomRowColumn)
				data, _ := service.ReadTile(gm.Tile{
					Z: z, X: i, Y: j,
				})
				if len(data) != 0 {
					if !ds.IsExist(dirPath) {
						ds.Save(dirPath, data)
						msg := fmt.Sprintf("下载瓦片成功：%s", zoomRowColumn)
						fmt.Println(msg)
					}
				} else {
					msg := fmt.Sprintf("瓦片为空：%s", zoomRowColumn)
					fmt.Println(msg)
				}
			}
		}
	}

	return nil
}

// Load 下载瓦片服务中的所有瓦片数据
func Load(lc *LoadConfig) error {

	minZoom := lc.MinZoom
	maxZoom := lc.MaxZoom
	for i := minZoom; i < maxZoom; i++ {
		loadingByZoom(lc, i)
	}
	return nil
}

// loadingByZoom 下载tile
func loadingByZoom(lc *LoadConfig, zoom int) error {

	size := int(math.Pow(2, float64(zoom)))
	count := size * size
	fns := make([]func(), 0, count)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fns = append(fns, func(i, j int) func() {
				return func() {
					loadingByXYZ(lc, &gm.Tile{X: i, Y: j, Z: zoom})
				}
			}(i, j))
		}
	}
	step := 1
	if count > 16 {
		step = count / 8
	}

	fn.GoFns(step, fns).Wait()
	return nil
}

func loadingByXYZ(lc *LoadConfig, tile *gm.Tile) error {

	if tile.Z >= 12 {
		lonlat := tile.GetLonLat()
		if lonlat[0] < 70 && lonlat[0] > 140 {
			return nil
		}

		if lonlat[1] < 0 && lonlat[0] > 60 {
			return nil
		}
	}

	rootUrl := lc.URL
	dirName := lc.DirName
	origin := lc.Origin

	x, y, z := tile.X, tile.Y, tile.Z

	zoomRowColumn := fmt.Sprintf("%d/%d/%d.png", z, x, y)
	url := fn.FmtString(rootUrl, map[string]any{"z": z, "x": x, "y": y})
	dirPath := filepath.Join(dirName, zoomRowColumn)
	if !ds.IsExist(dirPath) {

		err := web_request.Save(url, dirPath, origin)
		if err == nil {
			msg := fmt.Sprintf("下载瓦片成功：%s", url)
			console.Log(msg)
			// 休眠一小会，防止被认为是爬虫抓取数据
			// time.Sleep(time.Millisecond * 100)
		} else {
			msg := fmt.Sprintf("下载瓦片失败：%s", err)
			console.Log(msg)
		}
	}
	return nil
}
