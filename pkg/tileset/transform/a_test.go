package transform_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_gltf"
	"github.com/mocheer/xena/pkg/tileset/transform"
	"github.com/qmuntal/gltf"
)

func TestXxx(t *testing.T) {

	doc, _ := gltf.Open("./testdata/tile_1_128_0_tex_children.glb")
	doc2, _ := gltf.Open("./testdata/tile_1_0_128_tex_children.glb")
	transform.MergeTo(doc, doc2)
	// doc.Nodes = []*gltf.Node{}
	// doc.Scenes = []*gltf.Scene{}
	// doc.Meshes = []*gltf.Mesh{}
	// doc.Materials = []*gltf.Material{}

	ds_gltf.Save("./testdata/merge2.gltf", doc, false)
	//
	// ds_gltf.Save("./testdata/a.gltf", doc, false)
	// ds_gltf.Save("./testdata/b.gltf", doc2, false)
}
