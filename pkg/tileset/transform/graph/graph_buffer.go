package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

// gltf 文件可以有多个缓冲区
// glb 文件只有一个缓冲区，当gltf导出成glb时会将资源分组，通过缓冲区视图访问（这里不确定是否是标准规范）
type GraphBuffer struct {
	Graph *Graph
	*gltf.Buffer
}

// IsUsed
func (m *GraphBuffer) IsUsed() bool {
	index := m.Index()
	if index != -1 {
		for _, bv := range m.Graph.BufferViews {
			if bv.Buffer == uint32(index) {
				return true
			}
		}
	}
	return false
}

// Index
func (m *GraphBuffer) Index() int {
	return slices.Index(m.Graph.Buffers, m.Buffer)
}

// Dispose
// 销毁后需要修正所有后置索引
func (m *GraphBuffer) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Buffers = slices.Delete(m.Graph.Buffers, index, index+1)
		// 修改索引
		for i := range m.Graph.BufferViews {
			if m.Graph.BufferViews[i].Buffer > uint32(index) {
				m.Graph.BufferViews[i].Buffer--
			}
		}
	}
}

// IsEmptyExtras
func (m *GraphBuffer) IsEmptyExtras() bool {
	return m.Extras == nil
}
