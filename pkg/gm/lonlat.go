package gm

type LonLat [2]float64

// Lon 获取经度
func (m LonLat) Lon() float64 {
	return m[0]
}

// Lat 获取纬度
func (m LonLat) Lat() float64 {
	return m[1]
}

// ToPoint 获取点（将经纬度转换为Point对象）
func (m LonLat) ToPoint() Point {
	return Point(m)
}

// ConvertToPoint 将经纬度转换成平面坐标
func (m LonLat) ConvertToPoint() *Point {
	return &Point{}
}
