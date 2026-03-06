package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

type GraphScene struct {
	Graph *Graph
	*gltf.Scene
}

// Index
func (m *GraphScene) Index() int {
	return slices.Index(m.Graph.Scenes, m.Scene)
}

// IsUsed
func (m *GraphScene) IsUsed() bool {
	return *m.Graph.Scene == uint32(m.Index())
}

// Dispose
func (m *GraphScene) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Scenes = slices.Delete(m.Graph.Scenes, index, index+1)
	}
}

// IsEmptyExtras
func (m *GraphScene) IsEmptyExtras() bool {
	return m.Extras == nil
}

// GetNodes
// 这里的节点一般不会重复，但理论上也可以重复，考虑去重
func (m *GraphScene) GetNodes() []*GraphNode {
	nodes := make([]*GraphNode, 0, len(m.Nodes))
	for _, index := range m.Nodes {
		nodes = append(nodes, m.Graph.GetNode(index))
	}
	return nodes
}

// EachNodes
func (m *GraphScene) EachNodes(fn func(node *GraphNode)) {
	for _, node := range m.GetNodes() {
		fn(node)
	}
}

// EachAllNodes
// 遍历所有子节点，这里包括多层子节点
func (m *GraphScene) EachUsedNodes(fn func(node *GraphNode)) {
	for _, node := range m.GetNodes() {
		node.EachUsedNodes(fn) //
	}
}

// FindAllNodes
// 遍历所有子节点，这里包括多层子节点，这里的节点不用担心是否重复，因为只判断一次
func (m *GraphScene) FindUsedNodes(fn func(node *GraphNode) bool) bool {
	for _, node := range m.GetNodes() {
		if node.FindUsedNodes(fn) {
			return true
		}
	}
	return false
}
