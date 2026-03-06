package pack

import (
	_ "embed"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/sys"
	"github.com/mocheer/xena/pkg/tileset/transform/graph"
)

//go:embed gltfpack.exe
var gltfPack []byte

type PackOptions struct {
	InputPath             string
	OutPath               string
	CompressPosition      bool
	CompressTextureFormat string
	Lossless              bool // 无损压缩，禁用所有损失质量的操作
}

// Pack
// TODO 可以在linux环境下做虚拟目录，拦截读写操作
// https://github.com/zeux/meshoptimizer/blob/master/gltf/README.md
func Pack(options *PackOptions) func(graph *graph.Graph) error {
	name := "./gltfpack.exe"

	if !ds.IsExist(name) {
		err := ds.Save(name, gltfPack)
		if err != nil {
			panic(err)
		}
	}
	return func(graph *graph.Graph) error {

		// 这里应该是对 graph 进行处理，而不是直接操作文件
		if options.InputPath == "" {
			options.InputPath = "./.temp/-input.gltf"
			graph.Save(options.InputPath)
		}
		if options.OutPath == "" {
			options.OutPath = "./.temp/-out.glb"
		}

		params := []string{
			`-i`, options.InputPath,
			`-o`, options.OutPath,
		}
		// EXT_meshopt_compression
		if options.CompressPosition {
			params = append(params, "-cc")
		}
		// params = append(params, "-sp")
		switch options.CompressTextureFormat {
		// KHR_texture_basisu 显存和体积变小
		case "basisu":
			params = append(params, "-tc")
		// EXT_texture_webp  体积显著变小
		case "webp":
			params = append(params, "-tw")
		}
		if options.Lossless {
			params = append(params, "-noq")
		}

		// fmt.Println(fmt.Sprintf(`%s -i "%s" -o "%s"  -cc -tc`, name, options.InputPath, options.OutPath))
		// err := sys.Exec(name, fmt.Sprintf(`%s -i "%s" -o "%s"  -cc -tc`, name, options.InputPath, options.OutPath))
		err := sys.Exec(name, params...)
		if err != nil {
			return err
		}
		// 无法解析 EXT_meshopt_compression 拓展
		// graph.Read(options.OutPath)
		return nil
	}
}
