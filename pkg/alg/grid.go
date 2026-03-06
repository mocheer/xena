package alg

import "math"

// GetAroundGrids 获取网格点周围的8个方格
func GetAroundGrids(p [2]float64, size float64) [][2]float64 {
	return [][2]float64{
		{p[0] + size, p[1]},
		{p[0], p[1] + size},
		{p[0] + size, p[1] + size},
		{p[0] - size, p[1]},
		{p[0], p[1] - size},
		{p[0] - size, p[1] - size},
		{p[0] + size, p[1] - size},
		{p[0] - size, p[1] + size},
	}
}

// InterpolateGridData 将一份网格数据插值获取到一份分辨率更高的网格
func InterpolateGridData(gridData [][]float64, newWidth, newHeight int) [][]float64 {
	oldHeight, oldWidth := len(gridData), len(gridData[0])
	res := make([][]float64, newHeight)
	for i := 0; i < newHeight; i++ {
		res[i] = make([]float64, newWidth)
		for j := 0; j < newWidth; j++ {
			x := float64(j) / float64(newWidth-1) * float64(oldWidth-1)
			y := float64(i) / float64(newHeight-1) * float64(oldHeight-1)
			x0, y0 := int(math.Floor(x)), int(math.Floor(y))
			x1, y1 := int(math.Ceil(x)), int(math.Ceil(y))
			if x1 >= oldWidth {
				x1 = oldWidth - 1
			}
			if y1 >= oldHeight {
				y1 = oldHeight - 1
			}
			dx, dy := x-float64(x0), y-float64(y0)
			v00, v01, v10, v11 := gridData[y0][x0], gridData[y0][x1], gridData[y1][x0], gridData[y1][x1]
			res[i][j] = (1-dx)*(1-dy)*v00 + dx*(1-dy)*v01 + (1-dx)*dy*v10 + dx*dy*v11
		}
	}
	return res
}
