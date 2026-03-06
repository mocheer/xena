package transform_test

import (
	"os"
	"testing"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/xena/pkg/tileset/transform"
	"github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/mocheer/xena/pkg/tileset/transform/save"
)

func TestDefault(t *testing.T) {
	transformer := transform.New(transform.TransformOptions{ 
		InPath:  "./.temp/内林泵站.gltf",
		OutPath: "./.temp/transform-out.glb",
	})
	err := transformer.RunWithDefault()
	t.Log(err)
}

func TestSave(t *testing.T) {
	ds.EachFiles("\\\\disk1214\\data\\PrefabHierarchyObject", func(filename string, fi os.FileInfo) {
		transformer := transform.New(transform.TransformOptions{
			InPath: filename,
		})
		err := transformer.Run(
			func(g *graph.Graph) error {
				g.EachPrimitives(func(p *graph.GraphPrimitive) {
					p.DisposeAttributeColor0()
				})
				return nil
			},
			save.Save(&save.SaveOptions{OutPath: "./.temp/PrefabHierarchyObject/" + fi.Name() + ".gltf"}),
		)
		t.Log(err)
	})

}

// TestSaveTextures
func TestSaveTextures(t *testing.T) {
	transformer := transform.FromPath("./testdata/scene.gltf")
	err := transformer.RunGetTextures()
	t.Log(err)
}
