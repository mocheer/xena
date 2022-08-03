package proj4

import "fmt"

// 专用于UTM投影WGS84基准面的proj
type UTM_WGS84_ZONE int

// EPSGCode
// UTM_WGS84_ZONE(50).EPSGCode()
func (m UTM_WGS84_ZONE) EPSGCode() EPSGCode {
	code := EPSGCode(32600 + m)
	_, ok := projStrings[code]
	if !ok {
		Defs(int(code), fmt.Sprintf("+proj=utm +zone=%d +datum=WGS84 +units=m +no_defs", m))
	}
	return code
}

// Convert 经纬度转投影坐标系
func (m UTM_WGS84_ZONE) Convert(xy []float64) ([]float64, error) {
	return Convert(m.EPSGCode(), xy)
}

// Inverse 投影坐标系转经纬度
func (m UTM_WGS84_ZONE) Inverse(xy []float64) ([]float64, error) {
	return Inverse(m.EPSGCode(), xy)
}
