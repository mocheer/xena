package gd

import (
	geojson "github.com/paulmach/go.geojson"
	"github.com/rubenv/topojson"
)

func NewTopology(fc *geojson.FeatureCollection) *topojson.Topology {
	return topojson.NewTopology(fc, nil)
}
