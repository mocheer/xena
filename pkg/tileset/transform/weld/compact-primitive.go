package weld

import (
	"math"

	. "github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/qmuntal/gltf"
)

// compactPrimitive 重写一个 Primitive，移除其顶点属性中未使用的顶点。
// 当多个 Primitive 共享顶点属性，且每个 Primitive 只使用其中一部分时，
// 可以通过压缩生成每个 Primitive 都有更小、独立的顶点流。
func compactPrimitive(prim *GraphPrimitive, remap []uint32, dstVertexCount uint32) *GraphPrimitive {
	doc := prim.Graph.Doc

	if remap == nil || dstVertexCount == 0 {
		remap, dstVertexCount = createCompactPlan(prim)
	}

	// 重新映射索引。
	srcIndices, err := prim.ReadIndices()
	if err != nil {
		return nil
	}
	srcIndicesArray := srcIndices
	srcIndicesCount := getPrimitiveVertexCount(prim, VertexCountMethodRender)

	dstIndices := doc.CreateAccessor()
	dstIndicesCount := srcIndicesCount // Primitive 的顶点数不会改变。
	dstIndicesArray := make([]uint32, dstIndicesCount)

	for i := uint32(0); i < dstIndicesCount; i++ {
		dstIndicesArray[i] = remap[srcIndicesArray[i]]
	}

	prim.Indices = dstIndices.SetData(dstIndicesArray)

	// 重新映射顶点。
	srcAttributesPrev := deepListAttributes(prim)

	for _, srcAttribute := range prim.Attributes {
		dstAttribute := shallowCloneAccessor(doc, srcAttribute)
		compactAttribute(srcAttribute, srcIndices, remap, dstAttribute, dstVertexCount)
		prim.Attributes[srcAttribute.Name] = dstAttribute
	}
	for _, target := range prim.Targets {
		for _, srcAttribute := range target {
			dstAttribute := shallowCloneAccessor(doc, srcAttribute)
			compactAttribute(srcAttribute, srcIndices, remap, dstAttribute, dstVertexCount)
			target[srcAttribute.Name] = dstAttribute
		}
	}

	// 清理 Accessor。
	if srcIndices != nil && len(srcIndices.GetParents()) == 1 {
		srcIndices.Dispose()
	}
	for _, srcAttribute := range srcAttributesPrev {
		if len(srcAttribute.GetParents()) == 1 {
			srcAttribute.Dispose()
		}
	}

	return prim
}

// compactAttribute 将 srcAttribute 复制到 dstAttribute，使用给定的索引和映射（srcIndex -> dstIndex）。
// dstAttribute 中的现有数组将被替换。未被索引使用的顶点将被移除，生成一个紧凑的属性。
func compactAttribute(srcAttribute *gltf.Accessor, srcIndices *gltf.Accessor, remap []uint32, dstAttribute *gltf.Accessor, dstVertexCount uint32) *gltf.Accessor {
	elementSize := srcAttribute.ElementSize
	srcArray := srcAttribute.Data
	srcIndicesArray := srcIndices.Data
	srcIndicesCount := uint32(len(srcIndicesArray))
	if srcIndices == nil {
		srcIndicesCount = uint32(len(srcArray))
	}
	dstArray := make([]float32, dstVertexCount*elementSize)
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

	dstAttribute.SetData(dstArray)
	return dstAttribute
}

// createCompactPlan 为索引化的 Primitive 创建一个“映射”和“目标顶点数”计划，
// 以便通过 compactPrimitive 删除任何未渲染的顶点。
func createCompactPlan(prim *GraphPrimitive) ([]uint32, uint32) {
	srcVertexCount := getPrimitiveVertexCount(prim, VertexCountMethodUpload)

	indices := prim.Indices
	indicesArray := indices.Data
	if indices == nil || indicesArray == nil {
		return createIndices(srcVertexCount, math.MaxUint32), srcVertexCount
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

// Helper functions (假设这些函数已实现)
func getPrimitiveVertexCount(prim *gltf.Primitive, method VertexCountMethod) uint32 {
	// 根据方法计算顶点数
	return 0
}

func createIndicesEmpty(count, max uint32) []uint32 {
	return make([]uint32, count)
}

func createIndices(count, max uint32) []uint32 {
	indices := make([]uint32, count)
	for i := range indices {
		indices[i] = uint32(i)
	}
	return indices
}

func deepListAttributes(prim *gltf.Primitive) []*gltf.Accessor {
	// 深度列出所有属性
	return nil
}

func shallowCloneAccessor(doc *gltf.Document, src *gltf.Accessor) *gltf.Accessor {
	// 浅克隆一个 Accessor
	return nil
}

// VertexCountMethod 用于指定计算顶点数的方法
type VertexCountMethod int

const (
	VertexCountMethodRender VertexCountMethod = iota // 根据渲染顶点数计算
	VertexCountMethodUpload                          // 根据上传顶点数计算
)
