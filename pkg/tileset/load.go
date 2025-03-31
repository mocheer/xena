package tileset

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/tileset/tile"
)

// Load 加载所有数据并保存到本地
// 下载远程地址中的tileset以及其所有的tile
// 支持断点续抓 -- 有些站点有limit限制一定时间内的请求次数，或者批量访问的时候会出现错误
// 这里需要考虑并发，并发请求数据会更快
func Load(urlstr string, filename string) error {

	data, err := ctp.Get(urlstr)
	if err == nil {
		t := tile.FromBytes(data)
		if !ds.IsExist(filename) {

			err = ds.Save(filename, data)
			if err != nil {
				return err
			}
		}

		return loadTile(t.Root, urlstr[:strings.LastIndex(urlstr, "/")], filepath.Dir(filename))
	}
	return err
}

// loadTile 下载tile以及tile的所有子级
func loadTile(m *tile.Tile, baseURL, dirName string) error {
	if m.Content != nil {
		err := loadContent(m, baseURL, dirName)
		if err != nil {
			return err
		}
	}
	children := m.Children

	// 当content是一个tileset.json一般当前children节点为空
	// 当content是一个gltf模型时，如果精度很高，一般会有children节点来分级展示更高精度的模型
	if len(children) > 0 {
		// 优先加载模型
		// slices.SortFunc(children, func(a, b *Tile) int {
		// 	if a.Content.IsTileset() == b.Content.IsTileset() {
		// 		return 0
		// 	}
		// 	if a.Content.IsTileset() {
		// 		return -1
		// 	}
		// 	return 1
		// })
		for _, item := range children {
			err := loadTile(item, baseURL, dirName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// loadContent 加载
func loadContent(m *tile.Tile, baseURL string, dir string) error {

	name, err := url.QueryUnescape(m.Content.GetURL())
	if err != nil {
		return err
	}
	// urlstr := path.Join(baseURL, name)
	urlstr := baseURL + "/" + name
	fmt.Println("url:", urlstr)

	filename := filepath.Join(dir, name)
	if m.Content.IsTileset() {
		err := Load(urlstr, filename)
		if err != nil {
			return err
		}
	} else { // 一个b3dm/gltf文件
		if !ds.IsExist(filename) {
			err = ctp.Save(urlstr, filename)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
