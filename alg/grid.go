package alg

// GetGridAround 获取网格点周围的8个方格
func GetGridAround(p [2]float64, size float64) ([2]float64, [2]float64, [2]float64, [2]float64, [2]float64, [2]float64, [2]float64, [2]float64) {
	return [2]float64{p[0] + size, p[1]}, [2]float64{p[0], p[1] + size}, [2]float64{p[0] + size, p[1] + size}, [2]float64{p[0] - size, p[1]}, [2]float64{p[0], p[1] - size}, [2]float64{p[0] - size, p[1] - size}, [2]float64{p[0] + size, p[1] - size}, [2]float64{p[0] - size, p[1] + size}
}
