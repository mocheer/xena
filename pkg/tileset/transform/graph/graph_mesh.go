package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

type GraphMesh struct {
	*gltf.Mesh
	Graph *Graph
}

// Index
func (m *GraphMesh) Index() int {
	return slices.Index(m.Graph.Meshes, m.Mesh)
}

// IsUsed
func (m *GraphMesh) IsUsed() bool {
	index := uint32(m.Index())
	return m.Graph.FindUsedNodes(func(node *GraphNode) bool {
		return *node.Mesh == index
	})
}

// IsEmptyExtras
func (m *GraphMesh) IsEmptyExtras() bool {
	return m.Extras == nil
}

// Dispose
// 当Primitives为空时，可以直接销毁
// TODO 没有mesh的node，一般也需要销毁，所有用到这个位置后面顺序的mesh都要-1
func (m *GraphMesh) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Meshes = slices.Delete(m.Graph.Meshes, index, index+1)
		m.Graph.EachUsedNodes(func(node *GraphNode) {
			// 这里不直接销毁,而是标识为nil以待销毁
			if *node.Mesh == uint32(index) {
				node.Mesh = nil
			}
			if *node.Mesh > uint32(index) {
				*node.Mesh--
			}
		})
	}
}

// GetPrimitives
func (m *GraphMesh) GetPrimitives() []*GraphPrimitive {
	primitives := make([]*GraphPrimitive, 0, len(m.Primitives))
	for _, p := range m.Primitives {
		primitives = append(primitives, &GraphPrimitive{
			Graph:     m.Graph,
			Primitive: p,
		})
	}
	return primitives
}

// EachMaterial
func (m *GraphMesh) EachMaterial(fn func(material *GraphMaterial)) {
	for _, p := range m.GetPrimitives() {
		fn(p.GetMaterial())
	}
}

// FindMaterial
func (m *GraphMesh) FindMaterial(fn func(material *GraphMaterial) bool) bool {
	for _, p := range m.GetPrimitives() {
		if fn(p.GetMaterial()) {
			return true
		}
	}
	return false
}
