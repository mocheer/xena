package tianditu_test

import (
	"testing"

	"github.com/mocheer/xena/pkg/gm"
	"github.com/mocheer/xena/pkg/provider/tianditu"
)

func TestXxx(t *testing.T) {
	tmap := tianditu.NewSatellite()
	t.Log(tmap.GetTileURL(&gm.Tile{}))
}
