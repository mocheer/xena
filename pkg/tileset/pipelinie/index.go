package pipelinie

import "github.com/mocheer/xena/pkg/tileset/tile"

func New() *Pipeline {
	return &Pipeline{}
}

// Transform
func From(url string) *Pipeline {
	p := New()
	fromTileset, err := tile.FromPath(url)
	if err != nil {
		panic(err)
	}
	t := &TilesetWrapper{}
	t.Tileset = fromTileset
	t.URL = url
	p.Tileset = t
	return p
}
