package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

// GraphCamera
// Camera 定义了场景中的视角和观察方式，但大部分引擎都有自己的默认相机，经常为空
type GraphCamera struct {
	*gltf.Camera
	Graph *Graph
}

// Index
func (m *GraphCamera) Index() int {
	return slices.Index(m.Graph.Cameras, m.Camera)
}

// IsUsed
func (m *GraphCamera) IsUsed() bool {
	return false
}

// Dispose
func (m *GraphCamera) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Cameras = slices.Delete(m.Graph.Cameras, index, index+1)
	}
}

// IsEmptyExtras
func (m *GraphCamera) IsEmptyExtras() bool {
	return m.Extras == nil
}
