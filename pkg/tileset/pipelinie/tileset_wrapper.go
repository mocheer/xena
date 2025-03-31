package pipelinie

import (
	"path"

	"github.com/mocheer/xena/pkg/tileset/tile"
)

type TilesetWrapper struct {
	*tile.Tileset
	// 不在tileset规范中，这里用作content的相对路径
	// 用于辅助加载和保存
	OwnerTile    *TileWrapper
	URL          string
	RelativePath string
}

// SetBaseURL
func (m *TilesetWrapper) GetBaseURL() string {
	return path.Dir(m.URL)
}

// GetFirstTiles
// 用于重建第一层，第一层只有一个模型
func (m *TilesetWrapper) GetFirstTiles() (tiles []*TileWrapper) {
	var eachTile func(t *TileWrapper)
	eachTile = func(t *TileWrapper) {
		if t.Content != nil {
			if !t.Content.IsTileset() {
				tiles = append(tiles, t)
			} else {

				tileset, err := t.LoadTileset()
				if err != nil {
					panic(err)
				}
				eachTile(&TileWrapper{Tile: tileset.Root, OwnerTileset: tileset})
			}
		} else {
			for _, tile := range t.Children {
				tw := &TileWrapper{Tile: tile, OwnerTileset: t.OwnerTileset}
				eachTile(tw)
			}
		}

	}
	eachTile(&TileWrapper{Tile: m.Tileset.Root, OwnerTileset: m})

	return
}
