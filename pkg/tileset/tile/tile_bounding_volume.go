package tile

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
	Box *BoundingBox `json:"box,omitempty"`
	// 表示瓦片或瓦片集的轴对齐边界框。
	// [centerX,centerY,centerZ,radius]
	Sphere *[4]float64 `json:"sphere,omitempty"`
}

// SetBox
func (b *BoundingVolume) SetBox(box BoundingBox) {
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
		halfW, halfH, halfZ := m.Box[3], m.Box[7], m.Box[11]

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

// GetRegion
func (m *BoundingVolume) GetBox() *BoundingBox {
	if m.Box != nil {
		return m.Box
	}
	// 这里默认坐标轴对称
	if m.Region != nil {
		west, south, east, north, minZ, maxZ := m.Region[0], m.Region[1], m.Region[2], m.Region[3], m.Region[4], m.Region[5]
		return &BoundingBox{
			(east + west) / 2,
			(south + north) / 2,
			(minZ + maxZ) / 2,
			(east - west) / 2,
			0,
			0,
			0,
			(north - south) / 2,
			0,
			0,
			0,
			(maxZ - minZ) / 2,
		}
	}
	return nil
}

// QuadtreeBoundingVolume
// 获取当前BoundingVolumn的四叉树子级
// 这里中心点高度不变
// https://github.com/vladimirpajic/cesium_3d_tiles_generator/blob/master/src/quadtree.rs
func (m *BoundingVolume) QuadtreeBoundingVolumeBox() [4]BoundingVolume {
	centerX, centerY, centerZ := m.Box[0], m.Box[1], m.Box[2]
	halfW, halfH, halfZ := m.Box[4], m.Box[8], m.Box[11]
	return [4]BoundingVolume{
		{
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
			Box: &BoundingBox{
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
