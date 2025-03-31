package graph

import (
	"github.com/qmuntal/gltf"
)

type GraphBuffer struct {
	Graph *Graph
	*gltf.Buffer
}

// IsUsed
func (m *GraphBuffer) IsUsed() bool {
	return false
}

func (m *GraphBuffer) Dispose() {

}

func (m *GraphBuffer) IsEmptyExtras() bool {
	return m.Extras == nil
}
