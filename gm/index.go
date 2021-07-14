package gm

import "github.com/mocheer/pluto/fs"

func NewPolygonFromJSON(fileName string) Polygon {
	var p [][][2]float64
	fs.ReadJSON(fileName, &p)
	return Polygon(p)
}
