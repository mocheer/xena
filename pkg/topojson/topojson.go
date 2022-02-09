package gd

import (
	geojson "github.com/paulmach/go.geojson"
	"github.com/rubenv/topojson"
)

func NewTopology(data []byte, opts *topojson.TopologyOptions) *topojson.Topology {
	fc, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		panic(err)
	}
	return topojson.NewTopology(fc, opts)
}
