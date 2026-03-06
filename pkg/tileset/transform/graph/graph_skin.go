package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

// GraphSkin
// Skin 是用于描述模型骨骼动画的关键组件。它定义了模型的骨骼结构以及顶点如何与骨骼相关联
// joints 是一个数组，存储了骨骼节点的索引。这些索引指向 glTF 中的节点（nodes）
// skeleton 是一个节点索引，指向骨骼层级的根节点。
// inverseBindMatrices 是一个缓冲视图（BufferView）的索引，指向存储逆绑定矩阵（Inverse Bind Matrices）的数据。
type GraphSkin struct {
	Graph *Graph
	*gltf.Skin
}

// Index
func (m *GraphSkin) Index() int {
	return slices.Index(m.Graph.Skins, m.Skin)
}

// IsUsed
func (m *GraphSkin) IsUsed() bool {
	return true
}

// Dispose
func (m *GraphSkin) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Skins = slices.Delete(m.Graph.Skins, index, index+1)
	}
}

// IsEmptyExtras
func (m *GraphSkin) IsEmptyExtras() bool {
	return m.Extras == nil
}
