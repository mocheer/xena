package graph

import (
	"log"
	"slices"

	"github.com/qmuntal/gltf"
)

type Graph struct {
	*gltf.Document
}

// GraphBase
// gltf图表数据节点的基本抽象接口，用于Prune
type GraphBase interface {
	// 索引，这里用int，可用于 -1 区分是否被移除销毁
	Index() int
	// 是否有用
	IsUsed() bool
	//
	Dispose()
	// Extras 属性是用来存储用户数据的，可以任意数据结构
	IsEmptyExtras() bool
}

// New
func New(doc *gltf.Document) *Graph {
	graph := &Graph{Document: doc}
	return graph
}

// Open
func Open(f string) *Graph {
	doc, err := gltf.Open(f)
	if err != nil {
		panic(err)
	}
	return New(doc)
}

// Read
func (m *Graph) Read(f string) {
	doc, err := gltf.Open(f)
	if err != nil {
		panic(err)
	}
	m.Document = doc
}

// GetScenes
func (m *Graph) GetScenes() []*GraphScene {
	scenes := make([]*GraphScene, 0, len(m.Scenes))
	for _, scene := range m.Scenes {
		scenes = append(scenes, &GraphScene{
			Graph: m,
			Scene: scene,
		})
	}
	return scenes
}

// GetScene
func (m *Graph) GetScene() *GraphScene {
	return &GraphScene{
		Graph: m,
		Scene: m.Scenes[*m.Scene],
	}
}

// GetNodes
func (m *Graph) GetNodes() []*GraphNode {
	nodes := make([]*GraphNode, 0, len(m.Nodes))
	for _, node := range m.Nodes {
		nodes = append(nodes, &GraphNode{
			Graph: m,
			Node:  node,
		})
	}
	return nodes
}

// GetNode
func (m *Graph) GetNode(i uint32) *GraphNode {
	return &GraphNode{
		Graph: m,
		Node:  m.Nodes[i],
	}
}

// GetMeshes
func (m *Graph) GetMeshes() []*GraphMesh {
	meshes := make([]*GraphMesh, 0, len(m.Meshes))
	for _, mesh := range m.Meshes {
		meshes = append(meshes, &GraphMesh{
			Graph: m,
			Mesh:  mesh,
		})
	}
	return meshes
}

// GetMesh
func (m *Graph) GetMesh(i uint32) *GraphMesh {
	if int(i) >= 0 && int(i) < len(m.Meshes) {
		return &GraphMesh{
			Graph: m,
			Mesh:  m.Meshes[i],
		}
	} else {
		log.Println("获取mesh失败，索引为：", i, len(m.Meshes))
	}
	return nil
}

// GetMaterials
func (m *Graph) GetMaterials() []*GraphMaterial {
	materials := make([]*GraphMaterial, 0, len(m.Materials))
	for _, material := range m.Materials {
		materials = append(materials, &GraphMaterial{
			Graph:    m,
			Material: material,
		})
	}
	return materials
}

// GetMaterial
func (m *Graph) GetMaterial(i uint32) *GraphMaterial {
	return &GraphMaterial{
		Graph:    m,
		Material: m.Materials[i],
	}
}

// GetTextures
func (m *Graph) GetTextures() []*GraphTexture {
	textures := make([]*GraphTexture, 0, len(m.Textures))
	for _, texture := range m.Textures {
		textures = append(textures, &GraphTexture{
			Graph:   m,
			Texture: texture,
		})
	}
	return textures
}

// GetTexture
func (m *Graph) GetTexture(i uint32) *GraphTexture {
	return &GraphTexture{
		Graph:   m,
		Texture: m.Textures[i],
	}
}

// GetImages
func (m *Graph) GetImages() []*GraphImage {
	images := make([]*GraphImage, 0, len(m.Images))
	for _, image := range m.Images {
		images = append(images, &GraphImage{
			Graph: m,
			Image: image,
		})
	}
	return images
}

// GetImage
func (m *Graph) GetImage(i uint32) *GraphImage {
	return &GraphImage{
		Graph: m,
		Image: m.Images[i],
	}
}

// GetSkins
func (m *Graph) GetSkins() []*GraphSkin {
	skins := make([]*GraphSkin, 0, len(m.Skins))
	for _, skin := range m.Skins {
		skins = append(skins, &GraphSkin{
			Graph: m,
			Skin:  skin,
		})
	}
	return skins

}

// GetSamplers
func (m *Graph) GetSamplers() []*GraphSampler {
	samplers := make([]*GraphSampler, 0, len(m.Samplers))
	for _, sampler := range m.Samplers {
		samplers = append(samplers, &GraphSampler{
			Graph:   m,
			Sampler: sampler,
		})
	}
	return samplers
}
func (m *Graph) GetSampler(i uint32) *GraphSampler {
	return &GraphSampler{
		Graph:   m,
		Sampler: m.Samplers[i],
	}
}

// GetAccessors
func (m *Graph) GetAccessors() []*GraphAccessor {
	accessors := make([]*GraphAccessor, 0, len(m.Accessors))
	for _, accessor := range m.Accessors {
		accessors = append(accessors, &GraphAccessor{
			Graph:    m,
			Accessor: accessor,
		})
	}
	return accessors
}

func (m *Graph) GetAccessor(i uint32) *GraphAccessor {
	return &GraphAccessor{
		Graph:    m,
		Accessor: m.Accessors[i],
	}
}

// GetBufferViews
func (m *Graph) GetBufferViews() []*GraphBufferView {
	bufferViews := make([]*GraphBufferView, 0, len(m.BufferViews))
	for _, bufferview := range m.BufferViews {
		bufferViews = append(bufferViews, &GraphBufferView{
			Graph:      m,
			BufferView: bufferview,
		})
	}
	return bufferViews
}

func (m *Graph) GetBufferView(i uint32) *GraphBufferView {
	return &GraphBufferView{
		Graph:      m,
		BufferView: m.BufferViews[i],
	}
}

// GetBuffers
func (m *Graph) GetBuffers() []*GraphBuffer {
	buffers := make([]*GraphBuffer, 0, len(m.Buffers))
	for _, buffer := range m.Buffers {
		buffers = append(buffers, &GraphBuffer{
			Graph:  m,
			Buffer: buffer,
		})
	}
	return buffers
}

func (m *Graph) GetBuffer(i uint32) *GraphBuffer {
	return &GraphBuffer{
		Graph:  m,
		Buffer: m.Buffers[i],
	}
}

// GetCameras
func (m *Graph) GetCameras() []*GraphCamera {
	cameras := make([]*GraphCamera, 0, len(m.Cameras))
	for _, camera := range m.Cameras {
		cameras = append(cameras, &GraphCamera{
			Graph:  m,
			Camera: camera,
		})
	}
	return cameras

}

// AddMesh
func (m *Graph) AddMesh(mesh *gltf.Mesh) {
	meshIndex := len(m.Meshes)
	nodeIndex := len(m.Nodes)
	m.Meshes = append(m.Meshes, mesh)
	node := &gltf.Node{Mesh: gltf.Index(uint32(meshIndex))}
	m.Nodes = append(m.Nodes, node)
	scene := m.GetScene()
	scene.Nodes = append(scene.Nodes, uint32(nodeIndex))
}

// DiscardUnusedScenes
// 直接丢弃没有用到的场景
// TODO 丢弃后有些node没有用到，相关资源也会没有用到，需要后续处理
func (m *Graph) DiscardUnusedScenes() {
	m.Scenes = []*gltf.Scene{m.GetScene().Scene}
	*m.Scene = 0
}

// FindVisibleNodes
func (m *Graph) FindVisibleNodes(fn func(node *GraphNode) bool) bool {
	return m.GetScene().FindUsedNodes(fn)
}

// FindUsedNodes
func (m *Graph) FindUsedNodes(fn func(node *GraphNode) bool) bool {
	for _, scene := range m.GetScenes() {
		if scene.FindUsedNodes(fn) {
			return true
		}
	}
	return true
}

// EachVisiableNodes 只遍历默认场景下可见的节点
func (m *Graph) EachVisibleNodes(fn func(node *GraphNode)) {
	m.GetScene().EachUsedNodes(fn)
}

// EachUsedNodes 遍历所有场景用到的节点
func (m *Graph) EachUsedNodes(fn func(node *GraphNode)) {
	for _, scene := range m.GetScenes() {
		scene.EachUsedNodes(fn)
	}
}

// EachPrimitives
func (m *Graph) EachPrimitives(fn func(*GraphPrimitive)) {
	for _, node := range m.GetNodes() {
		mesh := node.GetMesh()
		if mesh != nil {
			for _, p := range mesh.GetPrimitives() {
				fn(p)
			}
		}

	}
}

// 添加非必需拓展
// 如果需要声明拓展是必需的，请再将该拓展添加到 extensionsRequired，注意 extensionsRequired的拓展必须也在 extensionsUsed 中声明
func (m *Graph) AddExtension(name string, required bool) {
	if !slices.Contains(m.ExtensionsUsed, name) {
		m.ExtensionsUsed = append(m.ExtensionsUsed, name)
	}
	if required {
		m.ExtensionsRequired = append(m.ExtensionsRequired, name)
	}
}

// RemoveExtension 移除拓展
func (m *Graph) RemoveExtension(name string) {
	m.ExtensionsUsed = slices.DeleteFunc(m.ExtensionsUsed, func(ext string) bool {
		return ext == name
	})
	m.ExtensionsRequired = slices.DeleteFunc(m.ExtensionsRequired, func(ext string) bool {
		return ext == name
	})
}
