package graph

import (
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

type GraphBufferView struct {
	Graph *Graph
	*gltf.BufferView
}

func (m *GraphBufferView) IsEmptyExtras() bool {
	return m.Extras == nil
}

func (m GraphBufferView) Read() ([]byte, error) {
	return modeler.ReadBufferView(m.Graph.Doc, m.BufferView)
}

// Write
// 不会销毁当前BufferView，实际上是追加新的BufferView和Buffer数据
func (m GraphBufferView) Write(data []byte) uint32 {
	doc := m.Graph.Doc
	return modeler.WriteBufferView(doc, m.BufferView.Target, data)
}
