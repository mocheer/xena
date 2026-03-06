package graph

import (
	"slices"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/qmuntal/gltf"
)

type GraphNode struct {
	*gltf.Node
	Graph *Graph
}

// GetMesh
func (m *GraphNode) GetMesh() *GraphMesh {
	if m.Mesh != nil {
		return m.Graph.GetMesh(*m.Mesh)
	}
	return nil
}

// GetChldren
// 这里的子节点有可能是重复的，考虑去重
func (m *GraphNode) GetChldren() []*GraphNode {
	nodes := make([]*GraphNode, 0, len(m.Children))
	for _, index := range m.Children {
		graphNode := m.Graph.GetNode(index)
		nodes = append(nodes, graphNode)
	}
	return nodes
}

// EachNodes
func (m *GraphNode) EachNodes(fn func(node *GraphNode)) {
	for _, node := range m.GetChldren() {
		fn(node)
	}
}

// IsUsed
func (m *GraphNode) IsUsed() bool {
	return m.Graph.FindUsedNodes(func(node *GraphNode) bool {
		return node == m
	})
}

// IsEmptyExtras
func (m *GraphNode) IsEmptyExtras() bool {
	return m.Extras == nil
}

// EachUsedNodes
// 这里的子节点有可能是重复的
func (m *GraphNode) EachUsedNodes(fn func(node *GraphNode)) {
	// 用于防止重复遍历
	// nodeMap := map[*GraphNode]bool{}
	// isUsed := nodeMap[m]
	// if !isUsed {
	// 	nodeMap[m] = true
	// }
	fn(m)
	for _, node2 := range m.GetChldren() {
		node2.EachUsedNodes(fn)
	}
}

func (m *GraphNode) EachUsedMesh(fn func(node *GraphMesh)) {
	m.EachUsedNodes(func(node *GraphNode) {
		fn(node.GetMesh())
	})
}

// GetTransformMatrix 从节点数据获取变换矩阵
// 互斥性：一个node不应同时提供matrix和rotation/scale/translation中的任何一个
// 优先级：如果提供了rotation/scale/translation（即使部分提供，缺失的视为默认值），则必须忽略matrix。
// 默认值：如果matrix、rotation、scale、translation都未提供，则变换是单位矩阵（即无变换）
func (m *GraphNode) GetMatrix() mgl64.Mat4 {
	// 否则根据TRS（平移、旋转、缩放）构建矩阵
	if m.Translation != [3]float64{} || m.Rotation != [4]float64{} || m.Scale != [3]float64{} {
		if m.Scale == [3]float64{} {
			m.Scale = [3]float64{1, 1, 1}
		}
		translation := mgl64.Translate3D(m.Translation[0], m.Translation[1], m.Translation[2])
		rotation := mgl64.QuatRotate(m.Rotation[3], mgl64.Vec3{m.Rotation[0], m.Rotation[1], m.Rotation[2]}).Mat4()
		scale := mgl64.Scale3D(m.Scale[0], m.Scale[1], m.Scale[2])
		return translation.Mul4(rotation).Mul4(scale)
	}
	// 如果节点直接提供了矩阵，使用它
	if m.Matrix != [16]float64{} {
		return mgl64.Mat4(m.Matrix)
	}
	return mgl64.Ident4()
}

// GetWorldMatrix 计算节点的世界矩阵
// 世界矩阵不应该在这里设置，因为节点可能有多个父节点，世界矩阵应该根据具体的场景和父节点来计算
// func (m *GraphNode) GetWorldMatrix() mgl64.Mat4 {
// 	martix := m.GetMatrix()
// 	return martix
// }

// FindUsedNodes
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
// 目前删除node之后，只改变了Scene和Node下的所有节点的引用
// TODO 删除node之后，其内部的材质、纹理、访问器、缓冲区视图、缓冲区、顶点可能已经完全无用，这里可能需要重新遍历去掉无用的资源
func (m *GraphNode) Dispose() {
	graph := m.Graph
	index := slices.Index(graph.Nodes, m.Node)
	graph.Nodes = slices.Delete(graph.Nodes, index, index+1)
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
	for _, scene := range m.Graph.GetScenes() {
		scene.Nodes = resetNodeIndexFunc(scene.Nodes)
		// 每个nodes还有children
		for _, node := range scene.GetNodes() {
			resetNodeFunc(node)
		}
	}

}
