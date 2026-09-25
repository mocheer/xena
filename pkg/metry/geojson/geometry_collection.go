package geojson

// GeometryCollection
type GeometryCollection struct {
	Type       string      `json:"type"`
	Geometries []*Geometry `json:"geometries"`
}

func NewGeometryCollection() *GeometryCollection {
	return &GeometryCollection{
		Type:       TypeGeometryCollection,
		Geometries: []*Geometry{},
	}
}

// Append
func (m *GeometryCollection) Append(g *Geometry) *GeometryCollection {
	m.Geometries = append(m.Geometries, g)
	return m
}
