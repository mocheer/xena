package gm

// 这里是弧度制
type Cartographic [3]float64

// ToPoint 获取点（将经纬度转换为Point对象）
func (m Cartographic) ToPoint() Point {
	return Point{m[0] * DEGREE_PER_RADIANS, m[1] * DEGREE_PER_RADIANS}
}
