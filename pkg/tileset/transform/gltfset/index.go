package gltfset

import (
	"log"
	"os"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/xena/pkg/cesium"
	"github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/qmuntal/gltf"
)

// Gltfset
type Gltfset struct {
	Root           Root                  `json:"root"`
	BoundingSphere cesium.BoundingSphere `json:"boundingSphere"`
}

// Root
type Root struct {
	Children []ModelInfo `json:"children"`
}

// ModelInfo
type ModelInfo struct {
	Name           string                `json:"name"`
	BoundingSphere cesium.BoundingSphere `json:"boundingSphere"`
}

// FromPath
func FromPath(path string) *Gltfset {
	m := &Gltfset{}
	m.Root.Children = []ModelInfo{}
	ds.EachFiles(path, func(f string, fi os.FileInfo) {
		doc, err := gltf.Open(f)
		if err != nil {
			log.Println(err)
		}
		gh := graph.New(doc)
		bs := gh.ComputedBoundingShpere()
		m.Root.Children = append(m.Root.Children, ModelInfo{Name: fi.Name(), BoundingSphere: bs})
	})
	return m
}
