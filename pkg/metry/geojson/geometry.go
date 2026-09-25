package geojson

import (
	"encoding/json"

	gm "github.com/mocheer/xena/pkg/metry"
	"github.com/tidwall/gjson"
)

type GeometryCoordinates interface {
	// Center() gm.Point
}

// Geometry
type Geometry struct {
	Type        string              `json:"type"`
	Coordinates GeometryCoordinates `json:"coordinates"`
	Properties  Properties          `json:"properties"`
}

// 实现UnmarshalJSON方法来自定义 Geometry 的解析
func (m *Geometry) UnmarshalJSON(data []byte) error {
	b := gjson.ParseBytes(data)
	t := b.Get("type").Str
	var g any
	switch t {
	case TypeGeomPoint:
		g = new(gm.Point)
	case TypeGeomLineString:
		g = new(gm.LineString)
	case TypeGeomMultiLineString:
		g = new(gm.MultiLineString)
	case TypeGeomPolygon:
		g = new(gm.Polygon)
	case TypeGeomMultiPolygon:
		g = new(gm.MultiPolygon)
	}
	if g != nil {
		cdata := b.Get("coordinates").Raw
		err := json.Unmarshal([]byte(cdata), g)
		if err != nil {
			return err
		}
		m.Type = t
		m.Coordinates = g
	} else {
		panic("解析geometry失败,类型不存在")
	}
	return nil
}

// AsPoint
func (m *Geometry) AsPoint() *gm.Point {
	return m.Coordinates.(*gm.Point)
}

// AsLineString
func (m *Geometry) AsLineString() *gm.LineString {
	return m.Coordinates.(*gm.LineString)
}

// AsMultiLineString
func (m *Geometry) AsMultiLineString() *gm.MultiLineString {
	return m.Coordinates.(*gm.MultiLineString)
}

// AsPolygon
func (m *Geometry) AsPolygon() *gm.Polygon {
	return m.Coordinates.(*gm.Polygon)
}

// AsMultiPolygon
func (m *Geometry) AsMultiPolygon() *gm.MultiPolygon {
	return m.Coordinates.(*gm.MultiPolygon)
}

// Feature
func (m *Geometry) Feature() *Feature {
	return NewFeature(m)
}

// FeatureCollection
func (m *Geometry) FeatureCollection() *FeatureCollection {
	fc := NewFeatureCollection()
	fc.Features = append(fc.Features, m.Feature())
	return fc
}

// NewPointGeom
func NewPointGeom(geom *gm.Point) *Geometry {
	return &Geometry{
		Type:        TypeGeomPoint,
		Coordinates: geom,
	}
}

// NewLineStringGeom
func NewLineStringGeom(geom *gm.LineString) *Geometry {
	return &Geometry{
		Type:        TypeGeomLineString,
		Coordinates: geom,
	}
}

// NewLineStringGeometry
func NewPolygonGeom(geom *gm.Polygon) *Geometry {
	return &Geometry{
		Type:        TypeGeomPolygon,
		Coordinates: geom,
	}
}

// NewMultiLineStringGeom
func NewMultiLineStringGeom(geom *gm.MultiLineString) *Geometry {
	return &Geometry{
		Type:        TypeGeomMultiLineString,
		Coordinates: geom,
	}
}

// NewMultiPolygonGeom
func NewMultiPolygonGeom(geom *gm.MultiPolygon) *Geometry {
	return &Geometry{
		Type:        TypeGeomMultiPolygon,
		Coordinates: geom,
	}
}
