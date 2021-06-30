package proj4

import "github.com/go-spatial/proj"

//
func Inverse(x, y float64) ([]float64, error) {
	return proj.Inverse(proj.WorldMercator, []float64{x, y})
}
