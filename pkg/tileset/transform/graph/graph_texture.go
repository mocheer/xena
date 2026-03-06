package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

// GraphTexture
// 使用 TexCoord 来定义使用哪种纹理坐标，0代表TEXCOORD_0，1代表TEXCOORD_1
type GraphTexture struct {
	Graph *Graph
	*gltf.Texture
}

// GetImage
func (m GraphTexture) GetImage() *GraphImage {
	return m.Graph.GetImage(*m.Source)
}

// GetSampler
func (m GraphTexture) GetSampler() *GraphSampler {
	return m.Graph.GetSampler(*m.Sampler)
}

// IsUsed
func (m *GraphTexture) IsUsed() bool {
	return m.Graph.FindUsedNodes(func(node *GraphNode) bool {
		return node.GetMesh().FindMaterial(func(material *GraphMaterial) bool {
			return m.IsUsedWithMaterial(material)
		})
	})
}

// Index
func (m *GraphTexture) Index() int {
	return slices.Index(m.Graph.Textures, m.Texture)
}

// IsUsedWithMaterial
func (m *GraphTexture) IsUsedWithMaterial(material *GraphMaterial) bool {
	index := uint32(m.Index())
	return (material.PBRMetallicRoughness != nil && material.PBRMetallicRoughness.BaseColorTexture != nil && index == material.PBRMetallicRoughness.BaseColorTexture.Index) ||
		(material.NormalTexture != nil && index == *material.NormalTexture.Index) ||
		(material.EmissiveTexture != nil && index == material.EmissiveTexture.Index) ||
		(material.OcclusionTexture != nil && index == *material.OcclusionTexture.Index)
}

// GetLinkMaterial
// 获取使用到当前纹理的材质对象
func (m *GraphTexture) GetParents() []*GraphMaterial {
	parents := []*GraphMaterial{}
	materials := m.Graph.GetMaterials()
	for _, material := range materials {
		if m.IsUsedWithMaterial(material) {
			parents = append(parents, material)
		}
	}
	return parents
}

// Dispose
func (m *GraphTexture) Dispose() {
	index := slices.Index(m.Graph.Textures, m.Texture)
	if index != -1 {
		m.Graph.Textures = slices.Delete(m.Graph.Textures, index, index+1)
		//
		targetIndex := uint32(index)
		m.Graph.EachUsedNodes(func(node *GraphNode) {
			node.GetMesh().EachMaterial(func(material *GraphMaterial) {
				if material.PBRMetallicRoughness.BaseColorTexture.Index > targetIndex {
					material.PBRMetallicRoughness.BaseColorTexture.Index--
				}
				if *material.NormalTexture.Index > targetIndex {
					*material.NormalTexture.Index--
				}
			})
		})
	}

}

func (m *GraphTexture) IsEmptyExtras() bool {
	return m.Extras == nil
}
