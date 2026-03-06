package pbf_test

import (
	"fmt"
	"testing"

	"github.com/mocheer/xena/pkg/pbf"
)

func TestTopojson(t *testing.T) {
	data := pbf.ReadOSM("D:\\data\\pbf\\osm\\jiangxi.osm.pbf")
	fmt.Println(len(data.Nodes))
}
