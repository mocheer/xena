package graph_test

import (
	"os"
	"strings"
	"testing"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/xena/pkg/cesium"
	"github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/qmuntal/gltf"
)

type Gltfset struct {
	Root           Root                  `json:"root"`
	BoundingSphere cesium.BoundingSphere `json:"boundingSphere"`
}

type Root struct {
	Children []ModelInfo `json:"children"`
}

type ModelInfo struct {
	Name           string                `json:"name"`
	BoundingSphere cesium.BoundingSphere `json:"boundingSphere"`
}

func TestBoundingShpere(t *testing.T) {
	g := Gltfset{}
	g.Root.Children = []ModelInfo{}
	ds.EachFiles(`D:\code\web\@planet\demo\gltf-to-3d-tiles\1020`, func(f string, fi os.FileInfo) {
		doc, err := gltf.Open(f)
		if err != nil {
			t.Log(err)
		}
		gh := graph.New(doc)
		bs := gh.ComputedBoundingShpere()
		g.Root.Children = append(g.Root.Children, ModelInfo{Name: fi.Name(), BoundingSphere: bs})
		t.Log(f, bs)
		err = gh.Save("../testdata/models/" + strings.Replace(fi.Name(), "glb", "gltf", 1))
		t.Log(err)
	})
	ds_json.SaveWithIndent("../testdata/gltfset.json", g)
}
