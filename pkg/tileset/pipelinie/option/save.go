package option

import (
	"bytes"
	"path"
	"path/filepath"

	"github.com/mocheer/pluto/pkg/ds"
	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/xena/pkg/tileset/pipelinie"
	"github.com/qmuntal/gltf"
)

type SaveOption struct {
	Path string
}

// TransformTileset
func (m *SaveOption) TransformTileset(t *pipelinie.TilesetWrapper) error {
	// 只需要保存根路径的tileset.json
	// 子集的tileset已经由tile的content保存了。
	if t.OwnerTile == nil {
		return ds_json.Save(path.Join(m.Path, "tileset.json"), t.Tileset)
	}
	url := filepath.Join(m.Path, t.OwnerTile.OwnerTileset.RelativePath, t.OwnerTile.Content.GetURL())
	return ds_json.Save(url, t.Tileset)

}

// TransformTile
func (m *SaveOption) TransformTile(t *pipelinie.TileWrapper) error {
	if t.Content == nil || t.ContentTileset != nil {
		return nil
	}
	url := filepath.Join(m.Path, t.OwnerTileset.RelativePath, t.Content.GetURL())
	if t.Doc != nil {
		var buf bytes.Buffer
		enc := gltf.NewEncoder(&buf)
		enc.AsBinary = true
		enc.Encode(t.Doc)
		ds.Save(url, buf.Bytes())
		return nil
	}
	data, err := t.Load()
	if err != nil {
		return err
	}
	return ds.Save(url, data)
}
