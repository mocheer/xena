package gm

import "math"

// Polygon 三维数组
type Polygon [][][2]float64

// BBox 获取多边形的边界范围
func (m Polygon) BBox() BBox {
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64
	for _, lines := range m {
		for _, point := range lines {
			if point[0] > maxX {
				maxX = point[0]
			}
			if point[1] > maxY {
				maxY = point[1]
			}
			if point[0] < minX {
				minX = point[0]
			}
			if point[1] < minY {
				minY = point[1]
			}
		}
	}
	return BBox{minX, minY, maxX, maxY}
}

// ContainsPoint 射线法判断点是否在多边形内部，这里会先判断是否在bbox中（在大部分场景中，这样做能提高性能）
func (m Polygon) ContainsPoint(pt [2]float64) bool {
	return m.BBox().ContainsPoint(pt) && m.ContainsPoint2(pt)
}

// ContainsPoint 射线法判断点是否在多边形内部
func (m Polygon) ContainsPoint2(point [2]float64) bool {
	for _, lines := range m {
		nverts := len(lines)
		intersect := false
		j := nverts - 1
		x, y := point[0], point[1]
		for i := 0; i < nverts; i++ {
			if ((lines[i][1] > y) != (lines[j][1] > y)) &&
				(x < (lines[j][0]-lines[i][0])*(y-lines[i][1])/(lines[j][1]-lines[i][1])+lines[i][0]) {
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

// Grids 多边形网格化，返回网格点集合
func (m Polygon) Grids(cellSize float64) [][2]float64 {
	bbox := m.BBox()
	minX, minY := bbox.MinX(), bbox.MinY()
	maxX, maxY := bbox.MaxX(), bbox.MaxY()
	data := [][2]float64{}
	for i := minX; i < maxX; i += cellSize {
		for j := minY; j < maxY; j += cellSize {
			p := [2]float64{i, j}
			if m.ContainsPoint2(p) {
				data = append(data, p)
			}
		}
	}
	return data
}
