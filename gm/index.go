package gm

import "github.com/mocheer/pluto/ds/dsjson"

func NewPolygonFromJSON(fileName string) Polygon {
	var p [][][2]float64
	dsjson.Read(fileName, &p)
	return Polygon(p)
}
