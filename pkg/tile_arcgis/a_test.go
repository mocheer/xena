package tile_arcgis_test

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_xml"
	"github.com/mocheer/xena/pkg/tile_arcgis"
)

func Test1(t *testing.T) {
	o := int64(binary.LittleEndian.Uint64([]byte{60, 0, 0, 0, 0}))
	fmt.Println(o)
}

func Test2(t *testing.T) {
	var cdi tile_arcgis.ArcgisTileCDI
	err := ds_xml.ReadFile("./testdata/3857/esriMapCacheStorageModeCompact/conf.cdi", &cdi)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(cdi.GetCenter())
}
