package graph

import (
	"github.com/qmuntal/gltf"
)

type GraphScene struct {
	Graph *Graph
	*gltf.Scene
}

func (m *GraphScene) GetNodes() []*GraphNode {
	nodes := make([]*GraphNode, 0, len(m.Nodes))
	for _, index := range m.Nodes {
		nodes = append(nodes, m.Graph.Nodes[index])
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
		node.EachUsedNodes(fn)
	}
}

// FindAllNodes
// 遍历所有子节点，这里包括多层子节点
func (m *GraphScene) FindUsedNodes(fn func(node *GraphNode) bool) bool {
	for _, node := range m.GetNodes() {
		if node.FindUsedNodes(fn) {
			return true
		}
	}
	return false
}
