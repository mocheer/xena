package weld

import (
	"math"

	. "github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

// compactPrimitive 重写一个 Primitive，移除其顶点属性中未使用的顶点。
// 当多个 Primitive 共享顶点属性，且每个 Primitive 只使用其中一部分时，
// 可以通过压缩生成每个 Primitive 都有更小、独立的顶点流。
//
// remap是重新映射的hash索引
// dstVertexCount是去重后的顶点个数
func compactPrimitive(prim *GraphPrimitive, remap []uint32, dstVertexCount uint32) *GraphPrimitive {
	graph := prim.Graph
	//
	if remap == nil || dstVertexCount == 0 {
		remap, dstVertexCount = createCompactPlan(prim)
	}

	// 重新映射索引。
	srcIndices := prim.GetIndices()
	var srcIndicesArray []uint32
	var srcIndicesCount uint32
	if srcIndices == nil {
		count := prim.GetVertexCount()
		srcIndicesArray = make([]uint32, count)
		for i := range count {
			srcIndicesArray[i] = i
		}
		srcIndicesCount = count
	} else {
		srcIndicesArray, _ = srcIndices.ReadAsIndices()
		srcIndicesCount = srcIndices.Count
	}

	dstIndicesCount := srcIndicesCount // Primitive 的顶点数不会改变。
	dstIndicesArray := make([]uint32, dstIndicesCount)

	for i := uint32(0); i < dstIndicesCount; i++ {
		dstIndicesArray[i] = remap[srcIndicesArray[i]]
	}

	prim.Indices = gltf.Index(graph.WriteIndices(dstIndicesArray))

	// 重新映射顶点。
	accessors := prim.GetAttributeAccessors()

	for _, srcAttribute := range accessors {
		prim.Attributes[srcAttribute.Name] = compactAttribute(srcAttribute, srcIndices, remap, dstVertexCount)
	}
	// TODO primitive.Target
	// 这里应该优化动画目标数据

	return prim
}

// compactAttribute 将 srcAttribute 复制到 dstAttribute，使用给定的索引和映射（srcIndex -> dstIndex）。
// dstAttribute 中的现有数组将被替换。未被索引使用的顶点将被移除，生成一个紧凑的属性。
func compactAttribute(srcAttribute *GraphAccessor, srcIndices *GraphAccessor, remap []uint32, dstVertexCount uint32) uint32 {
	elementSize := srcAttribute.ComponentType.ByteSize()
	srcArray, _ := srcAttribute.ReadBufferView()

	var srcIndicesArray []uint32
	var srcIndicesCount uint32
	//
	if srcIndices == nil {
		srcIndicesCount = uint32(len(srcArray))
	} else {
		srcIndicesArray, _ = srcIndices.ReadAsIndices()
		srcIndicesCount = uint32(len(srcIndicesArray))
	}
	dstArray := make([]byte, dstVertexCount*elementSize)
	dstDone := make([]bool, dstVertexCount)

	for i := uint32(0); i < srcIndicesCount; i++ {
		srcIndex := uint32(srcIndicesArray[i])
		if srcIndices == nil {
			srcIndex = i
		}
		dstIndex := remap[srcIndex]
		if dstDone[dstIndex] {
			continue
		}

		for j := uint32(0); j < elementSize; j++ {
			dstArray[dstIndex*elementSize+j] = srcArray[srcIndex*elementSize+j]
		}

		dstDone[dstIndex] = true
	}

	return modeler.WriteAccessor(srcAttribute.Graph.Document, srcAttribute.GetBufferView().Target, dstArray)
}

// createCompactPlan 为索引化的 Primitive 创建一个“映射”和“目标顶点数”计划，
// 以便通过 compactPrimitive 删除任何未渲染的顶点。
func createCompactPlan(prim *GraphPrimitive) ([]uint32, uint32) {
	srcVertexCount := prim.GetVertexCount()
	indicesArray, err := prim.ReadIndices()
	// 索引读取失败，一般是该索引没有数据
	if err != nil {
		return createIndices(srcVertexCount), srcVertexCount
	}

	remap := make([]uint32, srcVertexCount)
	for i := range remap {
		remap[i] = math.MaxUint32
	}

	dstVertexCount := uint32(0)

	for _, srcIndex := range indicesArray {
		if remap[srcIndex] == math.MaxUint32 {
			remap[srcIndex] = dstVertexCount
			dstVertexCount++
		}
	}

	return remap, dstVertexCount
}

func createIndices(count uint32) []uint32 {
	indices := make([]uint32, count)
	for i := range indices {
		indices[i] = uint32(i)
	}
	return indices
}
