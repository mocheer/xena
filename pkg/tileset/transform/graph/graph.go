package graph

import (
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

type Graph struct {
	Doc         *gltf.Document
	Scene       *GraphScene
	Scenes      []*GraphScene //有点雪碧图的味道
	Nodes       []*GraphNode
	Meshes      []*GraphMesh
	Materials   []*GraphMaterial
	Textures    []*GraphTexture
	Images      []*GraphImage
	Samplers    []*GraphSampler
	Skins       []*GraphSkin
	Cameras     []*GraphCamera
	Animations  []*GraphAnimation
	BufferViews []*GraphBufferView
	Accessors   []*GraphAccessor
	Buffers     []*GraphBuffer
}

type GraphBase interface {
	IsUsed() bool
	//
	Dispose()
	// Extras 属性是用来存储用户数据的，可以任意数据结构
	IsEmptyExtras() bool
}

// 连接两个实例的数据信息
type Edge struct {
	Parent Parent
}

type Parent struct {
	PropertyType string
}

// NewGraph
func NewGraph(doc *gltf.Document) *Graph {

	graph := &Graph{Doc: doc}
	graph.InitWithDocument(doc)
	return graph
}

// InitWithDocument
func (m *Graph) InitWithDocument(doc *gltf.Document) {
	m.Doc = doc
	m.Scenes = make([]*GraphScene, 0, len(doc.Scenes))
	for _, scene := range doc.Scenes {
		m.AppendScene(scene)
	}
	m.Scene = m.Scenes[*doc.Scene]
	m.Nodes = make([]*GraphNode, 0, len(doc.Nodes))
	for _, node := range doc.Nodes {
		m.AppendNode(node)
	}
	m.Meshes = make([]*GraphMesh, 0, len(doc.Meshes))
	for _, mesh := range doc.Meshes {
		m.AppendMesh(mesh)
	}
	m.Materials = make([]*GraphMaterial, 0, len(doc.Materials))
	for _, material := range doc.Materials {
		m.AppendMaterial(material)
	}
	m.Textures = make([]*GraphTexture, 0, len(doc.Textures))
	for _, texture := range doc.Textures {
		m.AppendTexture(texture)
	}
	m.Images = make([]*GraphImage, 0, len(doc.Images))
	for _, image := range doc.Images {
		m.AppendImage(image)
	}
	m.Samplers = make([]*GraphSampler, 0, len(doc.Samplers))
	for _, sampler := range doc.Samplers {
		m.AppendSampler(sampler)
	}
	m.Skins = make([]*GraphSkin, 0, len(doc.Skins))
	for _, skin := range doc.Skins {
		m.AppendSkin(skin)
	}

	m.Accessors = make([]*GraphAccessor, 0, len(doc.Accessors))
	for _, bufferview := range doc.Accessors {
		m.AppendAccessor(bufferview)
	}

	m.BufferViews = make([]*GraphBufferView, 0, len(doc.BufferViews))
	for _, bufferview := range doc.BufferViews {
		m.AppendBufferview(bufferview)
	}

	m.Buffers = make([]*GraphBuffer, 0, len(doc.Buffers))
	for _, bufferview := range doc.Buffers {
		m.AppendBuffer(bufferview)
	}
}

// AppendScene
func (m *Graph) AppendScene(scene *gltf.Scene) {
	m.Scenes = append(m.Scenes, &GraphScene{
		Graph: m,
		Scene: scene,
	})
}

// AppendNode
func (m *Graph) AppendNode(node *gltf.Node) {
	// index := uint32(len(m.Nodes))
	m.Nodes = append(m.Nodes, &GraphNode{
		Graph: m,
		Node:  node,
		// Index: &index,
	})
}

// AppendMesh
func (m *Graph) AppendMesh(mesh *gltf.Mesh) {
	m.Meshes = append(m.Meshes, &GraphMesh{
		Graph: m,
		Mesh:  mesh,
	})
}

// AppendMaterial
func (m *Graph) AppendMaterial(material *gltf.Material) {
	m.Materials = append(m.Materials, &GraphMaterial{
		Graph:    m,
		Material: material,
	})
}

// AppendSkin
func (m *Graph) AppendSkin(skin *gltf.Skin) {
	m.Skins = append(m.Skins, &GraphSkin{
		Graph: m,
		Skin:  skin,
	})
}

// AppendTexture
func (m *Graph) AppendTexture(texture *gltf.Texture) {
	m.Textures = append(m.Textures, &GraphTexture{
		Graph:   m,
		Texture: texture,
	})
}

// AppendImage
func (m *Graph) AppendImage(image *gltf.Image) {
	m.Images = append(m.Images, &GraphImage{
		Graph: m,
		Image: image,
	})
}

// AppendSampler
func (m *Graph) AppendSampler(sampler *gltf.Sampler) {
	m.Samplers = append(m.Samplers, &GraphSampler{
		Graph:   m,
		Sampler: sampler,
	})
}

// AppendAccessor
func (m *Graph) AppendAccessor(accessor *gltf.Accessor) {
	m.Accessors = append(m.Accessors, &GraphAccessor{
		Graph:    m,
		Accessor: accessor,
	})
}

// AppendBuffer
func (m *Graph) AppendBuffer(buffer *gltf.Buffer) {
	m.Buffers = append(m.Buffers, &GraphBuffer{
		Graph:  m,
		Buffer: buffer,
	})
}

// AppendBufferview
func (m *Graph) AppendBufferview(bufferview *gltf.BufferView) {
	m.BufferViews = append(m.BufferViews, &GraphBufferView{
		Graph:      m,
		BufferView: bufferview,
	})
}

// AppendBufferview
func (m *Graph) AddMesh(mesh *gltf.Mesh) {
	doc := m.Doc
	meshIndex := len(doc.Meshes)
	nodeIndex := len(doc.Nodes)
	doc.Meshes = append(doc.Meshes, mesh)
	m.AppendMesh(mesh)
	node := &gltf.Node{Mesh: gltf.Index(uint32(meshIndex))}
	doc.Nodes = append(doc.Nodes, node)
	m.AppendNode(node)
	m.Scene.Nodes = append(m.Scene.Nodes, uint32(nodeIndex))
}

// DiscardScenes
// 直接丢弃没有用到的场景
func (m *Graph) DiscardScenes() {
	m.Scenes = []*GraphScene{m.Scene}
	m.Doc.Scenes = []*gltf.Scene{m.Scene.Scene}
	*m.Doc.Scene = 0
}

// FindVisableNodes
func (m *Graph) FindVisableNodes(fn func(node *GraphNode) bool) bool {
	return m.Scene.FindUsedNodes(fn)
}

// FindUsedNodes
func (m *Graph) FindUsedNodes(fn func(node *GraphNode) bool) bool {
	for _, scene := range m.Scenes {
		if scene.FindUsedNodes(fn) {
			return true
		}
	}
	return true
}

func (m *Graph) WriteIndices(indices []uint32) uint32 {
	return modeler.WriteIndices(m.Doc, indices)
}

func (m *Graph) WritePostion(position [][3]float32) uint32 {
	return modeler.WritePosition(m.Doc, position)
}

func (m *Graph) WriteTextureCoord(textureCoord [][2]float32) uint32 {
	return modeler.WriteTextureCoord(m.Doc, textureCoord)
}

// EachVisiableNodes
func (m *Graph) EachVisiableNodes(fn func(node *GraphNode)) {
	m.Scene.EachUsedNodes(fn)
}

// EachUsedNodes
func (m *Graph) EachUsedNodes(fn func(node *GraphNode)) {
	for _, scene := range m.Scenes {
		scene.EachUsedNodes(fn)
	}
}
