package option

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mocheer/xena/pkg/tileset/b3dm"
	"github.com/mocheer/xena/pkg/tileset/pipelinie"
)

type UpgradeOption struct {
	Version string
}

// Upgrade 升级版本
func (m UpgradeOption) TransformTileset(t *pipelinie.TilesetWrapper) error {
	if t.Asset.Version != m.Version {
		switch t.Asset.Version {
		case "1.0":
			switch m.Version {
			case "1.1":

			}
		}
		//
		t.Asset.Version = m.Version
	}
	return nil
}

// Upgrade 升级版本
func (m UpgradeOption) TransformTile(t *pipelinie.TileWrapper) error {
	// b3dm to gltf
	if t.Content != nil && !t.Content.IsTileset() {
		url := t.GetContentURL()
		b, err := b3dm.Open(url)
		if err != nil {
			return err
		}
		if filepath.Ext(url) == ".b3dm" {
			t.Content.Uri = strings.Replace(t.Content.GetURL(), ".b3dm", ".glb", -1)
			t.Content.Url = nil
		}

		t.Doc = b.Model
		fmt.Println(b.BatchTable, b.FeatureTable)
	}

	return nil
}
