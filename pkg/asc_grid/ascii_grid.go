package asc_grid

import (
	_ "embed"
	"regexp"

	"github.com/mocheer/pluto/pkg/d3/d3_contour"
	"github.com/mocheer/pluto/pkg/fn"
	"github.com/mocheer/xena/pkg/proj4"
	"github.com/samber/lo"

	"github.com/mocheer/pluto/pkg/ds/ds_text"
)

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
		return nil
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

// ToGeoJSON
func (m AsciiGrid) ToGeoJSON(legends []float64, zone int) (data []*d3_contour.ContourPolygon) {
	data = d3_contour.Contour().Size([]int{m.Ncols, m.Nrows}).Thresholds(legends).Contours(lo.Flatten(m.Data))[1:]

	if zone > 0 {
		proj := proj4.UTM_WGS84_ZONE(zone).Inverse
		lo.ForEach(data, func(polygon *d3_contour.ContourPolygon, _ int) {
			polygon.Coordinates = lo.Map(polygon.Coordinates, func(coor3 [][][2]float64, _ int) [][][2]float64 {
				return lo.Map(coor3, func(coor2 [][2]float64, _ int) [][2]float64 {
					return lo.Map(coor2, func(coor [2]float64, _ int) [2]float64 {
						p, _ := proj([]float64{coor[0]*m.Cellsize + m.Xllcorner, m.Yllcorner + float64(m.Nrows)*m.Cellsize - coor[1]*m.Cellsize})
						return [2]float64{fn.Round(p[0], 5), fn.Round(p[1], 5)}
					})
				})
			})
		})
	}

	return
}
