package asc_grid

import (
	_ "embed"
	"regexp"

	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/pluto/pkg/series/d3/d3_contour"
	"github.com/mocheer/xena/pkg/proj4"
	"github.com/samber/lo"

	"github.com/mocheer/pluto/pkg/ds/ds_text"
)

// AsciiGrid
type AsciiGrid struct {
	Ncols       int
	Nrows       int
	Xllcorner   float64
	Yllcorner   float64
	Cellsize    float64
	NodataValue float64
	Data        [][]float64
}

// Read
func Read(fileName string) *AsciiGrid {
	s, err := ds_text.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return ReadText(s)
}

// ReadText
func ReadText(s string) *AsciiGrid {
	ag := &AsciiGrid{}
	splitRegexp := regexp.MustCompile("\\s+")
	lo.ForEach(fn.SplitLines(s), func(line string, _ int) {
		kv := splitRegexp.Split(line, -1)
		switch kv[0] {
		case "ncols":
			ag.Ncols = fn.ParseInt(kv[1])
		case "nrows":
			ag.Nrows = fn.ParseInt(kv[1])
			ag.Data = make([][]float64, 0, ag.Nrows)
		case "xllcorner":
			ag.Xllcorner = fn.ParseFloat64(kv[1])
		case "yllcorner":
			ag.Yllcorner = fn.ParseFloat64(kv[1])
		case "cellsize":
			ag.Cellsize = fn.ParseFloat64(kv[1])
		case "NODATA_value":
			ag.NodataValue = fn.ParseFloat64(kv[1])
		default:
			arr := make([]float64, len(kv))
			for i, str := range kv {
				arr[i] = fn.ParseFloat64(str)
			}
			ag.Data = append(ag.Data, arr)
		}
	})

	return ag
}

// Contour
func (m AsciiGrid) Contour(legends []float64) (data []*d3_contour.ContourPolygon) {
	return d3_contour.Contour().Size([]int{m.Ncols, m.Nrows}).Thresholds(legends).Contours(lo.Flatten(m.Data))[1:]
}

// ToUTMGeoJSON
func (m AsciiGrid) ToUTMGeoJSON(legends []float64, zone int, precision int) (data []*d3_contour.ContourPolygon) {
	data = m.Contour(legends)
	proj := proj4.UTM_WGS84_ZONE(zone).Inverse
	minX := m.Xllcorner
	maxY := m.Yllcorner + float64(m.Nrows)*m.Cellsize
	lo.ForEach(data, func(polygon *d3_contour.ContourPolygon, _ int) {
		polygon.Coordinates = lo.Map(polygon.Coordinates, func(coor3 [][][2]float64, _ int) [][][2]float64 {
			return lo.Map(coor3, func(coor2 [][2]float64, _ int) [][2]float64 {
				return lo.Map(coor2, func(coor [2]float64, _ int) [2]float64 {
					p, _ := proj([]float64{coor[0]*m.Cellsize + minX, maxY - coor[1]*m.Cellsize})
					c1 := p[0]
					c2 := p[1]
					if precision > 0 {
						c1, c2 = fn.Round(c1, precision), fn.Round(c2, precision)
					}
					return [2]float64{c1, c2}
				})
			})
		})
	})
	return
}

// ToGeoJSON
func (m AsciiGrid) ToGeoJSON(legends []float64, precision int) (data []*d3_contour.ContourPolygon) {
	data = m.Contour(legends)
	minX := m.Xllcorner
	maxY := m.Yllcorner + float64(m.Nrows)*m.Cellsize
	lo.ForEach(data, func(polygon *d3_contour.ContourPolygon, _ int) {
		polygon.Coordinates = lo.Map(polygon.Coordinates, func(coor3 [][][2]float64, _ int) [][][2]float64 {
			return lo.Map(coor3, func(coor2 [][2]float64, _ int) [][2]float64 {
				return lo.Map(coor2, func(coor [2]float64, _ int) [2]float64 {
					c1 := coor[0]*m.Cellsize + minX
					c2 := maxY - coor[1]*m.Cellsize
					if precision > 0 {
						c1, c2 = fn.Round(c1, precision), fn.Round(c2, precision)
					}
					return [2]float64{c1, c2}
				})
			})
		})
	})
	return
}

// func (m AsciiGrid) Union(legends []float64) []polygol.Geom {
// 	result := make([][][2]int, len(legends)+1, len(legends)+1)
// 	for y, row := range m.Data {
// 		for x, val := range row {
// 			if val != m.NodataValue {
// 				index := getIndexLegend(legends, val)
// 				p1 := [2]int{x, y}
// 				result[index] = append(result[index], p1)
// 			}
// 		}
// 	}
// 	result2 := make([]polygol.Geom, len(result), len(result))

// 	for i, p := range result {
// 		if len(p) > 0 {
// 			var err error
// 			geoms := lo.Map(p, func(p [2]int, index int) polygol.Geom {
// 				x := float64(p[0])
// 				y := float64(p[1])

// 				return polygol.Geom{
// 					{
// 						{
// 							[]float64{x, y},
// 							[]float64{x + 1, y},
// 							[]float64{x + 1, y + 1},
// 							[]float64{x, y + 1},
// 						},
// 					},
// 				}
// 			})
// 			result2[i], err = polygol.Union(geoms[0], geoms[1:]...)
// 			if err != nil {
// 				panic(err)
// 			}
// 		}
// 	}
// 	return result2

// }

func getIndexLegend(legends []float64, val float64) int {
	for i, v := range legends {
		if val <= v {
			return i
		}
	}
	return len(legends) - 1
}
