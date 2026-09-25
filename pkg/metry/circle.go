package gm

type Circle [3]float64

// Center
func (m Circle) Center() [2]float64 {
	// var p [2]float64
	// copy(p[:], m[:2])
	return *(*[2]float64)(m[:2])
}

// Radius
func (m Circle) Radius() float64 {
	return m[2]
}

// SetCenter
func (m *Circle) SetCenter(xy [2]float64) {
	m[0] = xy[0]
	m[1] = xy[1]
}

// SetRadius
func (m *Circle) SetRadius(r float64) {
	m[2] = r
}
