package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

type GraphNode struct {
	*gltf.Node
	Graph *Graph
	// Index *uint32
}

func (m *GraphNode) GetMesh() *GraphMesh {
	return m.Graph.Meshes[*m.Mesh]
}

func (m *GraphNode) GetChldren() []*GraphNode {
	nodes := make([]*GraphNode, 0, len(m.Children))
	for _, index := range m.Children {
		graphNode := m.Graph.Nodes[index]
		nodes = append(nodes, graphNode)
	}
	return nodes
}

// IsUsed
func (m *GraphNode) IsUsed() bool {
	return m.Graph.FindUsedNodes(func(node *GraphNode) bool {
		return node == m
	})
}
func (m *GraphNode) IsEmptyExtras() bool {
	return m.Extras == nil
}

// IsVisible
func (m *GraphNode) IsVisible() bool {
	return m.Graph.Scene.FindUsedNodes(func(node *GraphNode) bool {
		return node == m
	})
}

// EachChildren
func (m *GraphNode) EachUsedNodes(fn func(node *GraphNode)) {
	fn(m)
	for _, node2 := range m.GetChldren() {
		node2.EachUsedNodes(fn)
	}
}

// FindChildren
func (m *GraphNode) FindUsedNodes(fn func(node *GraphNode) bool) bool {
	if fn(m) {
		return true
	}
	for _, node2 := range m.GetChldren() {
		if node2.FindUsedNodes(fn) {
			return true
		}
	}
	return false
}

// Dispose
func (m *GraphNode) Dispose() {
	graph := m.Graph
	index := slices.Index(graph.Nodes, m)
	graph.Nodes = slices.Delete(graph.Nodes, index, index+1)
	graph.Doc.Nodes = slices.Delete(graph.Doc.Nodes, index, index+1)
	// TODO 这里需要改变所有Scene和Node对应的索引,甚至应该删除当前节点所引用的顶点、纹理对应的buffer数据（如果没有节点引用）
	u32Index := uint32(index)
	resetNodeIndexFunc := func(nodes []uint32) []uint32 {
		resetNodes := []uint32{}
		for _, nodeIndex := range nodes {
			// 如果包含当前node，则直接删除
			if nodeIndex != u32Index {
				if nodeIndex > u32Index {
					resetNodes = append(resetNodes, nodeIndex-1)
				} else {
					resetNodes = append(resetNodes, nodeIndex)
				}
			}
		}
		return resetNodes
	}
	var resetNodeFunc func(node *GraphNode)
	resetNodeFunc = func(node *GraphNode) {
		node.Children = resetNodeIndexFunc(node.Children)
		for _, child := range node.GetChldren() {
			resetNodeFunc(child)
		}
	}
	for _, scene := range m.Graph.Scenes {
		scene.Nodes = resetNodeIndexFunc(scene.Nodes)
		// 每个nodes还有children
		for _, node := range scene.GetNodes() {
			resetNodeFunc(node)
		}

	}

}
