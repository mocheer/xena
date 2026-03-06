package tile

import (
	"encoding/json"
)

const (
	TILE_REFINE_ADD     = "ADD"
	TILE_REFINE_REPLACE = "REPLACE" //
)

var (
	TileDefaultTransform = [16]float64{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0}
)

type Tileset struct {
	Asset              Asset              `json:"asset"`
	GeometricError     float64            `json:"geometricError"` //核心值，单位米，表示简化后模型与原始模型的最大空间距离，核心作用是为瓦片调度提供LOD切换依据。
	Root               *Tile              `json:"root"`
	Properties         *map[string]Schema `json:"properties,omitempty"`
	ExtensionsUsed     *[]string          `json:"extensionsUsed,omitempty"`
	ExtensionsRequired *[]string          `json:"extensionsRequired,omitempty"`
}

// ToJSON
func (m *Tileset) ToJSON() (string, error) {
	b, e := json.Marshal(m)
	return string(b), e
}
