/*
Package geojson is a library for encoding and decoding GeoJSON into Go structs
using the geometries in the orb package. Supports both the json.Marshaler and
json.Unmarshaler interfaces as well as helper functions such as
`UnmarshalFeatureCollection` and `UnmarshalFeature`.
*/
package geojson

import "github.com/mocheer/xena/pkg/gm"

// FeatureCollection
type FeatureCollection struct {
	Type     string     `json:"type"`
	BBox     *gm.BBox   `json:"bbox,omitempty"`
	Features []*Feature `json:"features"`
}

// NewFeatureCollection
func NewFeatureCollection() *FeatureCollection {
	return &FeatureCollection{
		Type:     TypeFeatureCollection,
		Features: []*Feature{},
	}
}

// Append
func (fc *FeatureCollection) Append(feature *Feature) *FeatureCollection {
	fc.Features = append(fc.Features, feature)
	return fc
}
