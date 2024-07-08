package tileset

import (
	"errors"
	"path"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts/ctp"
)

// Load 加载所有数据并保存到本地
// 下载远程地址中的tileset以及其所有的tile
// 支持断点续抓 -- 有些站点有limit限制一定时间内的请求次数，或者批量访问的时候会出现错误
func Load(url string, filename string) error {
	data, err := ctp.Get(url)
	if err == nil {
		t := FromBytes(data)
		err = ds.Save(filename, data)
		if err != nil {
			return err
		}
		return loadTile(t.Root, path.Dir(url), filepath.Dir(filename))
	}
	return err
}

// loadTile 下载tile以及tile的所有子级
func loadTile(m *Tile, baseURL, dirName string) error {
	if m.Content != nil {
		loadContent(m, baseURL, dirName)
	}
	children := m.Children
	// 当content是一个tileset.json一般当前children节点为空
	// 当content是一个gltf模型时，如果精度很高，一般会有children节点来分级展示更高精度的模型
	if len(children) > 0 {
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
func loadContent(m *Tile, baseURL string, dir string) error {
	url := path.Join(baseURL, m.Content.GetURL())
	if url != "" {
		filename := filepath.Join(dir, url)
		if !ds.IsExist(filename) {
			if m.Content.IsTileset() {
				err := Load(url, filename)
				if err != nil {
					return err
				}
			} else { // 一个b3dm文件
				err := ctp.Save(url, filename)
				if err != nil {
					return err
				}
			}
		}
	}
	return errors.New("url is invalid")
}
