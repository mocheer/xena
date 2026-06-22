package texture_merge

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"

	"github.com/mocheer/pluto/pkg/ts/img"

	"github.com/mocheer/xena/pkg/tileset/transform/graph"
	"github.com/mocheer/xena/pkg/tileset/transform/graph/extension"

	"github.com/qmuntal/gltf"
)

// glTF默认的采样器是重复（Repeat）模式。
// UV大于1.0是为了实现纹理重复的效果。例如，如果将一个2x2的纹理贴在一个大平面上，我们可以将UV设置为[0,2]，这样纹理就会重复4次。
// uv > 1.0 时默认时采样同一纹理的其他部分，但合成大图后会采样其他区域的纹理
// 目前只支持不需要平铺的纹理

// 方案一：读取uv并将对应的纹理展开（文件和尺寸变大）
// 方案二：当平铺的uv坐标的三角面都没有跨越1.0那么可以直接归一, 这种做法当一个三角顶点的uv坐标跨度比较大的时候容易异常，如果都是 1.0-2.0 又或者3.0-4.0都将正常，但如果是2.0-5.0（一个三角面内部平铺三次）就会异常
// 方案三：重写着色器
// 方案三：只合并不需要平铺的纹理（png图片一般不会平铺），甚至做分组合并，比如说：场景地图、树木

type TextureMergeOptions struct {
	Format string
}

// TextureMerge
// 目前只做粗暴的合并，合并所有图片纹理且不考虑材质、采样器不同的情况，且默认纹理没有KHR_texture_transform拓展
func TextureMerge(options TextureMergeOptions) func(graph *graph.Graph) error {
	return func(graph *graph.Graph) error {
		images, err := graph.ReadImagesAsImages()
		if err != nil {
			return err
		}
		if len(images) > 1 {
			atlas, atlasOptions := createTextureAtlas(images)
			update(graph, atlas, atlasOptions)
		}
		return nil
	}
}

// 创建纹理集并返回一张大图和UV映射关系
// 矩形装箱问题：Rectangle Packing
// 算法有：首次适应递减高度算法（First Fit Decreasing Height, FFDH），最佳适应算法（Best Fit），以及一些更复杂的算法如MaxRects算法
// https://github.com/InfinityTools/go-binpack2d 参考C实现的MaxRectsBinPack的算法，年代较久
// https://github.com/depp/skelly64/tree/main/lib/rectpack 任天堂的算法
// https://github.com/lewisgibson/go-binpack MaxRects算法，空间利用率极高，但性能小于skyline
// https://github.com/pekim/skyline 用于打包2D矩形的天际线算法。 空间利用率只是中等
// @see ds_ase.FromImages
func createTextureAtlas(images []image.Image) (image.Image, []extension.KHR_texture_transform_options) {
	// 计算大图尺寸 (简单堆叠)
	totalWidth := 0
	maxHeight := 0
	for _, img := range images {
		bounds := img.Bounds()
		totalWidth += bounds.Dx()
		if bounds.Dy() > maxHeight {
			maxHeight = bounds.Dy()
		}
	}
	// totalWidth = int(totalWidth/256) * 256
	// totalWidth = 16384
	// maxHeight = 16384
	// 创建空白大图
	atlas := image.NewRGBA(image.Rect(0, 0, totalWidth, maxHeight))
	// // 填充颜色
	// for y := 0; y < maxHeight; y++ {
	// 	for x := 0; x < totalWidth; x++ {
	// 		atlas.Set(x, y, color.RGBA{0, 255, 0, 255})
	// 	}
	// }
	atlasOptions := make([]extension.KHR_texture_transform_options, len(images))

	// 合并图像并记录UV变换
	// 这里横向排列纹理图
	// TODO 横向+垂直排列
	xOffset := 0
	for i, img := range images {
		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		// 绘制到纹理集
		draw.Draw(atlas, image.Rect(xOffset, 0, xOffset+width, height),
			img, bounds.Min, draw.Src)

		// 计算UV变换矩阵
		scaleX := float32(width) / float32(totalWidth)
		scaleY := float32(height) / float32(maxHeight)
		offsetX := float32(xOffset) / float32(totalWidth)
		//
		atlasOptions[i] = extension.KHR_texture_transform_options{
			Scale:  []float32{scaleX, scaleY},
			Offset: []float32{offsetX, 0},
		}

		xOffset += width
	}
	return atlas, atlasOptions
}

// 创建纹理集并返回一张大图和UV映射关系
// 垂直合并
func createTextureAtlasVertical(images []image.Image) (image.Image, []extension.KHR_texture_transform_options) {
	// 计算大图尺寸 (简单堆叠)
	totalHeight := 0
	maxWidth := 0
	for _, img := range images {
		bounds := img.Bounds()
		totalHeight += bounds.Dy()
		if bounds.Dx() > maxWidth {
			maxWidth = bounds.Dx()
		}
	}
	// 创建空白大图
	atlas := image.NewRGBA(image.Rect(0, 0, maxWidth, totalHeight))
	atlasOptions := make([]extension.KHR_texture_transform_options, len(images))
	// 合并图像并记录UV变换
	// 这里垂直排列纹理图
	// TODO 横向+垂直排列
	yOffset := 0
	for i, img := range images {
		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()

		// 绘制到纹理集
		draw.Draw(atlas, image.Rect(0, yOffset, width, yOffset+height),
			img, bounds.Min, draw.Src)

		// 计算UV变换矩阵
		scaleX := float32(width) / float32(maxWidth)
		scaleY := float32(height) / float32(totalHeight)
		offsetY := float32(yOffset) / float32(totalHeight)
		//
		atlasOptions[i] = extension.KHR_texture_transform_options{
			Scale:  []float32{scaleX, scaleY},
			Offset: []float32{0, offsetY},
		}
		//
		yOffset += height
	}
	return atlas, atlasOptions
}

// 更新glTF文档
func update(graph *graph.Graph, image image.Image, atlasOptions []extension.KHR_texture_transform_options) {
	data, err := img.ToJpegBytes(image, &jpeg.Options{Quality: 75})
	// ds.Save("./.temp/test.jpeg", data)
	if err != nil {
		panic(err)
	}
	updateTextInfo := func(texInfo *gltf.TextureInfo) {
		// 添加扩展

		texture := graph.GetTexture(texInfo.Index)
		imageIndex := texture.Source
		options := atlasOptions[*imageIndex]
		// 应用 KHR_texture_transform 扩展

		options.BindToTextureInfo(texInfo)

		// 指向新纹理
		// texInfo.Index = 0
		texture.Source = gltf.Index(0)
	}

	// 更新所有材质
	for _, node := range graph.GetNodes() {
		for _, p := range node.GetMesh().GetPrimitives() {
			// uv, err := p.ReadTEXCOORD_0()
			// if err != nil {
			// 	continue
			// }
			// flag := false
			// for i, xy := range uv {
			// 	if xy[0] > 1 {
			// 		xy[0] = float32(math.Mod(float64(xy[0]), 1))
			// 		flag = true
			// 	}
			// 	if xy[1] > 1 {
			// 		xy[1] = float32(math.Mod(float64(xy[1]), 1))
			// 		flag = true
			// 	}
			// 	uv[i] = xy
			// }
			// if flag {
			// 	p.WriteTEXCOORD_0(uv)
			// }

			mat := p.GetMaterial()
			// 处理PBR材质
			if mat.PBRMetallicRoughness != nil {
				// 处理基础色贴图
				if mat.PBRMetallicRoughness.BaseColorTexture != nil {
					updateTextInfo(mat.PBRMetallicRoughness.BaseColorTexture)

				}

				// 处理金属粗糙度贴图
				if mat.PBRMetallicRoughness.MetallicRoughnessTexture != nil {
					updateTextInfo(mat.PBRMetallicRoughness.MetallicRoughnessTexture)
				}
			}
			// 处理法线贴图
			if mat.NormalTexture != nil {
				texInfo := mat.NormalTexture
				texture := graph.GetTexture(*texInfo.Index)
				imageIndex := texture.Source
				options := atlasOptions[*imageIndex]
				// 应用 KHR_texture_transform 扩展
				options.BindToNormalTexture(texInfo)
				// 指向新纹理
				// texInfo.Index = gltf.Index(0)
				texture.Source = gltf.Index(0)
			}
			// 处理发光贴图
			if mat.EmissiveTexture != nil {
				updateTextInfo(mat.EmissiveTexture)
			}
			// 处理遮挡贴图
			if mat.OcclusionTexture != nil {
				texInfo := mat.OcclusionTexture
				texture := graph.GetTexture(*texInfo.Index)
				imageIndex := texture.Source
				options := atlasOptions[*imageIndex]
				// 应用 KHR_texture_transform 扩展
				options.BindToOcclusionTexture(texInfo)
				// 指向新纹理
				// texInfo.Index = gltf.Index(0)
				texture.Source = gltf.Index(0)
			}
		}

	}
	// TODO 这里销毁后不一定有对齐
	images := graph.GetImages()
	for _, image := range images {
		if image.BufferView != nil {
			// 这里会销毁当前bufferview和无用的buffer，同时修改其他image的bufferview中的buffer
			image.GetBufferView().Dispose()
		}
	}

	// 保留一个默认采样器
	// graph.Samplers = graph.Samplers[:1]
	log.Println(len(graph.Samplers))
	graph.Samplers[0].WrapS = gltf.WrapRepeat
	graph.Samplers[0].WrapT = gltf.WrapRepeat
	// 创建新纹理集对应的图像
	newImage := &gltf.Image{MimeType: "image/jpeg"}
	// 创建新纹理，并指向默认采样器
	// newTexture := &gltf.Texture{Source: gltf.Index(0), Sampler: gltf.Index(0)}

	// 添加到文档
	graph.Images = []*gltf.Image{newImage}
	// graph.Textures = []*gltf.Texture{newTexture}
	//
	graph.GetImage(0).WriteBufferView(data)
	graph.AddExtension(extension.KHR_texture_transform, true)
}
