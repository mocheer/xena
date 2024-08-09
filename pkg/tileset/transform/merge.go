package transform

import (
	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"
)

type transformMerge struct {
	target *gltf.Document
	doc    *gltf.Document
	//
	cacheNode map[uint32]uint32
	cacheMesh map[*uint32]*uint32
	// 存储索引
	cacheIndices      map[*uint32]*uint32
	cachePositions    map[uint32]uint32
	cacheTextureCoord map[uint32]uint32
	cacheMaterial     map[*uint32]*uint32
	cacheTexture      map[uint32]uint32
	// image
	cacheSource  map[*uint32]*uint32
	cacheSampler map[*uint32]*uint32
}

// MergeTo
// 只合并Scene中的Node，目前不合并Scenes的所有场景
func MergeTo(target *gltf.Document, doc *gltf.Document) error {
	m := (&transformMerge{target: target, doc: doc})
	m.cacheNode = map[uint32]uint32{}
	m.cacheMesh = map[*uint32]*uint32{}
	// 存储索引
	m.cacheIndices = map[*uint32]*uint32{}
	m.cachePositions = map[uint32]uint32{}
	m.cacheTextureCoord = map[uint32]uint32{}
	m.cacheMaterial = map[*uint32]*uint32{}
	m.cacheTexture = map[uint32]uint32{}
	// image
	m.cacheSource = map[*uint32]*uint32{}
	m.cacheSampler = map[*uint32]*uint32{}
	return m.merge()
}

// merge
func (m *transformMerge) merge() error {
	target := m.target
	doc := m.doc
	cacheNode := m.cacheNode
	//
	scene := doc.Scenes[*doc.Scene]
	for _, nodeIndex := range scene.Nodes {
		nIndex, ok := cacheNode[nodeIndex]
		if !ok {
			m.appendNode(nodeIndex, &nIndex)
		}
		// 可能是空的scenes
		if target.Scenes == nil {
			target.Scenes = []*gltf.Scene{{Nodes: []uint32{}}}
		}
		target.Scenes[0].Nodes = append(target.Scenes[0].Nodes, nIndex)
	}
	return nil
}

// appendNode
func (m *transformMerge) appendNode(index uint32, newIndex *uint32) error {
	target := m.target
	doc := m.doc
	//
	cacheNode := m.cacheNode
	cacheMesh := m.cacheMesh
	// 存储索引
	cacheIndices := m.cacheIndices
	cachePositions := m.cachePositions
	cacheTextureCoord := m.cacheTextureCoord
	cacheMaterial := m.cacheMaterial
	//
	node := doc.Nodes[index]
	meshIndex, ok := cacheMesh[node.Mesh]
	if !ok {
		mesh := doc.Meshes[*node.Mesh]
		newMesh := &gltf.Mesh{
			Name:       mesh.Name,
			Extensions: mesh.Extensions,
			Weights:    mesh.Weights,
		}
		for _, p := range mesh.Primitives {
			//
			primitive := &gltf.Primitive{
				Extensions: p.Extensions,
				Mode:       p.Mode,
				Attributes: gltf.Attribute{},
			}
			newMesh.Primitives = append(newMesh.Primitives, primitive)
			// 索引
			indicesIndex, ok := cacheIndices[p.Indices]
			if ok {
				primitive.Indices = indicesIndex
			} else {
				indices, err := modeler.ReadIndices(doc, doc.Accessors[*p.Indices], nil)
				if err == nil {
					primitive.Indices = gltf.Index(modeler.WriteIndices(target, indices))
					cacheIndices[p.Indices] = primitive.Indices
				}
			}
			// 位置
			POSITION, ok := p.Attributes[gltf.POSITION]
			if ok {
				positionIndex, ok := cachePositions[POSITION]
				if ok {
					primitive.Attributes[gltf.POSITION] = positionIndex
				} else {
					position, err := modeler.ReadPosition(doc, doc.Accessors[POSITION], nil)
					if err == nil {
						primitive.Attributes[gltf.POSITION] = modeler.WritePosition(target, position)
						cachePositions[POSITION] = primitive.Attributes[gltf.POSITION]
					}
				}
			}
			// 纹理
			TEXCOORD_0, ok := p.Attributes[gltf.TEXCOORD_0]
			if ok {
				textureCoord, ok := cacheTextureCoord[TEXCOORD_0]
				if ok {
					primitive.Attributes[gltf.TEXCOORD_0] = textureCoord
				} else {
					textureCoord, err := modeler.ReadTextureCoord(doc, doc.Accessors[TEXCOORD_0], nil)
					if err == nil {
						primitive.Attributes[gltf.TEXCOORD_0] = modeler.WriteTextureCoord(target, textureCoord)
						cacheTextureCoord[TEXCOORD_0] = primitive.Attributes[gltf.TEXCOORD_0]
					}
				}
			}
			// 纹理
			TEXCOORD_1, ok := p.Attributes[gltf.TEXCOORD_1]
			if ok {
				textureCoord, err := modeler.ReadTextureCoord(doc, doc.Accessors[TEXCOORD_1], nil)
				if err == nil {
					primitive.Attributes[gltf.TEXCOORD_1] = modeler.WriteTextureCoord(target, textureCoord)
				}
			}
			// 材质
			materialIndex, ok := cacheMaterial[p.Material]
			if ok {
				primitive.Material = materialIndex
			} else {
				material := doc.Materials[*p.Material]
				newMaterial := &gltf.Material{}
				materialIndex = gltf.Index(uint32(len(target.Materials)))
				target.Materials = append(target.Materials, newMaterial)
				primitive.Material = materialIndex
				cacheMaterial[p.Material] = materialIndex
				// 纹理
				if material.PBRMetallicRoughness != nil {
					newMaterial.PBRMetallicRoughness = &gltf.PBRMetallicRoughness{}

					if material.PBRMetallicRoughness.BaseColorTexture != nil {
						newMaterial.PBRMetallicRoughness.BaseColorTexture = &gltf.TextureInfo{
							Extensions: material.PBRMetallicRoughness.BaseColorTexture.Extensions,
							// pbrMetallicRoughness中的metallicFactor和roughnessFactor 没有解析
							// Extras: material.PBRMetallicRoughness.BaseColorTexture.Extras,
						}
						m.appendTexture(material.PBRMetallicRoughness.BaseColorTexture.Index, &newMaterial.PBRMetallicRoughness.BaseColorTexture.Index)
					}
					if material.PBRMetallicRoughness.MetallicRoughnessTexture != nil {
						newMaterial.PBRMetallicRoughness.MetallicRoughnessTexture = &gltf.TextureInfo{}
						m.appendTexture(material.PBRMetallicRoughness.MetallicRoughnessTexture.Index, &newMaterial.PBRMetallicRoughness.MetallicRoughnessTexture.Index)
					}
				}
				if material.NormalTexture != nil {
					newMaterial.NormalTexture = &gltf.NormalTexture{}
					m.appendTexture(*material.NormalTexture.Index, newMaterial.NormalTexture.Index)
				}
				if material.OcclusionTexture != nil {
					newMaterial.OcclusionTexture = &gltf.OcclusionTexture{}
					m.appendTexture(*material.OcclusionTexture.Index, newMaterial.OcclusionTexture.Index)
				}
				if material.EmissiveTexture != nil {

					newMaterial.EmissiveTexture = &gltf.TextureInfo{}
					m.appendTexture(material.EmissiveTexture.Index, &newMaterial.EmissiveTexture.Index)
				}

			}
		}
		meshIndex = gltf.Index(uint32(len(target.Meshes)))
		target.Meshes = append(target.Meshes, newMesh)
	}
	*newIndex = uint32(len(target.Nodes))
	// TODO 遍历children
	target.Nodes = append(target.Nodes, &gltf.Node{
		Name:        node.Name,
		Mesh:        meshIndex,
		Camera:      node.Camera,
		Skin:        node.Skin,
		Matrix:      node.Matrix,
		Extensions:  node.Extensions,
		Rotation:    node.Rotation,
		Scale:       node.Scale,
		Translation: node.Translation,
		Weights:     node.Weights,
	})
	children := make([]uint32, len(node.Children))
	for i, itemNodeIndex := range node.Children {
		nIndex, ok := cacheNode[itemNodeIndex]
		if !ok {
			m.appendNode(itemNodeIndex, &nIndex)
		}
		children[i] = nIndex
	}
	return nil
}

// appendTexture
func (m *transformMerge) appendTexture(index uint32, newIndex *uint32) error {
	doc := m.doc
	target := m.target
	cacheTexture := m.cacheTexture
	// image
	cacheSource := m.cacheSource
	cacheSampler := m.cacheSampler
	//
	textureIndex, ok := cacheTexture[index]
	if ok {
		*newIndex = textureIndex
	} else {
		//
		t := doc.Textures[index]
		// 纹理图片
		sourceIndex, ok := cacheSource[t.Source]
		if !ok { //

			image := doc.Images[*t.Source]
			newImage := &gltf.Image{
				MimeType: image.MimeType,
				URI:      image.URI,
			}
			// 有 URI 就没有BufferView
			if image.BufferView != nil {
				bfv := doc.BufferViews[*image.BufferView]
				data, err := modeler.ReadBufferView(doc, bfv)
				if err != nil {
					return err
				}
				bufferviewIndex := modeler.WriteBufferView(target, bfv.Target, data)
				newImage.BufferView = &bufferviewIndex
			}
			sourceIndex = gltf.Index(uint32(len(target.Images)))
			target.Images = append(target.Images, newImage)
			cacheSource[t.Source] = sourceIndex

		}
		// 采样器
		samplerIndex, ok := cacheSampler[t.Sampler]
		if !ok {
			sampler := doc.Samplers[*t.Sampler]
			//
			samplerIndex = gltf.Index(uint32(len(target.Samplers)))
			target.Samplers = append(target.Samplers, sampler)
			cacheSampler[t.Sampler] = samplerIndex
		}
		//
		*newIndex = uint32(len(target.Textures))
		target.Textures = append(target.Textures, &gltf.Texture{
			Extensions: t.Extensions,
			Name:       t.Name,
			Source:     sourceIndex,
			Sampler:    samplerIndex,
		})

	}
	return nil
}
