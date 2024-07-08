package pipelinie

import (
	"github.com/mocheer/xena/pkg/tileset"
)

func New() *Pipeline {
	return &Pipeline{}
}

// Transform
func From(url string) *Pipeline {
	p := New()
	fromTileset, err := tileset.FromPath(url)
	if err != nil {
		panic(err)
	}
	t := &TilesetWrapper{}
	t.Tileset = fromTileset
	t.URL = url
	p.Tileset = t
	return p
}
