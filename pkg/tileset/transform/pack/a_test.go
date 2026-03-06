package pack_test

import (
	"testing"

	"github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/mocheer/xena/pkg/tileset/transform/pack"
	"github.com/qmuntal/gltf"
)

func TestPack(t *testing.T) {
	doc, _ := gltf.Open("./testdata/ship.glb")
	g := graph.New(doc)
	run := pack.Pack(&pack.PackOptions{
		InputPath: "./testdata/ship.glb",
		OutPath:   "./testdata/ship-pack.glb",
		// CompressPosition: true,
		// CompressTextureFormat: "basisu",
		Lossless: true,
	})
	err := run(g)
	if err != nil {
		t.Log(err)
	}
}
