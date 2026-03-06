package graph

import (
	"github.com/qmuntal/gltf"
)

// GraphTextureInfo
// BaseColorTexture、MetallicRoughnessTexture、EmissiveTexture是 TextureInfo 类型
// 使用 TexCoord 来定义使用哪种纹理坐标，0代表TEXCOORD_0，1代表TEXCOORD_1
type GraphTextureInfo struct {
	Graph *Graph
	*gltf.TextureInfo
}

// IsUsed
func (m *GraphTextureInfo) IsUsed() bool {
	return true
}

// Dispose
func (m *GraphTextureInfo) Dispose() {

}

// IsEmptyExtras
func (m *GraphTextureInfo) IsEmptyExtras() bool {
	return m.Extras == nil
}
