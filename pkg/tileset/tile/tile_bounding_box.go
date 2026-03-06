package tile

import "math"

type BoundingBox [12]float64

// GetDiagonal
func (m *BoundingBox) GetHalfDiagonal() float64 {
	return math.Sqrt(m[3]*m[3] + m[7]*m[7] + m[11]*m[11])
}

// GetHalfDiagonalXY
func (m *BoundingBox) GetHalfDiagonalXY() float64 {
	return math.Sqrt(m[3]*m[3] + m[7]*m[7])
}
