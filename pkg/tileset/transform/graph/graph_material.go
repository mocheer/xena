package graph

import (
	"github.com/qmuntal/gltf"
)

// GraphMaterial
// 当材质有纹理数据，但Primitive没有纹理坐标数据（TEXCOORD），有些渲染器会直接忽略纹理直接使用默认颜色，比如Three.js
// 但有些渲染器会直接使用默认纹理坐标，或将顶点的 X 和 Y 坐标直接用作纹理坐标，或作做一个平面投影再作为纹理坐标，比如Babylon.js
type GraphMaterial struct {
	Graph *Graph
	*gltf.Material
}

// GetPBRMetallicRoughness PBRMetallicRoughness是默认使用的材质模型,一般渲染图片都是通过 BaseColorTexture 来实现
func (m GraphMaterial) GetBaseColorTexture() *GraphTexture {
	return m.Graph.Textures[m.PBRMetallicRoughness.BaseColorTexture.Index]
}

func (m GraphMaterial) GetMetallicRoughnessTexture() *GraphTexture {
	return m.Graph.Textures[m.PBRMetallicRoughness.MetallicRoughnessTexture.Index]
}

// GetNormalTexture
func (m GraphMaterial) GetNormalTexture() *GraphTexture {
	return m.Graph.Textures[*m.NormalTexture.Index]
}

// GetOcclusionTexture
func (m GraphMaterial) GetOcclusionTexture() *GraphTexture {
	return m.Graph.Textures[*m.OcclusionTexture.Index]
}

// GetEmissiveTexture
func (m GraphMaterial) GetEmissiveTexture() *GraphTexture {
	return m.Graph.Textures[*&m.EmissiveTexture.Index]
}

func (m *GraphMaterial) IsUsed() bool {
	return false
}

func (m *GraphMaterial) Dispose() {

}

func (m *GraphMaterial) IsEmptyExtras() bool {
	return m.Extras == nil
}
