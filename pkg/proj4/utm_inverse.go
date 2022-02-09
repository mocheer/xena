package proj4

import "fmt"

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

// ConvertToWGS84
func (m UTM_WGS84_ZONE) ConvertToWGS84(xy []float64) ([]float64, error) {
	return Inverse(m.EPSGCode(), xy)
}
