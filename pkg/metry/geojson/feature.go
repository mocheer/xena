package geojson

import "github.com/mocheer/xena/pkg/gm"

// A Feature corresponds to GeoJSON feature object
type Feature struct {
	ID         *string    `json:"id,omitempty"`
	Type       string     `json:"type"`
	BBox       *gm.BBox   `json:"bbox,omitempty"`
	Geometry   *Geometry  `json:"geometry"`
	Properties Properties `json:"properties"`
}

// NewFeature creates and initializes a GeoJSON feature given the required attributes.
func NewFeature(geometry *Geometry) *Feature {
	return &Feature{
		Type:       TypeFeature,
		Geometry:   geometry,
		Properties: make(map[string]any),
	}
}
