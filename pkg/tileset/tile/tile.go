package tile

type Tile struct {
	Content             *Content        `json:"content,omitempty"`
	BoundingVolume      BoundingVolume  `json:"boundingVolume,omitempty"`
	ViewerRequestVolume *BoundingVolume `json:"viewerRequestVolume,omitempty"`
	GeometricError      float64         `json:"geometricError"`
	Refine              string          `json:"refine"`
	Transform           *[16]float64    `json:"transform,omitempty"`
	Children            []*Tile         `json:"children,omitempty"`
}

// CaclGeometricError
// 计算几何误差，遵循层级递减原则，确保父级误差 ≥ 子级误差之和；
// 顶点坐标的量化误差需控制在GeometricError的 1/5～1/10 以内
// 有些算法是对角线一半的0.05，有些算法是长宽高中最大值的0.05
func (m *Tile) CaclGeometricError() float64 {
	box := m.BoundingVolume.GetBox()
	if box != nil {
		return 0.05 * box.GetHalfDiagonalXY()
	}
	return 0
}

// SetDefaultRefine
func (m *Tile) SetDefaultRefine() {
	m.Refine = TILE_REFINE_REPLACE
}
