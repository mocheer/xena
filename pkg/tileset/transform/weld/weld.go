package weld

import (
	"math"

	. "github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/qmuntal/gltf"
)

type WeldOptions struct {
	Overwrite bool
}

func Weld(options WeldOptions) func(graph *Graph) error {
	return func(graph *Graph) error {
		return nil
	}
}



var WELD_DEFAULTS = WeldOptions{
	Overwrite: false,
}

func weldPrimitive(prim *GraphPrimitive, options WeldOptions) {
	// 如果有索引同时不想要重写
	if prim.Indices != nil && !options.Overwrite {
		return
	}
	// 如果是点，那就不需要weld
	if prim.Mode == gltf.PrimitivePoints {
		return
	}
	position, err := prim.ReadPostion()
	if err != nil {
		return
	}
	srcVertexCount := len(position)
	// 不一定有 indices ,没有的时候
	srcIndices, err := prim.ReadIndices()
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

	for i := 0; i < srcIndicesCount; i++ {
		var srcIndex uint32
		if srcIndices != nil {
			srcIndex = srcIndices[i]
		} else {
			srcIndex = uint32(i)
		}
		if writeMap[srcIndex] != EMPTY_U32 {
			continue
		}

		hashIndex := hashLookup(table, tableSize, prim, srcIndex, EMPTY_U32)
		dstIndex := table[hashIndex]

		if dstIndex == EMPTY_U32 {
			table[hashIndex] = srcIndex
			writeMap[srcIndex] = dstVertexCount
			dstVertexCount++
		} else {
			writeMap[srcIndex] = writeMap[dstIndex]
		}
	}

	compactPrimitive(prim, writeMap, dstVertexCount)
}
