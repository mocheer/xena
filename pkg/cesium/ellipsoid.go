package cesium

var (
	Wgs84OneOverRadii = Cartesian3{
		X: 1.0 / 6378137.0,
		Y: 1.0 / 6378137.0,
		Z: 1.0 / 6356752.3142451793,
	}
	Wgs84OneOverRadiiSquared = Cartesian3{
		X: 1.0 / (6378137.0 * 6378137.0),
		Y: 1.0 / (6378137.0 * 6378137.0),
		Z: 1.0 / (6356752.3142451793 * 6356752.3142451793),
	}
	Wgs84CenterToleranceSquared = Epsilon1
)

// https://cesium.com/learn/cesiumjs/ref-doc/Ellipsoid.html?classFilter=Ellipsoid
// https://github.com/CesiumGS/cesium/blob/1.131/packages/engine/Source/Core/Ellipsoid.js
type Ellipsoid struct {
	X float64
	Y float64
	Z float64
}
