package tile

import (
	"fmt"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/tile_arcgis"
)

// LoadArcgis 下载arcgis本地图包中的所有瓦片数据
func LoadArcgis(localPath, dirName string) error {
	service, err := tile_arcgis.NewTileArcgis(filepath.Join(localPath, "conf.xml"))
	srid := 3857
	if err != nil {
		return err
	}
	for z := 0; z < 18; z++ {
		startX, startY, endX, endY := GetChinaTileLimit(srid, z)
		for x := startX; x < endX; x++ {
			for y := startY; y < endY; y++ {
				zoomRowColumn := fmt.Sprintf("%d/%d/%d", z, x, y)
				dirPath := filepath.Join(dirName, zoomRowColumn)
				data, _ := service.ReadTile(gm.Tile{
					Z: z, X: x, Y: y,
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
