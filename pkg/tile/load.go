package tile

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/rh"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/tile_arcgis"
)

type LoadConfig struct {
	DirName string
	MinZoom int
	MaxZoom int
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
func Load(remoteURL, dirName string, maxZoom int) error {
	fns := []func(){}
	for i := 0; i < maxZoom; i++ {
		fns = append(fns, func(i int) func() {
			return func() {
				loadingByZoom(remoteURL, i, dirName)
			}
		}(i))
	}
	fn.GoFns(1, fns).Wait()
	return nil
}

// loadingByZoom 下载tile
func loadingByZoom(rootUrl string, zoom int, dirName string) error {
	size := int(math.Pow(2, float64(zoom)))
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			zoomRowColumn := fmt.Sprintf("%d/%d/%d", zoom, i, j)
			url := fn.FmtString(rootUrl, map[string]any{"z": zoom, "x": i, "y": j})
			dirPath := filepath.Join(dirName, zoomRowColumn)
			if !ds.IsExist(dirPath) {
				data, err := rh.Get(url)
				if err == nil {
					ds.Save(dirPath, data) // 保存到本地硬盘
					msg := fmt.Sprintf("下载瓦片成功：%s", url)
					fmt.Println(msg)
					// 休眠一小会，防止被认为是爬虫抓取数据
					// time.Sleep(time.Millisecond * 100)
				} else {
					msg := fmt.Sprintf("下载瓦片失败：%s", err)
					fmt.Println(msg)
				}
			}
		}
	}
	return nil
}
