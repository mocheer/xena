package tileset

import (
	"strings"
)

// Content
// url 一般是相对的，相对于引用的相对于引用的tileset文件
// url 可以是一个新的tileset.json，也可以是b3dm、gltf格式的模型文件
type Content struct {
	Uri            string          `json:"uri"`           // >=1.0版本时，使用uri
	Url            *string         `json:"url,omitempty"` // 弃用，<1.0版本时，使用url
	BoundingVolume *BoundingVolume `json:"boundingVolume,omitempty"`
	Extentions     *map[string]any `json:"extension,omitempty"`
	Extras         *interface{}    `json:"extras,omitempty"`
}

// type TileData interface {
// 	[]byte | *Tileset
// }

// GetURL
func (m *Content) GetURL() string {
	url := m.Uri
	if url == "" {
		url = *m.Url
	}
	return url
}

// IsTileset
func (m *Content) IsTileset() bool {
	return strings.HasSuffix(m.GetURL(), ".json")
}
