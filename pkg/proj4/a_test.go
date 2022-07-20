package proj4_test

import (
	"testing"

	"github.com/mocheer/xena/pkg/proj4"
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

// TestUTMToWGS84 将一个utm投影坐标转成wgs84的经纬度坐标
func TestUTMToWGS84_2(t *testing.T) {
	lonlats, err := proj4.UTM_WGS84_ZONE(50).ConvertToWGS84([]float64{300278.079510, 3806211.646500 + 1028.256247*100})
	if err != nil {
		t.Log(err)
	}
	// 0.011176382537186669
	// 114.82780571769993 34.378037440547914
	t.Log(lonlats)
	t.Error("")
}
