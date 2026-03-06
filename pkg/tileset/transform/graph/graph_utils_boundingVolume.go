package graph

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/mocheer/xena/pkg/cesium"
)

// ComputedBoundingShpere
func (m *Graph) ComputedBoundingShpere() cesium.BoundingSphere {
	//
	points := []mgl64.Vec3{}
	m.GetScene().EachNodes(func(node *GraphNode) {
		traverseNodeWithMatrix(node, mgl64.Ident4(), &points)
	})
	return computeRitterBoundingSphere(points)
}

// traverseNodeWithMatrix
func traverseNodeWithMatrix(node *GraphNode, parentMatrix mgl64.Mat4, points *[]mgl64.Vec3) {
	localMatrix := node.GetMatrix()
	worldMatrix := parentMatrix.Mul4(localMatrix)
	//
	mesh := node.GetMesh()
	if mesh != nil {
		for _, p := range mesh.GetPrimitives() {
			positions, err := p.ReadPostion()
			if err == nil {
				for _, p := range positions {
					worldVertex := applyTransform(worldMatrix, mgl64.Vec3{float64(p[0]), float64(p[1]), float64(p[2])})
					*points = append(*points, worldVertex)
				}
			}
		}
	}
	node.EachNodes(func(node *GraphNode) {
		traverseNodeWithMatrix(node, worldMatrix, points)
	})
}

// applyTransform 将4x4矩阵应用到3D向量
func applyTransform(matrix mgl64.Mat4, vertex mgl64.Vec3) mgl64.Vec3 {
	// 将3D向量转换为齐次坐标（w=1），应用变换，再转换回3D坐标
	v := matrix.Mul4x1(mgl64.Vec4{vertex[0], vertex[1], vertex[2], 1})
	return mgl64.Vec3{v[0] / v[3], v[1] / v[3], v[2] / v[3]}
}

// computeRitterBoundingSphere 使用Ritter算法计算点集的最小包围球
func computeRitterBoundingSphere(points []mgl64.Vec3) cesium.BoundingSphere {
	// 第一步：找到在X, Y, Z轴上分布最远的两个点作为初始直径
	var minPt, maxPt cesium.Cartesian3
	maxDistance := -1.0

	// 检查三个轴方向上的极值点对
	axes := []func(cesium.Cartesian3) float64{
		func(v cesium.Cartesian3) float64 { return v.X },
		func(v cesium.Cartesian3) float64 { return v.Y },
		func(v cesium.Cartesian3) float64 { return v.Z },
	}

	for _, axis := range axes {
		var minAxisPt, maxAxisPt cesium.Cartesian3
		minVal, maxVal := math.MaxFloat64, -math.MaxFloat64

		for _, p2 := range points {
			p := cesium.Cartesian3{X: p2[0], Y: p2[1], Z: p2[2]}
			val := axis(p)
			if val < minVal {
				minVal = val
				minAxisPt = p
			}
			if val > maxVal {
				maxVal = val
				maxAxisPt = p
			}
		}

		dist := minAxisPt.DistanceTo(maxAxisPt)
		if dist > maxDistance {
			maxDistance = dist
			minPt, maxPt = minAxisPt, maxAxisPt
		}
	}

	// 初始球体：以最远点对的中点为球心，半长为半径[citation:1]
	center := cesium.Cartesian3{
		X: (minPt.X + maxPt.X) / 2,
		Y: (minPt.Y + maxPt.Y) / 2,
		Z: (minPt.Z + maxPt.Z) / 2,
	}
	radius := maxDistance / 2

	// 第二步：遍历所有点，扩展球体以包含所有点[citation:1]
	for _, p2 := range points {
		p := cesium.Cartesian3{X: p2[0], Y: p2[1], Z: p2[2]}
		d := center.DistanceTo(p)
		if d > radius {
			// 点p在球外，需要扩展球体
			newRadius := (radius + d) / 2
			offset := (d - radius) / (2 * d) // 球心需要向点p方向移动的比例
			center.X += (p.X - center.X) * offset
			center.Y += (p.Y - center.Y) * offset
			center.Z += (p.Z - center.Z) * offset
			radius = newRadius
		}
	}

	return cesium.BoundingSphere{
		Center: [3]float64{center.X, center.Y, center.Z},
		Radius: radius,
	}
}
