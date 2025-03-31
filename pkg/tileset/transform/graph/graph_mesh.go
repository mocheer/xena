package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

type GraphMesh struct {
	*gltf.Mesh
	Graph      *Graph
	primitives []*GraphPrimitive
}

func (m *GraphMesh) GetPrimitives() []*GraphPrimitive {
	if m.primitives == nil {
		primitives := make([]*GraphPrimitive, 0, len(m.Primitives))
		for _, p := range m.Primitives {
			gp := &GraphPrimitive{
				Graph:     m.Graph,
				Primitive: p,
			}
			primitives = append(primitives, gp)
		}
		m.primitives = primitives
	}
	return m.primitives
}

func (m *GraphMesh) FindMaterial(fn func(material *GraphMaterial) bool) bool {
	for _, p := range m.GetPrimitives() {
		if fn(p.GetMaterial()) {
			return true
		}
	}
	return false
}

func (m *GraphMesh) IsUsed() bool {
	return m.Graph.FindUsedNodes(func(node *GraphNode) bool {
		return node.GetMesh() == m
	})
}

func (m *GraphMesh) IsEmptyExtras() bool {
	return m.Extras == nil
}

// Dispose
// 当Primitives为空时，可以直接销毁
func (m *GraphMesh) Dispose() {
	graph := m.Graph
	index := slices.Index(graph.Meshes, m)
	graph.Meshes = slices.Delete(graph.Meshes, index, index+1)
	graph.Doc.Meshes = slices.Delete(graph.Doc.Meshes, index, index+1)
	// TODO 没有mesh的node，一般也需要销毁，所有用到这个位置后面顺序的mesh都要-1
	graph.EachUsedNodes(func(node *GraphNode) {
		// 这里不直接销毁,而是标识为nil以待销毁
		if *node.Mesh == uint32(index) {
			node.Mesh = nil
		}
		if *node.Mesh > uint32(index) {
			*node.Mesh--
		}
	})
}
