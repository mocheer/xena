package graph

import (
	"bytes"
	"errors"
	"image"
	"os"
	"slices"
	"strings"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/qmuntal/gltf"
)

type GraphImage struct {
	*gltf.Image
	Graph *Graph
}

// Index
func (m *GraphImage) Index() int {
	return slices.Index(m.Graph.Images, m.Image)
}

// IsUsed
func (m *GraphImage) IsUsed() bool {
	targetIndex := uint32(m.Index())
	for _, texture := range m.Graph.Textures {
		if *texture.Source == targetIndex {
			return true
		}
	}
	return false
}

// UseCount
// 使用次数
func (m *GraphImage) UseCount() int {
	count := 0
	targetIndex := uint32(m.Index())
	for _, texture := range m.Graph.Textures {
		if *texture.Source == targetIndex {
			count++
		}
	}
	return count
}

// Dispose
func (m *GraphImage) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.Images = slices.Delete(m.Graph.Images, index, index+1)
		// 修改索引
		targetIndex := uint32(index)
		for _, texture := range m.Graph.Textures {
			if *texture.Source > targetIndex {
				*texture.Source--
			}
		}
		bfv := m.GetBufferView()
		if bfv != nil && !bfv.IsUsed() {
			bfv.Dispose()
		}
	}
}

// GetColorSpace
func (m *GraphImage) GetColorSpace() string {
	// 外部图像默认为 sRGB 色彩空间编码
	if m.URI != "" {
		return "sRGB"
	}
	return ""
}

// GetBufferView
func (m GraphImage) GetBufferView() *GraphBufferView {
	if m.BufferView != nil {
		return m.Graph.GetBufferView(*m.BufferView)
	}
	return nil
}

// Read
// 如果bufferView为空，则读取uri
func (m GraphImage) Read() ([]byte, error) {
	if m.BufferView != nil {
		data, err := m.ReadBufferView()
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	// 可能是base64
	if m.URI != "" {
		if strings.HasPrefix(m.URI, "data:image") {
			data := fn.Atob2BytesWithURI(m.URI)
			return data, nil
		}
		return ds.ReadFile(m.URI)
	}
	return nil, errors.New("当前数据为空")
}

// ReadBufferView
func (m GraphImage) ReadBufferView() ([]byte, error) {
	// 有 URI 就没有BuferView
	if m.BufferView != nil {
		bfv := m.Graph.GetBufferView(*m.BufferView)
		data, err := bfv.Read()
		return data, err
	}
	return nil, errors.New("not found")
}

// WriteBufferView
func (m GraphImage) WriteBufferView(data []byte) {
	if m.BufferView != nil {
		bfv := m.Graph.GetBufferView(*m.BufferView)
		bufferviewIndex := bfv.Write(data)
		m.BufferView = &bufferviewIndex
	} else {
		// TargetNone
		// ARRAY_BUFFER 一般用于几何数据
		// ELEMENT_ARRAY_BUFFER 一般用于索引数据
		m.BufferView = gltf.Index(m.Graph.WriteBufferView(data, gltf.TargetNone))
	}
}

// ReadImage 读取图像
func (m GraphImage) ReadAsImage() (image.Image, error) {
	img := m.Image
	var imgData []byte
	var err error
	if img.BufferView != nil {
		imgData, err = m.ReadBufferView()
		if err != nil {
			return nil, err
		}
	} else if img.URI != "" {
		// 从外部文件加载
		imgData, err = os.ReadFile(img.URI)
		if err != nil {
			return nil, err
		}
	}
	// 解码图像
	src, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}
	return src, nil
}
