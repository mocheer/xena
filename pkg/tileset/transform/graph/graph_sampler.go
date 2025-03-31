package graph

import (
	"github.com/qmuntal/gltf"
)

type GraphSampler struct {
	Graph *Graph
	*gltf.Sampler
}

func (m *GraphSampler) IsUsed() bool {
	return false
}

func (m *GraphSampler) Dispose() {

}

func (m *GraphSampler) IsEmptyExtras() bool {
	return m.Extras == nil
}
