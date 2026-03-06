package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

// GraphMaterial
// 当材质有纹理数据，但Primitive没有纹理坐标数据（TEXCOORD），有些渲染器会直接忽略纹理直接使用默认颜色，比如Three.js
// 但有些渲染器会直接使用默认纹理坐标，或将顶点的 X 和 Y 坐标直接用作纹理坐标，或作做一个平面投影再作为纹理坐标，比如Babylon.js
type GraphMaterial struct {
	Graph *Graph
	*gltf.Material
}

// Index
func (m *GraphMaterial) Index() int {
	return slices.Index(m.Graph.Materials, m.Material)
}

// IsUsed
func (m *GraphMaterial) IsUsed() bool {
	return true
}

// Dispose
func (m *GraphMaterial) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Materials = slices.Delete(m.Graph.Materials, index, index+1)
	}
}

// IsEmptyExtras
func (m *GraphMaterial) IsEmptyExtras() bool {
	return m.Extras == nil
}

// GetPBRMetallicRoughness PBRMetallicRoughness是默认使用的材质模型,一般渲染图片都是通过 BaseColorTexture 来实现
// 最终基础颜色 = baseColorFactor × baseColorTexture × COLOR_0 (顶点色)
// 有不少游戏资源用顶点色为纯黑vec4(0, 0, 0, 1)或者完全透明来混淆
// 默认为sRGB编码?
func (m GraphMaterial) GetBaseColorTexture() *GraphTexture {
	if m.PBRMetallicRoughness == nil || m.PBRMetallicRoughness.BaseColorTexture == nil {
		return nil
	}
	return m.Graph.GetTexture(m.PBRMetallicRoughness.BaseColorTexture.Index)
}

// GetBaseColorFactor
// 若存在 baseColorTexture（基础色贴图），BaseColorFactor会与之相乘，作为最终颜色的缩放因子
// 如果不存在，基础色完全由 BaseColorFactor 决定，省去贴图加载（节省显存）和纹理数据传输（降低带宽），尤其利好移动端或大场景渲染。
// 金属材质：BaseColor直接影响镜面反射颜色。
// 非金属材质：BaseColor仅影响漫反射，镜面反射由菲涅尔效应决定
// BaseColorFactor的值需在线性色彩空间（非sRGB）中定义。若输入为sRGB值（如常见的8位RGB），需先转换为线性值，
// sRGB 和 Linear 的转换，
func (m GraphMaterial) GetBaseColorFactor() *[4]float64 {
	if m.PBRMetallicRoughness == nil {
		return nil
	}
	return m.PBRMetallicRoughness.BaseColorFactor
}

// SetBaseColorFactor
// 设置材质的基础色，value的格式为rgba四维向量
func (m GraphMaterial) SetBaseColorFactor(value [4]float64) {
	*m.PBRMetallicRoughness.BaseColorFactor = value
}

// GetMetallicRoughnessTexture
// 默认为Linear编码
func (m GraphMaterial) GetMetallicRoughnessTexture() *GraphTexture {
	if m.PBRMetallicRoughness == nil || m.PBRMetallicRoughness.MetallicRoughnessTexture == nil {
		return nil
	}
	return m.Graph.GetTexture(m.PBRMetallicRoughness.MetallicRoughnessTexture.Index)
}

// GetMetallicFactor
// metallicFactor 控制材质金属感（0.0=非金属，1.0=金属）。
func (m GraphMaterial) GetMetallicFactor() *float64 {
	if m.PBRMetallicRoughness == nil {
		return nil
	}
	return m.PBRMetallicRoughness.MetallicFactor
}

// GetNormalTexture
// 法线纹理，控制凹凸感
// 默认为Linear编码
func (m GraphMaterial) GetNormalTexture() *GraphTexture {
	if m.NormalTexture == nil || m.NormalTexture.Index == nil {
		return nil
	}
	return m.Graph.GetTexture(*m.NormalTexture.Index)
}

// GetOcclusionTexture
func (m GraphMaterial) GetOcclusionTexture() *GraphTexture {
	if m.OcclusionTexture == nil || m.OcclusionTexture.Index == nil {
		return nil
	}
	return m.Graph.GetTexture(*m.OcclusionTexture.Index)
}

// GetEmissiveTexture
// （自发光纹理）是一个用于模拟物体表面自主发光效果的贴图。它能让模型的一部分（如屏幕、灯管、信号灯）看起来像是在发光，且其亮度不会受场景灯光影响
func (m GraphMaterial) GetEmissiveTexture() *GraphTexture {
	if m.EmissiveTexture == nil {
		return nil
	}
	return m.Graph.GetTexture(m.EmissiveTexture.Index)
}
