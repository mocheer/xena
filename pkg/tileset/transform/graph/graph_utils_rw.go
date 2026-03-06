package graph

import (
	"image"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

func (m *Graph) ReadBufferView(bv *gltf.BufferView) ([]byte, error) {
	return modeler.ReadBufferView(m.Document, bv)
}

// ReadAsIndices 读取索引
func (m *Graph) ReadAsIndices(accessor *gltf.Accessor) ([]uint32, error) {
	return modeler.ReadIndices(m.Document, accessor, nil)
}

// ReadAsPosition 读取顶点坐标集
func (m *Graph) ReadAsPosition(accessor *gltf.Accessor) ([][3]float32, error) {
	return modeler.ReadPosition(m.Document, accessor, nil)
}

// ReadAsTextureCoord 读取纹理坐标集
func (m *Graph) ReadAsTextureCoord(accessor *gltf.Accessor) ([][2]float32, error) {
	return modeler.ReadTextureCoord(m.Document, accessor, nil)
}

// Write
// 不会销毁当前BufferView和其对应的Buffer，实际上是追加新的BufferView，并且往最后一个Buffer数据写入数据（如果不存在buffer则创建）
// 当前BufferView有可能是根其他节点共享，所以不要贸然移除BufferView和其对应的Buffer
// target 是一个可选属性。对于不直接用于渲染的中间数据（如动画数据、 morph target 权重等），可以省略 target。但用于渲染的顶点属性和索引，必须正确设置 target。
// 返回新BufferView的索引
func (m Graph) WriteBufferView(data []byte, target gltf.Target) uint32 {
	//太夸张了，因为URI不为空，即使Data已经拓展了，保存的时候会直接读取反而导致Data的长度大于uri的长度,详细看gltf的encode
	lenBuffer := len(m.Buffers)
	if lenBuffer > 0 {
		lastBuffer := m.Buffers[lenBuffer-1]
		if lastBuffer.URI != "" {
			lastBuffer.URI = ""
		}
	}
	return modeler.WriteBufferView(m.Document, target, data)
}

// WriteIndices
func (m *Graph) WriteIndices(indices []uint32) uint32 {
	return modeler.WriteIndices(m.Document, indices)
}

// WritePostion
func (m *Graph) WritePostion(position [][3]float32) uint32 {
	return modeler.WritePosition(m.Document, position)
}

// WriteTextureCoord
func (m *Graph) WriteTextureCoord(textureCoord [][2]float32) uint32 {
	return modeler.WriteTextureCoord(m.Document, textureCoord)
}

// ReadImagesAsImages 加载glTF中的所有图像
// 用例1：用于纹理合并，用于雪碧图制作
func (m *Graph) ReadImagesAsImages() ([]image.Image, error) {
	images := make([]image.Image, len(m.Images))
	for i, img := range m.GetImages() {
		src, err := img.ReadAsImage()
		if err != nil {
			return nil, err
		}
		images[i] = src
	}
	return images, nil
}
