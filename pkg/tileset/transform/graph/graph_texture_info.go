package graph

import (
	"github.com/qmuntal/gltf"
)

// GraphTexture
// Texture
// 使用 TexCoord 来定义使用哪种纹理坐标，0代表TEXCOORD_0，1代表TEXCOORD_1
type GraphTextureInfo struct {
	Graph *Graph
	*gltf.TextureInfo
}

func (m *GraphTextureInfo) IsUsed() bool {
	return false
}

func (m *GraphTextureInfo) Dispose() {

}

func (m *GraphTextureInfo) IsEmptyExtras() bool {
	return m.Extras == nil
}
