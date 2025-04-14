package pipelinie

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/ds/ds_gltf"
	"github.com/mocheer/xena/pkg/tileset/tile"
	"github.com/mocheer/xena/pkg/tileset/transform/merge"
	"github.com/qmuntal/gltf"
)

type Pipeline struct {
	Tileset *TilesetWrapper
	Options []Option
}

type Option interface {
	TransformTileset(t *TilesetWrapper) error
	TransformTile(t *TileWrapper) error
}

// SetTransformFromMetadataXML 从osgb的metadata.xml中计算transform
// 示例2
// <ModelMetadata version="1">
// <SRS>ENU:26.034399819000001,119.216628185</SRS>
// <SRSOrigin>0,0,89.840000000000003</SRSOrigin>
// <Texture>
// <ColorSource>Visible</ColorSource>
// </Texture>
// </ModelMetadata>
// 实例2
// <ModelMetadata version="1">
// <SRS>EPSG:4549</SRS>
// <SRSOrigin>415156.8,2884887.2,-23.9</SRSOrigin>
// </ModelMetadata>
func (m *Pipeline) SetTransformFromENU(lon float64, lat float64, alt float64) {
	matrix := CalcEnuToEcefMatrix(lon, lat, alt)
	m.Tileset.Tileset.Root.Transform = &matrix
}

// Rebuild
// 从上到下重建顶层
func (m *Pipeline) Rebuild(dir string) error {
	tiles := m.Tileset.GetFirstTiles() //
	var build func(z int, num int, bv tile.BoundingVolume, tiles []*TileWrapper)
	//
	build = func(z int, num int, bv tile.BoundingVolume, tiles []*TileWrapper) {
		if z > 22 {
			return
		}
		// tiles
		doc := gltf.Document{}
		childrenTiles := []*TileWrapper{}
		for _, tile := range tiles {
			// 这里应该改用完全包含，当tile的范围被bv完全包含（包括相等）的时候，直接将tile合并，这里应该根据层级范围添加容错
			// 否则遍历tile子级，寻找被完全包含的块再包含
			// 注意，这里会有切片横跨 BoundingVolume，根据当前范围的容错，如果可以接受，直接添加到第一次遍历到的合并切片，之后不会再遍历到
			if IsBoundingBoxesOverlap(&bv, &tile.BoundingVolume, 0) {

			}

			if tile.Doc == nil {
				data, err := tile.Load()
				if err == nil {
					dr := bytes.NewReader(data)
					tile.Doc = &gltf.Document{}
					err = gltf.NewDecoder(dr).Decode(tile.Doc)
					if err != nil {
						panic(err)
					}
				} else {
					panic(err)
				}
			}
			childrenTiles = append(childrenTiles, tile)
			merge.MergeTo(&doc, tile.Doc)
		}
		filename := fmt.Sprintf("L%d_%d.glb", z, num)
		// 这里需要加一层简化、网格化、draco、ktx2
		ds_gltf.Save(filepath.Join(dir, filename), &doc, true)
		//
		z++
		for i, bv2 := range bv.QuadtreeBoundingVolumeBox() {
			n := 1<<num + i - 1
			// 这里不应该直接用tiles，应该是一个当切片被使用到就扔掉，当切片不符合要求时将切片对象的所有子级替换掉当前切片
			build(z, n, bv2, childrenTiles)
		}
	}
	// 这里不应该直接用tiles，应该是一个当切片被使用到就扔掉，当切片不符合要求时将切片对象的所有子级替换掉当前切片
	build(0, 0, m.Tileset.Root.BoundingVolume, tiles)
	return nil
}

// Transform
func (m *Pipeline) Transform(options ...Option) error {
	m.Options = append(m.Options, options...)
	err := m.transformTileset(m.Tileset)
	if err != nil {
		panic(err)
	}
	return nil
}

// transformTileset
func (m *Pipeline) transformTileset(t *TilesetWrapper) error {
	// 可能需要先遍历子级
	// 可能需要修改tile的url名称后续再保存tileset
	err := m.transformTile(&TileWrapper{Tile: t.Root, OwnerTileset: t})
	if err != nil {
		return err
	}
	//
	for _, o := range m.Options {
		err := o.TransformTileset(t)
		if err != nil {
			return err
		}
	}

	return nil

}

// transformTile
func (m *Pipeline) transformTile(t *TileWrapper) error {

	for _, o := range m.Options {
		err := o.TransformTile(t)
		if err != nil {
			return err
		}
	}

	for _, item := range t.Children {
		tile := &TileWrapper{Tile: item, OwnerTileset: t.OwnerTileset}
		err := m.transformTile(tile)
		if err != nil {
			return err
		}
		if tile.Content.IsTileset() {
			t, err := tile.LoadTileset()
			if err != nil {
				return err
			}
			m.transformTileset(t)
		}
	}
	return nil
}
