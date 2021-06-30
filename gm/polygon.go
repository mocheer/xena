package gm

// Polygon 三维数组
type Polygon [][][2]float64

// BBox 获取多边形的边界范围
func (m Polygon) BBox() BBox {
	var maxX, maxY, minX, minY float64
	for _, lines := range m {
		for _, side := range lines {
			if side[0] > maxX || maxX == 0.0 {
				maxX = side[0]
			}
			if side[1] > maxY || maxY == 0.0 {
				maxY = side[1]
			}
			if side[0] < minX || minX == 0.0 {
				minX = side[0]
			}
			if side[1] < minY || minY == 0.0 {
				minY = side[1]
			}
		}
	}
	return BBox{minX, minY, maxX, maxY}
}

// ContainPoint 射线算法判断点是否在多边形内部
func (m Polygon) ContainPoint(pt [2]float64) bool {
	if !m.BBox().ContainPoint(pt) {
		return false
	}
	//
	for _, lines := range m {
		nverts := len(lines)
		intersect := false
		j := 0
		for i := 1; i < nverts; i++ {
			if ((lines[i][1] > pt[1]) != (lines[j][1] > pt[1])) &&
				(pt[0] < (lines[j][0]-lines[i][0])*(pt[1]-lines[i][1])/(lines[j][1]-lines[i][1])+lines[i][0]) {
				intersect = !intersect
			}
			j = i
		}
		if intersect {
			return true
		}
	}
	return false
}
