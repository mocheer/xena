package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

type GraphSampler struct {
	Graph *Graph
	*gltf.Sampler
}

// Index
func (m *GraphSampler) Index() int {
	return slices.Index(m.Graph.Samplers, m.Sampler)
}

// IsUsed
func (m *GraphSampler) IsUsed() bool {
	return true
}

// Dispose
// TODO: 目前没有处理被引用的情况
func (m *GraphSampler) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Samplers = slices.Delete(m.Graph.Samplers, index, index+1)
	}
}

// IsEmptyExtras
func (m *GraphSampler) IsEmptyExtras() bool {
	return m.Extras == nil
}
