package gd_test

import (
	"testing"

	gd "github.com/mocheer/xena/pkg/topojson"
	"github.com/rubenv/topojson"
)

func TestTopojson(t *testing.T) {

	in := []byte(`{"type":"FeatureCollection","features":[{"type":"Feature","properties":{"id":"road"},"geometry":{"type":"LineString","coordinates":[[4.707126617431641,50.88714360752515],[4.708242416381835,50.886683361842216],[4.708693027496338,50.886514152727386],[4.70914363861084,50.886321253586765],[4.709406495094299,50.88633140619302],[4.709567427635193,50.88636524819791],[4.709604978561401,50.88647354244835],[4.709545969963074,50.88664275171071],[4.708666205406189,50.88698116839158],[4.707368016242981,50.88743464288961]]}}]}`)

	topo := gd.NewTopology(in, &topojson.TopologyOptions{
		IDProperty:   "id",
		PreQuantize:  1000000,
		PostQuantize: 10000,
	})

	data, _ := topo.MarshalJSON()
	t.Log(string(data))
}
