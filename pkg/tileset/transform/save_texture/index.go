package save_texture

import (
	"log/slog"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/xena/pkg/tileset/transform/graph"
)

type SaveTextureOptions struct {
	OutPath string
	Format  string //TODO
}

// SaveTexture
// graph.SaveImages() 同样的功能，不过这里作为插件，支持一些配置
func SaveTexture(options SaveTextureOptions) func(graph *graph.Graph) error {
	return func(graph *graph.Graph) error {

		images := graph.GetImages()
		slog.Info("模型图片个数", "count", len(images))
		if len(images) > 0 {
			for i, image := range images {
				data, err := image.Read()
				if err != nil {
					slog.Info("读取纹理失败", "err", err.Error())
					continue
				}
				// 文件名
				name := image.Name
				if name == "" {
					name = strconv.Itoa(i)
				}
				// 文件名后缀
				if !strings.Contains(name, ".") {
					suffix := strings.Replace(image.MimeType, "image/", ".", -1)
					name = name + suffix
				}
				fname := filepath.Join(options.OutPath, "images", name)
				slog.Info("保存图片", "fname", fname)
				err = ds.Save(fname, data)
				if err != nil {
					panic(err)
				}
			}
		}
		return nil
	}
}
