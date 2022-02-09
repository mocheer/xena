package tile

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/rh"
)

type LoadConfig struct {
	DirName string
	MinZoom int
	MaxZoom int
}

// Load 下载瓦片服务中的所有瓦片数据
func Load(remoteURL, dirName string) error {
	for i := 0; i < 10; i++ {
		loadingByZoom(remoteURL, i, dirName)
	}
	return nil
}

// loadingByZoom 下载tile
func loadingByZoom(rootUrl string, zoom int, dirName string) error {
	size := int(math.Pow(2, float64(zoom)))
	maxErrorCount := 10
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			zoomRowColumn := fmt.Sprintf("%d/%d/%d", zoom, i, j)
			url := rootUrl + zoomRowColumn
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
					maxErrorCount--
					if maxErrorCount < 0 {
						return nil
					}

				}
			}

		}
	}
	return nil
}
