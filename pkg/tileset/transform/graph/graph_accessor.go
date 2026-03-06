package graph

import (
	"errors"
	"slices"

	"github.com/qmuntal/gltf"
)

// mesh.primitive.attributes的 POSITION、NORMAL、TEXCOORD_0 以及 mesh.primitive.indices 使用了 accessor
// accessor 描述如何解析 bufferView 的数据格式。
// image是没用 accessor 直接访问bufferView的，因为它是原始二进制不需要额外解释

// GraphAccessor
// Sparse 用于读取稀疏数据，替换默认的BufferView数据
// 默认数据: [A, B, C, D, E, F, G, H, I, J]
// 稀疏索引: [3, 5, 10]       ← 要替换的位置
// 稀疏值:   [X, Y, Z]
// 结果:     [A, B, C, X, E, Y, G, H, I, Z]
// 在动态更新场景（如游戏、动画）中，只需传输稀疏部分，就可以更新数据。
type GraphAccessor struct {
	Graph *Graph
	*gltf.Accessor
}

// Index
func (m *GraphAccessor) Index() int {
	return slices.Index(m.Graph.Accessors, m.Accessor)
}

// IsUsed
func (m *GraphAccessor) IsUsed() bool {
	return true
}

// Dispose
func (m *GraphAccessor) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Accessors = slices.Delete(m.Graph.Accessors, index, index+1)
	}
}

// IsEmptyExtras
func (m *GraphAccessor) IsEmptyExtras() bool {
	return m.Extras == nil
}

// GetBufferView
func (m *GraphAccessor) GetBufferView() *GraphBufferView {
	return m.Graph.GetBufferView(*m.Accessor.BufferView)
}

// ReadBufferView
// 这里只读取BufferView，可能不准确，还有一部分数据放在Sparse，这部分数据后面要合并
func (m *GraphAccessor) ReadBufferView() ([]byte, error) {
	return m.GetBufferView().Read()
}

// ReadAsIndices 读取索引
func (m *GraphAccessor) ReadAsIndices() ([]uint32, error) {
	return m.Graph.ReadAsIndices(m.Accessor)
}

// ReadAsPosition 读取顶点坐标集
func (m *GraphAccessor) ReadAsPosition() ([][3]float32, error) {
	return m.Graph.ReadAsPosition(m.Accessor)
}

// ReadAsTextureCoord 读取纹理坐标集
func (m *GraphAccessor) ReadAsTextureCoord() ([][2]float32, error) {
	return m.Graph.ReadAsTextureCoord(m.Accessor)
}

// Valid
func (m GraphAccessor) Valid() error {
	if m.GetBufferView() == nil {
		return errors.New("accessor定义的bufferview不存在")
	}
	return nil
}
