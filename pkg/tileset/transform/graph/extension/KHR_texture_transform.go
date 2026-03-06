package extension

import "github.com/qmuntal/gltf"

// KHR_texture_transform
// @see https://github.com/KhronosGroup/glTF/blob/main/extensions/2.0/Khronos/KHR_texture_transform/README.md
// KHR_mesh_quantization 依赖此拓展
const KHR_texture_transform = "KHR_texture_transform"

// KHR_texture_transform_options
// offset   纹理坐标的平移量 [U, V]，范围 [0, 1]，例如 [0.5, 0] 表示向右平移 50%，默认	[0, 0]
// rotation 纹理绕中心点（[0.5, 0.5]）的旋转角度（弧度制），正值逆时针旋转，默认为 0
// scale 纹理缩放比例 [scaleU, scaleV]，例如 [2, 2] 表示放大 2 倍，默认[1, 1]
// texCoord 指定使用的纹理坐标集索引（对应 mesh 的 TEXCOORD_0、TEXCOORD_1 等），通常为0,对应TEXCOORD_0
// 以上全部是非必需的参数
type KHR_texture_transform_options struct {
	Offset   []float32
	Rotation float64
	Scale    []float32
	TexCoord int
}

// BindToTextureInfo
func (m KHR_texture_transform_options) BindToTextureInfo(textInfo *gltf.TextureInfo) {
	if textInfo.Extensions == nil {
		textInfo.Extensions = make(map[string]any)
	}
	m.BindTo(textInfo.Extensions)
}

// BindToNormalTexture
func (m KHR_texture_transform_options) BindToNormalTexture(textInfo *gltf.NormalTexture) {
	if textInfo.Extensions == nil {
		textInfo.Extensions = make(map[string]any)
	}
	m.BindTo(textInfo.Extensions)
}

// BindToOcclusionTexture
func (m KHR_texture_transform_options) BindToOcclusionTexture(textInfo *gltf.OcclusionTexture) {
	if textInfo.Extensions == nil {
		textInfo.Extensions = make(map[string]any)
	}
	m.BindTo(textInfo.Extensions)
}

// BindTo
// extensions 是 material 的 extensions 属性
// 这里用于设置拓展的参数，还要注意在根节点声明该拓展
func (m KHR_texture_transform_options) BindTo(extensions map[string]any) {
	extensions[KHR_texture_transform] = map[string]any{
		"offset": m.Offset,
		"scale":  m.Scale,
	}
}
