package weld

import (
	. "github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/qmuntal/gltf"
)

type WeldOptions struct {
	Tolerance float64
}

// Weld
func Weld(options WeldOptions) func(graph *Graph) error {
	return func(graph *Graph) error {
		graph.EachUsedNodes(func(node *GraphNode) {
			for _, p := range node.GetMesh().GetPrimitives() {
				weldPrimitive(p, options)
			}
		})
		return nil
	}
}

// weldPrimitive
func weldPrimitive(p *GraphPrimitive, options WeldOptions) {
	// 如果是点，那就不需要weld
	if p.Mode == gltf.PrimitivePoints {
		return
	}
	position, err := p.ReadPostion()
	if err != nil {
		return
	}
	srcVertexCount := len(position)
	// 不一定有 indices ,没有的时候
	srcIndices, err := p.ReadIndices()
	srcIndicesCount := len(srcIndices)
	if err != nil {
		srcIndicesCount = srcVertexCount
	}

	tableSize := ceilPowerOfTwo(srcVertexCount + srcVertexCount/4)
	table := make([]uint32, tableSize)
	for i := range table {
		table[i] = EMPTY_U32
	}
	writeMap := make([]uint32, srcVertexCount)
	for i := range writeMap {
		writeMap[i] = EMPTY_U32
	}

	dstVertexCount := uint32(0)
	for i := range srcIndicesCount {
		// 原始索引
		var srcIndex uint32
		if srcIndices != nil {
			srcIndex = srcIndices[i]
		} else {
			srcIndex = uint32(i)
		}
		// 已经写入了，就不要再重复
		if writeMap[srcIndex] != EMPTY_U32 {
			continue
		}
		// 目标索引
		hashIndex, _ := hashLookup(table, tableSize, NewVertexStream(p), srcIndex, EMPTY_U32)
		dstIndex := table[hashIndex]

		if dstIndex == EMPTY_U32 {
			table[hashIndex] = srcIndex //指向原始索引，顶点在Tolerance误差下看似成同一个点，但这个点指向最开始的第一个点
			writeMap[srcIndex] = dstVertexCount
			dstVertexCount++
		} else {
			writeMap[srcIndex] = writeMap[dstIndex]
		}
	}

	compactPrimitive(p, writeMap, dstVertexCount)
}
