package alg

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
