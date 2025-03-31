package graph

import (
	"github.com/qmuntal/gltf"
)

// GraphCamera 定义了场景中的视角和观察方式，但大部分引擎都有自己的默认相机，经常没有用到
type GraphCamera struct {
	graph *Graph
	*gltf.Camera
}

// IsUsed
func (m *GraphCamera) IsUsed() bool {
	return false
}

func (m *GraphCamera) Dispose() {

}

func (m *GraphCamera) IsEmptyExtras() bool {
	return m.Extras == nil
}
