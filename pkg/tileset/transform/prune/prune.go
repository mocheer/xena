package prune

import (
	"slices"

	. "github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/qmuntal/gltf"
)

const (
	EPS float64 = 3.0 / 255.0
)

type PruneOptions struct {
	PropertyTypes     []string
	KeepLeaves        bool
	KeepAttributes    bool
	KeepIndices       bool
	KeepSolidTextures bool
	KeepExtras        bool
}

var PRUNE_DEFAULTS = PruneOptions{
	PropertyTypes: []string{
		PropertyTypeNode,
		PropertyTypeSkin,
		PropertyTypeMesh,
		PropertyTypeCamera,
		PropertyTypePrimitive,
		PropertyTypeAnimation,
		PropertyTypeMaterial,
		PropertyTypeTexture,
		PropertyTypeAccessor,
		PropertyTypeBuffer,
	},
	KeepLeaves:        false,
	KeepAttributes:    false,
	KeepIndices:       false,
	KeepSolidTextures: false, //是否保留单色纹理，有的人用一张纯色图片来渲染纹理，完全没有必要
	KeepExtras:        false, //是否保留用户数据
}

func Prune(options PruneOptions) func(graph *Graph) error {
	// 这里理应返回Counter
	return func(graph *Graph) error {
		// 只保留单场景
		if !options.KeepLeaves {
			graph.DiscardScenes()
		}
		// 销毁没有数据的mesh
		if slices.Contains(options.PropertyTypes, PropertyTypeMesh) {
			for _, mesh := range graph.Meshes {
				if len(mesh.Primitives) == 0 {
					mesh.Dispose()
				}
			}
		}
		// 销毁没有用到的或者内容为空的node
		if slices.Contains(options.PropertyTypes, PropertyTypeNode) {
			for _, node := range graph.Nodes {
				nodeTreeShake(node, options.KeepExtras)
			}
		}
		// 销毁蒙皮
		if slices.Contains(options.PropertyTypes, PropertyTypeSkin) {
			for _, skin := range graph.Skins {
				treeShake(skin, options.KeepExtras)
			}
		}
		// 销毁没有用到的mesh，上面只是销毁内容为空的mesh，这里是销毁没有用的
		if slices.Contains(options.PropertyTypes, PropertyTypeMesh) {
			for _, mesh := range graph.Meshes {
				treeShake(mesh, options.KeepExtras)
			}
		}
		// 销毁没有用到的camera
		if slices.Contains(options.PropertyTypes, PropertyTypeCamera) {
			for _, camera := range graph.Cameras {
				treeShake(camera, options.KeepExtras)
			}
		}

		// Prune 用于去掉用不到的attribute属性
		// 有些模型会有各种私有的attribute属性，或者对应的数据已经删除，但这里的属性仍然存在
		// 尤其是一堆纹理数据没有删除，这里需要根据用到的纹理，重新排列，保证连续，所有数据都有被用到
		if !options.KeepAttributes && slices.Contains(options.PropertyTypes, PropertyTypeAccessor) {

		}

		// Prune 丢弃索引，直接用顶点数据
		// 有的三维模型是直接用顶点数据的顺序渲染的，但却又专门生成顶点索引
		if !options.KeepIndices && slices.Contains(options.PropertyTypes, PropertyTypeAccessor) {
			for _, mesh := range graph.Meshes {
				for _, prim := range mesh.GetPrimitives() {
					prim.DiscardIndices()
				}
			}
		}
		// 去掉动画
		if slices.Contains(options.PropertyTypes, PropertyTypeAnimation) {
			graph.Animations = nil
			// 这里还需要去掉PrimitiveTarget的目标位置数据
		}
		// 销毁没有用到的材质
		if slices.Contains(options.PropertyTypes, PropertyTypeMaterial) {
			for _, material := range graph.Materials {
				treeShake(material, options.KeepExtras)
			}
		}
		// 销毁没有用到的纹理
		if slices.Contains(options.PropertyTypes, PropertyTypeTexture) {
			for _, texture := range graph.Textures {
				treeShake(texture, options.KeepExtras)
			}
			if !options.KeepSolidTextures {
				if err := pruneSolidTextures(graph); err != nil {
					return err
				}
			}
		}

		//
		if slices.Contains(options.PropertyTypes, PropertyTypeAccessor) {
			for _, accessor := range graph.Accessors {
				treeShake(accessor, options.KeepExtras)
			}
		}

		if slices.Contains(options.PropertyTypes, PropertyTypeBuffer) {
			for _, buffer := range graph.Buffers {
				treeShake(buffer, options.KeepExtras)
			}
		}

		return nil
	}
}

func treeShake(target GraphBase, keepExtras bool) {
	needsExtras := keepExtras && !target.IsEmptyExtras()
	if target.IsUsed() && !needsExtras {
		target.Dispose()
	}
}

func nodeTreeShake(node *GraphNode, keepExtras bool) {
	isEmpty := len(node.Children) == 0 && node.Mesh == nil
	needsExtras := keepExtras && !node.IsEmptyExtras()
	// 销毁空内容或者没有被用到的节点
	if (isEmpty && !needsExtras) || !node.IsUsed() {
		node.Dispose()
	}
}

func pruneAttributes(prim *gltf.Primitive, unused []string) {
	for _, semantic := range unused {
		delete(prim.Attributes, semantic)
	}
}

func pruneSolidTextures(graph *Graph) error {
	return nil
}

// // applyMaterialFactor
// // 在js中很多属性是可以直接在material访问的，这里很多属性都是要到PBRMetallicRoughness才能访问，这里要再仔细查看gltf标准文档
// func applyMaterialFactor(material *gltf.Material, factor vec4.Vec4, slot string) bool {
// 	switch slot {
// 	case "baseColorTexture":
// 		baseColorFacotr := vec4.Multiply(&vec4.Vec4{}, factor, vec4.Vec4(*material.PBRMetallicRoughness.BaseColorFactor))
// 		material.PBRMetallicRoughness.BaseColorFactor = (*[4]float64)(&baseColorFacotr)
// 		return true
// 	case "emissiveTexture":
// 		material.EmissiveFactor = vec3.Multiply(vec3.Vec3{factor[0], factor[1], factor[2]}, material.EmissiveFactor)
// 		return true
// 	case "occlusionTexture":
// 		return math.Abs(float64(factor[0]-1)) <= EPS
// 	case "metallicRoughnessTexture":
// 		*material.PBRMetallicRoughness.RoughnessFactor *= factor[1]
// 		*material.PBRMetallicRoughness.MetallicFactor *= factor[2]
// 		return true
// 	case "normalTexture":
// 		return vec4.Length(vec4.Subtract(&vec4.Vec4{}, factor, vec4.Vec4{0.5, 0.5, 1, 1})) <= EPS
// 	}
// 	return false
// }

// func getTextureFactor(texture *GraphTexture) (vec4.Vec4, error) {
// 	//  读取纹理对应的图片资源
// 	bytes, _ := texture.GetImage().ReadBufferView()
// 	image, _, err := img.FromBytes(bytes)
// 	if err != nil {
// 		return vec4.Vec4{}, err
// 	}
// 	pixels := image.Image
// 	// rgba
// 	r, g, b, a := pixels.At(0, 0).RGBA()
// 	min := vec4.Vec4{float64(r), float64(g), float64(b), float64(a)}
// 	max := vec4.Vec4{float64(r), float64(g), float64(b), float64(a)}
// 	target := vec4.Vec4{}

// 	width, height := pixels.Bounds().Dx(), pixels.Bounds().Dy()

// 	for i := 0; i < width; i++ {
// 		for j := 0; j < height; j++ {
// 			r, g, b, a := pixels.At(i, j).RGBA()
// 			min[0] = math.Min(min[0], float64(r))
// 			min[1] = math.Min(min[1], float64(g))
// 			min[2] = math.Min(min[2], float64(b))
// 			min[3] = math.Min(min[3], float64(a))
// 			//
// 			max[0] = math.Max(max[0], float64(r))
// 			max[1] = math.Max(max[1], float64(g))
// 			max[2] = math.Max(max[2], float64(b))
// 			max[3] = math.Max(max[3], float64(a))
// 		}
// 		if vec4.Length(vec4.Subtract(&target, max, min))/255 > EPS {
// 			return vec4.Vec4{}, nil
// 		}
// 	}
// 	return vec4.Scale(&target, vec4.Add(&target, max, min), 0.5/255), nil
// }
