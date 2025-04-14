package graph

import (
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

// GraphAccessor
// Sparse 用于读取稀疏数据，替换默认的BufferView数据
// 默认数据: [A, B, C, D, E, F, G, H, I, J]
// 稀疏索引: [3, 5, 10]       ← 要替换的位置
// 稀疏值:   [X, Y, Z]
// 结果:    [A, B, C, X, E, Y, G, H, I, Z]
// 在动态更新场景（如游戏、动画）中，只需传输稀疏部分，就可以更新数据。
type GraphAccessor struct {
	Graph *Graph
	*gltf.Accessor
}

func (m *GraphAccessor) GetBufferView() *GraphBufferView {
	return m.Graph.BufferViews[*m.Accessor.BufferView]
}

// ReadBufferView
// 这里只读取BufferView，可能不准确，还有一部分数据放在Sparse，这部分数据后面要合并
func (m *GraphAccessor) ReadBufferView() ([]byte, error) {
	return modeler.ReadBufferView(m.Graph.Doc, m.GetBufferView().BufferView)
}

// ReadAsIndices
func (m *GraphAccessor) ReadAsIndices() ([]uint32, error) {

	return modeler.ReadIndices(m.Graph.Doc, m.Accessor, nil)
}

func (m *GraphAccessor) ReadAsPosition() ([][3]float32, error) {
	return modeler.ReadPosition(m.Graph.Doc, m.Accessor, nil)
}

func (m *GraphAccessor) ReadAsTextureCoord() ([][2]float32, error) {
	return modeler.ReadTextureCoord(m.Graph.Doc, m.Accessor, nil)
}

func (m *GraphAccessor) IsUsed() bool {
	return false
}

func (m *GraphAccessor) Dispose() {

}

func (m *GraphAccessor) IsEmptyExtras() bool {
	return m.Extras == nil
}
