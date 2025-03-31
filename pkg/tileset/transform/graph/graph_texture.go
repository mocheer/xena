package graph

import (
	"github.com/qmuntal/gltf"
)

// GraphTexture
// 使用 TexCoord 来定义使用哪种纹理坐标，0代表TEXCOORD_0，1代表TEXCOORD_1
type GraphTexture struct {
	Graph *Graph
	*gltf.Texture
}

func (m GraphTexture) GetImage() *GraphImage {
	return m.Graph.Images[*m.Source]
}

func (m GraphTexture) GetSampler() *GraphSampler {
	return m.Graph.Samplers[*m.Sampler]
}

func (m *GraphTexture) IsUsed() bool {
	m.Graph.FindUsedNodes(func(node *GraphNode) bool {
		return node.GetMesh().FindMaterial(func(material *GraphMaterial) bool {
			return m == material.GetBaseColorTexture() || m == material.GetNormalTexture() || m == material.GetEmissiveTexture() || m == material.GetOcclusionTexture()
		})
	})
	return false
}

func (m *GraphTexture) Dispose() {

}

func (m *GraphTexture) IsEmptyExtras() bool {
	return m.Extras == nil
}
