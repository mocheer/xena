package geojson

import (
	"encoding/json"

	"github.com/samber/lo"
	"github.com/tidwall/gjson"
)

type GeoJSON json.RawMessage

// From
func From(data []byte) GeoJSON {
	return GeoJSON(data)
}

func (m GeoJSON) Type() string {
	return gjson.GetBytes(m, "type").Str
}

func (m GeoJSON) Features() []*Feature {
	t := m.Type()
	if t == "" {
		return nil
	}
	switch t {
	case TypeFeatureCollection:
		return m.AsFeatureCollection().Features
	case TypeGeometryCollection:
		return lo.Map(m.AsGeometryCollection().Geometries, func(geom *Geometry, _ int) *Feature {
			return NewFeature(geom)
		})
	case TypeFeature:
		return []*Feature{m.AsFeature()}
	}
	return nil
}

func (m GeoJSON) Geometries() []*Geometry {
	t := m.Type()
	if t == "" {
		return nil
	}
	switch t {
	case TypeFeatureCollection:
		return lo.Map(m.AsFeatureCollection().Features, func(f *Feature, _ int) *Geometry {
			return f.Geometry
		})
	case TypeGeometryCollection:
		return m.AsGeometryCollection().Geometries
	case TypeFeature:
		return []*Geometry{m.AsFeature().Geometry}
	}
	return nil
}

// GeometryCoordinates
func (m GeoJSON) GeometryCoordinates() []GeometryCoordinates {
	t := m.Type()
	if t == "" {
		return nil
	}
	switch t {
	case TypeFeatureCollection:
		return lo.Map(m.AsFeatureCollection().Features, func(f *Feature, _ int) GeometryCoordinates {
			return f.Geometry.Coordinates
		})
	case TypeGeometryCollection:
		return lo.Map(m.AsGeometryCollection().Geometries, func(f *Geometry, _ int) GeometryCoordinates {
			return f.Coordinates
		})
	case TypeFeature:
		return []GeometryCoordinates{m.AsFeature().Geometry.Coordinates}
	default:
		return []GeometryCoordinates{m.AsGeometry().Coordinates}
	}
}

// AsFeatureCollection
func (m GeoJSON) AsFeatureCollection() *FeatureCollection {
	fc := &FeatureCollection{}
	err := json.Unmarshal(m, fc)
	if err != nil {
		panic(err)
	}
	return fc
}

// AsGeometryCollection
func (m GeoJSON) AsGeometryCollection() *GeometryCollection {
	gc := &GeometryCollection{}
	err := json.Unmarshal(m, gc)
	if err != nil {
		panic(err)
	}
	return gc
}

// AsFeature
func (m GeoJSON) AsFeature() *Feature {
	gc := &Feature{}
	err := json.Unmarshal(m, gc)
	if err != nil {
		panic(err)
	}
	return gc
}

// AsFeature
func (m GeoJSON) AsGeometry() *Geometry {
	g := &Geometry{}
	err := json.Unmarshal(m, g)
	if err != nil {
		panic(err)
	}
	return g
}
