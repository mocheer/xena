package topojson

import (
	geojson "github.com/paulmach/go.geojson"
	topoj "github.com/rubenv/topojson"
)

func NewTopology(data []byte, opts *topoj.TopologyOptions) *topoj.Topology {
	fc, err := geojson.UnmarshalFeatureCollection(data)
	if err != nil {
		panic(err)
	}
	return topoj.NewTopology(fc, opts)
}
