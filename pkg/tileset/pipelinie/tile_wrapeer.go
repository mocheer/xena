package pipelinie

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ts/ctp"
	"github.com/mocheer/xena/pkg/tileset"
	"github.com/qmuntal/gltf"
)

type TileWrapper struct {
	*tileset.Tile
	// 用于辅助加载和保存
	Doc            *gltf.Document
	OwnerTileset   *TilesetWrapper
	ContentTileset *TilesetWrapper
}

func (m *TileWrapper) GetContentURL() string {
	if m.Content == nil {
		return ""
	}
	return path.Join(m.OwnerTileset.GetBaseURL(), m.Content.GetURL())
}

func (m *TileWrapper) GetDoc() string {
	return path.Join(m.OwnerTileset.GetBaseURL(), m.Content.GetURL())
}

// IsTileset
func (m *TileWrapper) Load() (data []byte, err error) {
	url := m.GetContentURL()
	if strings.HasPrefix(url, "http") {
		data, err = ctp.Get(url)
	} else {
		data, err = ds.ReadFile(url)
	}
	return

}

// IsTileset
func (m *TileWrapper) LoadTileset() (*TilesetWrapper, error) {

	if m.Content.IsTileset() {
		data, err := m.Load()
		if err != nil {
			return nil, err
		}
		t := &TilesetWrapper{Tileset: tileset.FromBytes(data), OwnerTile: m}
		t.URL = m.GetContentURL()
		t.RelativePath = filepath.Join(m.OwnerTileset.RelativePath, path.Dir(m.Content.GetURL()))
		m.ContentTileset = t
	}
	return m.ContentTileset, nil
}
