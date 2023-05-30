package asc_grid_test

import (
	"testing"

	"github.com/mocheer/pluto/pkg/ds/ds_json"
	"github.com/mocheer/xena/pkg/asc_grid"
)

func TestAsciiGrid(t *testing.T) {
	ag := asc_grid.Read("./testdata/raincolor2022-r4000-c4000-d1-water_depth-shandong-000.asc")
	legends := []float64{0.5, 0.8, 1.2, 1.5}
	ag.ToGeoJSON(legends, 50)
}

func TestAsciiGrid2(t *testing.T) {
	ag := asc_grid.Read("./testdata/raincolor2022-r4000-c4000-d1-discharge-shandong-000.asc")
	legends := []float64{0.5, 1, 2}
	ag.ToGeoJSON(legends, 50)
}

func TestAsciiGrid3(t *testing.T) {
	ag := asc_grid.Read("./testdata/raincolor2022-r4000-c4000-d1-water_depth-shandong-000.asc")
	ds_json.Save(ag, "grid.json")
}

func TestAsciiGrid4(t *testing.T) {
	ag := asc_grid.Read("./testdata/raincolor2022-r4000-c4000-d1-discharge-shandong-000.asc")
	legends := []float64{0.5, 0.8, 1.2, 1.5}
	t.Log(ag.Cellsize)
	ds_json.Save(ag.ToGeoJSON(legends, 50), "grid2.json")
}
