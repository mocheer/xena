package cesium

import (
	"fmt"
	"math"
)

// https://cesium.com/learn/cesiumjs/ref-doc/Cartographic.html?classFilter=Cartographic
// 这里是弧度制
type Cartographic struct {
	Longitude float64
	Latitude  float64
	Height    float64
}

// CartographicFromCartesian3
// https://cesium.com/learn/cesiumjs/ref-doc/Cartographic.html?classFilter=cartog#.fromCartesian
// https://github.com/CesiumGS/cesium/blob/1.131/packages/engine/Source/Core/Cartographic.js#L118
func CartographicFromCartesian3(cs *Cartesian3) (*Cartographic, error) {
	oneOverRadii := Wgs84OneOverRadii
	oneOverRadiiSquared := Wgs84OneOverRadiiSquared
	centerToleranceSquared := Wgs84CenterToleranceSquared

	p, err := scaleToGeodeticSurface(cs, &oneOverRadii, &oneOverRadiiSquared, centerToleranceSquared)

	if err != nil {
		return nil, fmt.Errorf("failed to do scaleToGeoticSurface transformation: %v", err)
	}

	n := multiplyCartesian3Components(p, &oneOverRadiiSquared)
	n = Normalize(n)

	h := subtract(cs, p)

	longitude := math.Atan2(n.Y, n.X)
	latitude := math.Asin(n.Z)
	height := sign(dot(h, cs)) * h.Magnitude()

	return &Cartographic{
		Longitude: longitude,
		Latitude:  latitude,
		Height:    height,
	}, nil

}
