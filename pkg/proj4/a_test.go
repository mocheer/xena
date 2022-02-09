package proj4_test

import (
	"testing"

	"github.com/mocheer/xena/alg/proj4"
)

// TestUTMToWGS84 将一个utm投影坐标转成wgs84的经纬度坐标
func TestUTMToWGS84(t *testing.T) {
	lonlats, err := proj4.UTM_WGS84_ZONE(54).ConvertToWGS84([]float64{374424.86565261666, 3597865.9055234734})
	if err != nil {
		t.Log(err)
	}
	// 139.663118125,32.5109995234
	// height 42.774175278013054
	t.Log(lonlats)
}
