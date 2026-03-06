package graph

import (
	"slices"

	"github.com/qmuntal/gltf"
)

// bufferView 相当于 buffer的切片
type GraphBufferView struct {
	Graph *Graph
	*gltf.BufferView
}

// Index
func (m *GraphBufferView) Index() int {
	return slices.Index(m.Graph.BufferViews, m.BufferView)
}

// IsUsed
func (m *GraphBufferView) IsUsed() bool {
	index := m.Index()
	if index != -1 {
		targetIndex := uint32(index)
		// 除了accessor有依赖多图片纹理也会直接依赖BufferView
		for _, accessor := range m.Graph.Accessors {
			if *accessor.BufferView == targetIndex {
				return true
			}
		}
		// 检查材质纹理是否依赖此BufferView（如KHR_texture_basisu扩展）
		for _, image := range m.Graph.Images {
			if *image.BufferView == targetIndex {
				return true
			}
		}
	}
	return false
}

// Dispose
// 销毁后需要修正所有后置索引
// 当前BufferView销毁后，对应的Buffer可能没有任何节点使用，需要检测后销毁
func (m *GraphBufferView) Dispose() {
	index := m.Index()
	if index != -1 {
		m.Graph.BufferViews = slices.Delete(m.Graph.BufferViews, index, index+1)
		targetIndex := uint32(index)

		// 修改 Images 中的bufferview索引
		images := []*gltf.Image{}
		for _, image := range m.Graph.Images {
			if *image.BufferView == targetIndex {
				continue
			}
			if *image.BufferView > targetIndex {
				*image.BufferView -= 1
			}
			images = append(images, image)
		}
		m.Graph.Images = images

		// 尝试删除关联的buffer
		buffer := m.Graph.GetBuffer(m.Buffer)
		// 没有任何指向Buffer数据的索引时，删除整个Buffer,并更新索引
		if !buffer.IsUsed() {
			buffer.Dispose()
		} else {
			// 很多数据是单缓冲区，根据BufferView的偏移量+长度加载的数据，即使是多纹理，原始纹理也不一定是单独的Buffer，更不用说POSITION、NORMAL等
			// 这里应该要考虑删除这部分数据，并拿到所有BufferView修改偏移量
			// 不过前提还是确保这部分数据依旧没有其他节点引用
			// for _, view := range m.Graph.BufferViews {
			// 	// 一般来说只需要
			// 	// if view.Buffer == m.Buffer && view.ByteOffset >= m.ByteOffset && (view.ByteOffset+view.ByteLength) < (m.ByteOffset+m.ByteLength) {
			// 	if view.Buffer == m.Buffer && view.ByteOffset == m.ByteOffset {
			// 		// 共享
			// 		return
			// 	}
			// }

			// 这里没有移除用于组件类型占位对齐的字节
			// 这里可能导致对齐出现问题
			count := m.ByteLength
			if count%4 != 0 {
				count -= count % 4
			}
			newBufferData := append(
				buffer.Data[:m.ByteOffset],
				buffer.Data[m.ByteOffset+count:]...,
			)
			buffer.Data = newBufferData
			buffer.ByteLength = uint32(len(newBufferData))
			// 更新其他BufferView的偏移量
			for _, view := range m.Graph.BufferViews {

				if view.Buffer == m.Buffer && view.ByteOffset > m.ByteOffset {
					// 这里有可能包含用于组件对齐的字节，导致偏移量不是组件类型的倍数
					view.ByteOffset -= count
				}
			}
		}

		// 修改 Accessors 中的bufferview索引
		// 这里单纯修改BufferView的指向是不够的，访问器的总字节偏移必须是组件类型长度的整数倍
		accessors := []*gltf.Accessor{}
		for _, accessor := range m.Graph.Accessors {
			if *accessor.BufferView == targetIndex {
				continue
			}
			if *accessor.BufferView > targetIndex {
				*accessor.BufferView -= 1
				// TODO 这里需要验证下偏移量是否对齐
				// accessor.ComponentType.ByteSize()
			}
			accessors = append(accessors, accessor)
		}
		m.Graph.Accessors = accessors

	}
}

// IsEmptyExtras
func (m *GraphBufferView) IsEmptyExtras() bool {
	return m.Extras == nil
}

// Read
func (m GraphBufferView) Read() ([]byte, error) {
	return m.Graph.ReadBufferView(m.BufferView)
}

// Write
// 不会销毁当前BufferView和其对应的Buffer，实际上是追加新的BufferView和Buffer数据
// 不要贸然移除BufferView和其对应的Buffer，当前BufferView有可能是根其他节点共享
// 返回新BufferView的索引
func (m GraphBufferView) Write(data []byte) uint32 {
	return m.Graph.WriteBufferView(data, m.BufferView.Target)
}

// Replace
// TODO 未完成
// 如果 bufferview对应的buffer只有当前的bufferview使用，那可以直接丢弃buffer，修改所有后置的buffer索引（包括当前）
// 如果 有多个地方使用，那需要删除数据，并修改对应的bufferView的索引
// func (m GraphBufferView) Replace(data []byte) uint32 {
// 	// 记录当前索引，销毁当前数据
// 	return m.Write(data)
// }
