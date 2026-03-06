package weld

import (
	"errors"
	"math"

	. "github.com/mocheer/xena/pkg/tileset/transform/graph"
)

const (
	EMPTY_U32 = uint32(math.MaxUint32)
)

// VertexStream 用于高效处理顶点数据的哈希和比较。
type VertexStream struct {
	attributes []VertexStreamAttribute
	u8         []byte   // 临时缓冲区（4字节对齐的顶点数据）
	u32        []uint32 // 与 u8 共享内存，用于哈希计算
}

type VertexStreamAttribute struct {
	u8               []byte // 顶点属性数据（原始字节）
	byteStride       int    // 每个顶点的实际字节跨度
	paddedByteStride int    // 填充后的字节跨度（4字节对齐）
}

// NewVertexStream 创建一个新的 VertexStream，初始化所有顶点属性。
// 将顶点的所有属性（位置、法线、纹理坐标等）组合起来，计算一个唯一的哈希值。
func NewVertexStream(prim *GraphPrimitive) *VertexStream {
	vs := &VertexStream{}
	totalStride := 0

	// 遍历图元的所有顶点属性（例如位置、法线、UV等）
	for _, accessor := range prim.GetAttributeAccessors() {
		totalStride += vs.initAttribute(accessor)
	}

	// 初始化对齐的临时缓冲区
	vs.u8 = make([]byte, totalStride)
	vs.u32 = make([]uint32, totalStride/4) // 4 bytes per uint32
	return vs
}

// initAttribute 初始化单个顶点属性，返回填充后的字节跨度。
func (vs *VertexStream) initAttribute(accessor *GraphAccessor) int {
	elementSize := int(accessor.Type.Components())          // 例如 vec3=3
	componentSize := int(accessor.ComponentType.ByteSize()) // 例如 FLOAT=4
	byteStride := elementSize * componentSize
	paddedByteStride := padNumber(byteStride, 4) // 4字节对齐

	// 获取访问器的底层字节数据，这里只读取了BufferView,没有读取稀疏数据
	data, err := accessor.ReadBufferView()
	if err != nil {
		panic(err)
	}

	vs.attributes = append(vs.attributes, VertexStreamAttribute{
		u8:               data,
		byteStride:       byteStride,
		paddedByteStride: paddedByteStride,
	})

	return paddedByteStride
}

// Hash 计算指定索引顶点的哈希值。
func (vs *VertexStream) Hash(index uint32) uint32 {
	offset := 0
	// 将顶点数据复制到对齐的临时缓冲区
	for _, attr := range vs.attributes {
		for i := range attr.paddedByteStride {
			if i < attr.byteStride {
				vs.u8[offset+i] = attr.u8[int(index)*attr.byteStride+i]
			} else {
				vs.u8[offset+i] = 0 // 填充0
			}
		}
		offset += attr.paddedByteStride
	}
	// 计算 MurmurHash2
	return murmurHash2(0, vs.u32)
}

// Equal 比较两个顶点是否完全相同。
func (vs *VertexStream) Equal(a, b int) bool {
	for _, attr := range vs.attributes {
		aOffset := a * attr.byteStride
		bOffset := b * attr.byteStride
		for j := range attr.byteStride {
			if attr.u8[aOffset+j] != attr.u8[bOffset+j] {
				return false
			}
		}
	}
	return true
}

// padNumber 计算填充到指定对齐的字节数。
func padNumber(n, align int) int {
	return (n + align - 1) & ^(align - 1)
}

// murmurHash2 实现 MurmurHash2 算法。
func murmurHash2(h uint32, key []uint32) uint32 {
	const m uint32 = 0x5bd1e995
	const r = 24

	for _, k := range key {
		k *= m
		k ^= k >> r
		k *= m

		h = h*m ^ k
	}
	return h
}

// hashLookup 在哈希表中查找或插入顶点索引。
func hashLookup(table []uint32, buckets int, stream *VertexStream, key uint32, empty uint32) (int, error) {
	hashmod := uint32(buckets - 1)
	hashval := stream.Hash(key)
	bucket := hashval & hashmod

	for probe := 0; probe <= buckets; probe++ {
		item := table[bucket]
		if item == empty || stream.Equal(int(item), int(key)) {
			return int(bucket), nil
		}
		// 线性探测解决哈希冲突
		bucket = (bucket + uint32(probe) + 1) & hashmod
	}
	return 0, errors.New("hash table full")
}
