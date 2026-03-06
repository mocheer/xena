package graph

import (
	"errors"
	"iter"
	"maps"

	"github.com/qmuntal/gltf"
)

// GraphPrimitive
// Primitive 用于定义网格的静态几何数据，包括顶点、索引、材质等。
// PrimitiveTarget 用于定义网格的动态几何数据，主要用于形变动画。通过在动画中改变目标的权重，可以实现网格的平滑变形。
// PrimitiveTarget 是Primitive的一种，有些数据模型专门区分Primitive和PrimitiveTarget
// 每个Primitive对象都有一个targets字段，它是一个数组。如果这个数组不为空，说明该Primitive与PrimitiveTarget相关联

// Unity的一个游戏资源（江湖），导出的资源包含 COLOR_0 会导致纹理渲染异常
// 一是：在着色器中，将顶点颜色与从纹理采样得到的颜色相乘 (textureColor * vertexColor)。这样可以用顶点颜色来对纹理进行染色或变暗。
// 二是：直接用顶点颜色作为最终输出

// const (
//
//	POSITION   = "POSITION"
//	NORMAL     = "NORMAL"
//	TANGENT    = "TANGENT"
//	TEXCOORD_0 = "TEXCOORD_0"
//	TEXCOORD_1 = "TEXCOORD_1"
//	WEIGHTS_0  = "WEIGHTS_0"
//	JOINTS_0   = "JOINTS_0"
//	COLOR_0    = "COLOR_0"
//
// )
type GraphPrimitive struct {
	*gltf.Primitive
	Graph *Graph
}

/**
 * 列出与该图元（primitive）相关的所有顶点属性语义，不包括用于形变目标（morph targets）的语义。
 * 例如，`['POSITION', 'NORMAL', 'TEXCOORD_0']`。
 */
func (m GraphPrimitive) GetSemantics() iter.Seq[string] {
	return maps.Keys(m.Attributes)
}

// GetAttributeAccessors
func (m GraphPrimitive) GetAttributeAccessors() []*GraphAccessor {
	accessors := []*GraphAccessor{}
	for _, s := range m.Attributes {
		accessors = append(accessors, m.Graph.GetAccessor(s))
	}
	return accessors
}

// GetAttributeAccessor
func (m GraphPrimitive) GetAttributeAccessor(semantic string) *GraphAccessor {
	index, ok := m.Attributes[semantic]
	if ok {
		return m.Graph.GetAccessor(index)
	}
	return nil
}

// GetMaterial
func (m GraphPrimitive) GetMaterial() *GraphMaterial {
	if m.Material != nil {
		return m.Graph.GetMaterial(*m.Material)
	}
	return nil
}

// GetIndices
func (m GraphPrimitive) GetIndices() *GraphAccessor {
	if m.Indices != nil {
		return m.Graph.GetAccessor(*m.Indices)
	}
	return nil
}

// ReadIndices 读取顶点索引数据
// 顶点索引数据可为空，为空时，直接按照顶点数据绘制
func (m GraphPrimitive) ReadIndices() ([]uint32, error) {
	acc := m.GetIndices()
	if acc != nil {
		return acc.ReadAsIndices()
	}
	return nil, errors.New("not found")
}

// WriteIndices 这里写入，但可能需要销毁原来的数据
func (m GraphPrimitive) WriteIndices(indices []uint32) {
	*m.Indices = m.Graph.WriteIndices(indices)
}

// DiscardIndices
// 这里只是去掉索引数据，还需要保证这个索引顺序其他Primtive都没有使用，才能修改buffer和bufferView等
func (m GraphPrimitive) DiscardIndices() {
	m.Indices = nil
}

// ReadPostion 读取顶点数据，一般需要配合索引数据来构建模型
func (m GraphPrimitive) ReadPostion() ([][3]float32, error) {
	val := m.GetAttributeAccessor(gltf.POSITION)
	if val != nil {
		position, err := val.ReadAsPosition()
		if err != nil {
			return nil, err
		}
		return position, err
	}
	return nil, errors.New("not found")
}

// ReadTEXCOORD_0
// 常说的uv坐标
// TEXCOORD_0  通常用于主纹理映射（Primary Texture Mapping）。它是最常用的纹理坐标集，用于将基本纹理（如漫反射纹理、颜色纹理等）映射到模型表面。
func (m GraphPrimitive) ReadTEXCOORD_0() ([][2]float32, error) {
	val := m.GetAttributeAccessor(gltf.TEXCOORD_0) //不一定存在
	if val != nil {
		return val.ReadAsTextureCoord()
	}
	return nil, errors.New("not found")
}

// ReadTEXCOORD_1
// TEXCOORD_1
// 通常用于辅助纹理映射（Secondary Texture Mapping）。它可以用于多纹理技术，例如：
// 法线贴图（Normal Mapping）：用于存储法线纹理的坐标。
// 环境光遮蔽（AO）贴图：用于存储环境光遮蔽纹理的坐标。
// 其他辅助纹理：如高光贴图、遮罩贴图等。
func (m GraphPrimitive) ReadTEXCOORD_1() ([][2]float32, error) {
	val := m.GetAttributeAccessor(gltf.TEXCOORD_1)
	if val != nil {
		return val.ReadAsTextureCoord()
	}

	return nil, errors.New("not found")
}

// WritePostion
func (m GraphPrimitive) WritePostion(position [][3]float32) {
	m.Attributes[gltf.POSITION] = m.Graph.WritePostion(position)
}

// WriteTEXCOORD_0
func (m GraphPrimitive) WriteTEXCOORD_0(textureCoord [][2]float32) {
	m.Attributes[gltf.TEXCOORD_0] = m.Graph.WriteTextureCoord(textureCoord)
}

// WriteTEXCOORD_1
func (m GraphPrimitive) WriteTEXCOORD_1(textureCoord [][2]float32) {
	m.Attributes[gltf.TEXCOORD_1] = m.Graph.WriteTextureCoord(textureCoord)
}

// https://github.com/donmccurdy/glTF-Transform/blob/main/packages/functions/src/get-vertex-count.ts#L180

// GetVertexCount
// 获取顶点数量
// 类似 VertexCountMethod.UPLOAD,VertexCountMethod.UPLOAD_NAIVE
func (m GraphPrimitive) GetVertexCount() uint32 {
	val := m.GetAttributeAccessor(gltf.POSITION)
	if val != nil {
		return val.Count
	}
	return 0
}

// GetVertexCountWithRender
// 获取实际渲染的顶点数量
func (m GraphPrimitive) GetVertexCountWithRender() uint32 {
	val := m.GetIndices()
	if val != nil {
		return val.Count
	}
	return m.GetVertexCount()
}

// GetVertexCountWithRender
// 获取没有用到的顶点数量，目前的计算方式不准确，计算方式为：顶点数-索引数，其中索引值需唯一
// 这里的索引可能重复使用某个顶点
func (m GraphPrimitive) GetVertexCountWithUnused() uint32 {
	return m.GetVertexCount() - m.GetVertexCountWithRender()
	// return 0
}

// DisposeAttributeColor0
func (m *GraphPrimitive) DisposeAttributeColor0() {
	// TODO 移除对应的buffer信息
	delete(m.Attributes, gltf.COLOR_0)
}

// Valid
// 案例：桂林.glb 模型（由blender导出），存在Primitive，但Primitve下的attribute没有 TEXCOORD_0 字段，这是不标准的，但这种情况常常存在。
// Cesium错误：
// ERROR: 0:326: 'v_texCoord_0' : undeclared identifier
// ERROR: 0:326: '=' : dimension mismatch
// ERROR: 0:326: '=' : cannot convert from 'const highp float' to 'highp 2-component vector of float'
func (m GraphPrimitive) Valid() error {
	// 只有 POSITION 和 TEXCOORD_0 是必须的
	// 有的模型可能没有TEXCOORD_0，但有TEXCOORD_1，只是这样做不太符合标准
	_, ok := m.Attributes[gltf.POSITION]
	if !ok {
		return ErrPrimitiveUndefinedPosition
	}
	_, ok = m.Attributes[gltf.TEXCOORD_0]
	if !ok {
		return ErrPrimitvieUndefinedTexture0
	}
	// 但可选的也要存在Accessor，有的会指向错误的索引
	for _, accessor := range m.GetAttributeAccessors() {
		if accessor.Accessor == nil {
			return errors.New("primitive定义的属性对应的Accessor不存在")
		}
	}
	return nil
}
