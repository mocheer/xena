package tileset

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts/ctp"
)

// Load 下载远程地址中的tileset以及其所有的tile
// 断点续抓 -- 有些站点有limit限制一定时间内的请求次数，或者批量访问的时候会出现错误
func Load(remoteURL, dirName string) error {
	name := filepath.Base(remoteURL)         // tileset.json
	fileName := filepath.Join(dirName, name) // 文件保存路径
	var data []byte
	var err error
	var isFileExist = ds.IsExist(fileName)
	if isFileExist {
		data, err = os.ReadFile(fileName)
	} else {
		data, err = ctp.Get(remoteURL)

	}
	//
	if err == nil {
		tilesetObj := FromBytes(data)
		if tilesetObj == nil { //服务端错误，json文件不规范（limit插件问题） -- 重写请求
			fmt.Println("请求失败，未知异常")
			time.Sleep(time.Second)
			return Load(remoteURL, dirName)
		}
		if !isFileExist {
			ds.Save(fileName, data)
		}
		loading(tilesetObj.Root, remoteURL[:strings.LastIndex(remoteURL, "/")], dirName)
	}
	return err
}

// loading 下载tile以及tile的所有子级
func loading(t Tile, baseURL, dirName string) error {
	if t.Content != nil {
		contentURL := t.Content.Url
		if contentURL == "" {
			contentURL = t.Content.Uri
		}
		if contentURL != "" {
			contentRemoteURL := baseURL + "/" + contentURL
			relativePath := dirName + "/" + contentURL
			fmt.Println(contentRemoteURL)

			// 这是一个json描述文件
			if strings.HasSuffix(contentURL, ".json") {
				Load(contentRemoteURL, filepath.Dir(relativePath))
			} else { // 一个b3dm文件
				if !ds.IsExist(relativePath) {
					err := ctp.Save(contentRemoteURL, relativePath)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	children := t.Children
	if len(children) > 0 {
		for _, t := range children {
			err := loading(t, baseURL, dirName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
