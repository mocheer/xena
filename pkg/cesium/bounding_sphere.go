package cesium

// BoundingSphere
type BoundingSphere struct {
	Center [3]float64 `json:"center"`
	Radius float64    `json:"radius"`
}

func FromPoints(points [][3]float64) *BoundingSphere {
	return &BoundingSphere{}
}
