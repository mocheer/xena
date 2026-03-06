package weld_test

import (
	"testing"

	"github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/mocheer/xena/pkg/tileset/transform/weld"
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

//https://github.com/donmccurdy/glTF-Transform/blob/main/packages/functions/test/weld.test.ts

func TestWeld(t *testing.T) {
	doc := gltf.NewDocument()
	// Create position accessor
	positionArray := [][3]float32{
		{0, 0, 0},
		{0, 0, 1},
		{0, 0, -1},
		{0, 0, 0},
		{0, 0, 1},
		{0, 0, -1},
	}
	positionAccessor := modeler.WritePosition(doc, positionArray)
	// Create indices accessor
	indicesArray := []uint32{3, 4, 5, 0, 1, 2}
	indicesAccessor := modeler.WriteIndices(doc, indicesArray)

	// Create primitives
	prim1 := &gltf.Primitive{
		// Indices: &indicesAccessor,
		Attributes: gltf.Attribute{
			"POSITION": positionAccessor,
		},
		Mode: gltf.PrimitiveTriangles,
	}

	prim2 := &gltf.Primitive{
		Indices: &indicesAccessor,
		Attributes: gltf.Attribute{
			"POSITION": positionAccessor,
		},
		Mode: gltf.PrimitiveTriangles,
	}

	// Create mesh and add primitives
	mesh := &gltf.Mesh{
		Primitives: []*gltf.Primitive{prim1, prim2},
	}

	g := graph.New(doc)
	g.AddMesh(mesh)

	weld.Weld(weld.WeldOptions{})(g)
	g.EachUsedNodes(func(node *graph.GraphNode) {
		for i, p := range node.GetMesh().GetPrimitives() {
			t.Log(i)
			t.Log(p.ReadIndices())
			t.Log(p.ReadPostion())
		}
	})
}
