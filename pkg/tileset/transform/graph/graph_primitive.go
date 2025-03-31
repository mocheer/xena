package graph

import (
	"errors"
	"iter"
	"maps"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

// GraphPrimitive
// Primitive 用于定义网格的静态几何数据，包括顶点、索引、材质等。
// PrimitiveTarget 用于定义网格的动态几何数据，主要用于形变动画。通过在动画中改变目标的权重，可以实现网格的平滑变形。
// PrimitiveTarget 是Primitive的一种，有些数据模型专门区分Primitive和PrimitiveTarget
// 每个Primitive对象都有一个targets字段，它是一个数组。如果这个数组不为空，说明该Primitive与PrimitiveTarget相关联
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

// GetAttributes
func (m GraphPrimitive) GetAttributes() []*GraphAccessor {
	accessors := []*GraphAccessor{}
	for _, s := range m.Attributes {
		accessors = append(accessors, m.Graph.Accessors[s])
	}
	return accessors
}

// GetAttribute
func (m GraphPrimitive) GetAttribute(semantic string) *GraphAccessor {
	index, ok := m.Attributes[semantic]
	if ok {
		return m.Graph.Accessors[index]
	}
	return nil
}

func (m GraphPrimitive) GetMaterial() *GraphMaterial {
	return m.Graph.Materials[*m.Material]
}

// ReadIndices 读取顶点索引数据
// 顶点索引数据可为空，为空时，直接按照顶点数据绘制
func (m GraphPrimitive) ReadIndices() ([]uint32, error) {
	doc := m.Graph.Doc
	indices, err := modeler.ReadIndices(doc, doc.Accessors[*m.Indices], nil)
	return indices, err
}

// WriteIndices 这里写入，但可能需要销毁原来的数据
func (m GraphPrimitive) WriteIndices(indices []uint32) {
	doc := m.Graph.Doc
	index := gltf.Index(modeler.WriteIndices(doc, indices))
	m.Indices = index
}

// PruneIndices
// 这里只是去掉索引数据，还需要保证这个索引顺序其他Primtive都没有使用，才能修改buffer和bufferView等
func (m GraphPrimitive) DiscardIndices() {
	m.Indices = nil
}

// ReadPostion 读取顶点数据，一般需要配合索引数据来构建模型
func (m GraphPrimitive) ReadPostion() ([][3]float32, error) {
	val := m.GetAttribute(gltf.POSITION)
	if val != nil {
		return val.ReadAsPosition()
	}
	return nil, errors.New("not found")
}

// ReadTEXCOORD_0
// TEXCOORD_0  通常用于主纹理映射（Primary Texture Mapping）。它是最常用的纹理坐标集，用于将基本纹理（如漫反射纹理、颜色纹理等）映射到模型表面。
func (m GraphPrimitive) ReadTEXCOORD_0() ([][2]float32, error) {
	val := m.GetAttribute(gltf.TEXCOORD_0)
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
	val := m.GetAttribute(gltf.TEXCOORD_1)
	if val != nil {
		return val.ReadAsTextureCoord()
	}

	return nil, errors.New("not found")
}
func (m GraphPrimitive) WritePostion(position [][3]float32) {
	m.Attributes[gltf.POSITION] = modeler.WritePosition(m.Graph.Doc, position)
}

func (m GraphPrimitive) WriteTEXCOORD_0(textureCoord [][2]float32) {
	doc := m.Graph.Doc
	m.Attributes[gltf.TEXCOORD_0] = modeler.WriteTextureCoord(doc, textureCoord)
}

func (m GraphPrimitive) WriteTEXCOORD_1(textureCoord [][2]float32) {
	m.Attributes[gltf.TEXCOORD_1] = modeler.WriteTextureCoord(m.Graph.Doc, textureCoord)
}
