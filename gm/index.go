package gm

import "github.com/mocheer/pluto/ds/ds_json"

func NewPolygonFromJSON(fileName string) Polygon {
	var p [][][2]float64
	ds_json.Read(fileName, &p)
	return Polygon(p)
}
