package transform

import (
	"fmt"

	"github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/mocheer/xena/pkg/tileset/transform/pack"
	"github.com/mocheer/xena/pkg/tileset/transform/prune"
	"github.com/mocheer/xena/pkg/tileset/transform/save"
	"github.com/mocheer/xena/pkg/tileset/transform/save_texture"
	"github.com/mocheer/xena/pkg/tileset/transform/texture_merge"
)

type TransformOptions struct {
	InPath  string
	OutPath string
}
type Transform struct {
	options TransformOptions
}

type Transformer = func(graph *graph.Graph) error

// New
func New(options TransformOptions) *Transform {
	return &Transform{
		options: options,
	}
}

// FromPath
func FromPath(p string) *Transform {
	// 不能存在同名文件名和目录
	oupPath := fmt.Sprintf("%s.assets", p)
	return &Transform{
		options: TransformOptions{
			InPath:  p,
			OutPath: oupPath,
		},
	}
}

func (t *Transform) SetOutpath(value string) {
	t.options.OutPath = value
}

// RunWithDefault
func (t *Transform) RunWithDefault() error {
	return t.Run(
		prune.Prune(&prune.PruneOptions{
			KeepLeaves:        false,
			KeepAttributes:    false,
			KeepIndices:       false,
			KeepExtras:        false,
			KeepSolidTextures: false,
		}),
		texture_merge.TextureMerge(texture_merge.TextureMergeOptions{}),
		pack.Pack(&pack.PackOptions{
			OutPath:               t.options.OutPath,
			CompressPosition:      true,
			CompressTextureFormat: "webp",
		}),
		save.Save(&save.SaveOptions{OutPath: t.options.OutPath}),
	)
}

// RunGetTextures
func (m *Transform) RunGetTextures() error {
	return m.Run(
		save_texture.SaveTexture(save_texture.SaveTextureOptions{OutPath: m.options.OutPath}),
	)
}

// Run
func (t *Transform) Run(transformers ...Transformer) error {
	g := graph.Open(t.options.InPath)
	for _, transform := range transformers {
		err := transform(g)
		if err != nil {
			return err
		}
	}
	return nil
}
