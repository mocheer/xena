package tileset

// 切片的边界体积
// 一般只用 region、box、sphere中的一种来表达这个边界范围
type BoundingVolume struct {
	// 表示地理空间区域的边界
	// [west, south, east, north, minimum height, maximum height]
	// @see https://cesium.com/learn/cesiumjs/ref-doc/AxisAlignedBoundingBox.html?classFilter=bounding
	Region *[6]float64 `json:"region,omitempty"`
	// 表示包围瓦片或瓦片集的球体。
	// Oriented Bounding Box
	// [ centerX,centerY,centerZ, halfXLength,dxy,dxz, dyx,halfYLength,dyz, dzx,dzy,halfZLength]
	// @see https://cesium.com/learn/cesiumjs/ref-doc/OrientedBoundingBox.html?classFilter=bound
	Box *[12]float64 `json:"box,omitempty"`
	// 表示瓦片或瓦片集的轴对齐边界框。
	// [centerX,centerY,centerZ,radius]
	Sphere *[4]float64 `json:"sphere,omitempty"`
}

// SetBox
func (b *BoundingVolume) SetBox(box [12]float64) {
	b.Region = nil
	b.Sphere = nil
	b.Box = &box
}

// SetRegion
func (b *BoundingVolume) SetRegion(region [6]float64) {
	b.Box = nil
	b.Sphere = nil
	b.Region = &region
}

// SetSphere
func (b *BoundingVolume) SetSphere(sphere [4]float64) {
	b.Box = nil
	b.Region = nil
	b.Sphere = &sphere
}

// GetRegion
func (m *BoundingVolume) GetRegion() *[6]float64 {
	if m.Region != nil {
		return m.Region
	}
	// 这里默认坐标轴对称
	if m.Box != nil {
		centerX, centerY, centerZ := m.Box[0], m.Box[1], m.Box[2]
		halfW, halfH, halfZ := m.Box[4], m.Box[8], m.Box[11]

		return &[6]float64{
			centerX - halfW*2,
			centerY - halfH*2,
			centerX + halfW*2,
			centerY + halfH*2,
			centerZ - halfZ*2,
			centerZ + halfZ*2,
		}
	}
	return nil
}

// 这里判断两个aabb是否重叠
func IsBoundingBoxesOverlap(box1, box2 *BoundingVolume) bool {
	r1 := box1.GetRegion()
	r2 := box2.GetRegion()

	return r1[0] > r2[2] || r1[2] < r2[0] || r1[1] > r2[3] || r1[3] < r2[1] || r1[4] > r2[5] || r1[5] < r2[4]
}

// QuadtreeBoundingVolume
// 获取当前BoundingVolumn的四叉树子级
// 这里中心点高度不变
func (m *BoundingVolume) QuadtreeBoundingVolumeBox() [4]BoundingVolume {
	centerX, centerY, centerZ := m.Box[0], m.Box[1], m.Box[2]
	halfW, halfH, halfZ := m.Box[4], m.Box[8], m.Box[11]
	return [4]BoundingVolume{
		{
			Box: &[12]float64{
				centerX - halfW,
				centerY - halfH,
				centerZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ,
			},
		},
		{
			Box: &[12]float64{
				centerX + halfW,
				centerY - halfH,
				centerZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ,
			},
		},
		{
			Box: &[12]float64{
				centerX - halfW,
				centerY + halfH,
				centerZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ,
			},
		},
		{
			Box: &[12]float64{
				centerX + halfW,
				centerY + halfH,
				centerZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ,
			},
		},
	}

}

// OctreeBoundingVolume 八叉树
func (m *BoundingVolume) OctreeBoundingVolumeBox() [8]BoundingVolume {
	centerX, centerY, centerZ := m.Box[0], m.Box[1], m.Box[2]
	halfW, halfH, halfZ := m.Box[4], m.Box[8], m.Box[11]
	return [8]BoundingVolume{
		{
			Box: &[12]float64{
				centerX - halfW,
				centerY - halfH,
				centerZ - halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
		{
			Box: &[12]float64{
				centerX + halfW,
				centerY - halfH,
				centerZ - halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
		{
			Box: &[12]float64{
				centerX - halfW,
				centerY + halfH,
				centerZ - halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
		{
			Box: &[12]float64{
				centerX - halfW,
				centerY - halfH,
				centerZ + halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
		{
			Box: &[12]float64{
				centerX + halfW,
				centerY + halfH,
				centerZ - halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
		{
			Box: &[12]float64{
				centerX + halfW,
				centerY - halfH,
				centerZ + halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
		{
			Box: &[12]float64{
				centerX - halfW,
				centerY + halfH,
				centerZ + halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
		{
			Box: &[12]float64{
				centerX + halfW,
				centerY + halfH,
				centerZ + halfZ,
				halfW / 2,
				0,
				0,
				0,
				halfH / 2,
				0,
				0,
				0,
				halfZ / 2,
			},
		},
	}

}
