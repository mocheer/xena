package graph

import (
	"errors"

	"github.com/qmuntal/gltf"
)

type GraphImage struct {
	Graph *Graph
	*gltf.Image
}

func (m GraphImage) ReadBufferView() ([]byte, error) {
	// 有 URI 就没有BuferView
	if m.BufferView != nil {
		bfv := m.Graph.BufferViews[*m.BufferView]
		data, err := bfv.Read()
		return data, err
	}
	return nil, errors.New("not found")
}

// WriteBufferView
func (m GraphImage) WriteBufferView(data []byte) {
	bfv := m.Graph.BufferViews[*m.BufferView]
	bufferviewIndex := bfv.Write(data)
	m.BufferView = &bufferviewIndex
}
