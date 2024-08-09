package pipelinie

import (
	"math"

	"github.com/mocheer/xena/pkg/tileset"
)

// GetGeometricError 计算几何误差
// 只需要计算 对角线/系数 ，这个系数约为16,即可，因为切片范围越大zoom越小精度越低，切片范围越小精度越高
// 对角线可能不太合理，因为有空白切片
func GetGeometricError(bbox tileset.BoundingVolume) float64 {
	// 计算bbox的对角线长度
	//

	// var w = math.Abs(bbox.Xmax - bbox.Xmin)
	// var l = math.Abs(bbox.Ymax - bbox.Ymin)
	// var h = math.Abs(bbox.Zmax - bbox.Zmin)
	// return math.Sqrt(w*w + l*l + h*h)

	// geometric error is estimated as the maximum possible distance between two points lying in the cell
	// return n.cellSize * math.Sqrt(3) * 2
	//
	return 0
}

// CalcEnuToEcefMatrix 计算ENU到ECEF坐标转换矩阵
func CalcEnuToEcefMatrix(lnt, lat, heightMin float64) [16]float64 {
	const pi = math.Pi
	ellipsodA := 40680631590769.0
	ellipsodB := 40680631590769.0
	ellipsodC := 40408299984661.4

	radianX := lnt * pi / 180.0
	radianY := lat * pi / 180.0
	xn := math.Cos(radianX) * math.Cos(radianY)
	yn := math.Sin(radianX) * math.Cos(radianY)
	zn := math.Sin(radianY)

	x0 := ellipsodA * xn
	y0 := ellipsodB * yn
	z0 := ellipsodC * zn
	gamma := math.Sqrt(xn*x0 + yn*y0 + zn*z0)
	px := x0 / gamma
	py := y0 / gamma
	pz := z0 / gamma

	dx := xn * heightMin
	dy := yn * heightMin
	dz := zn * heightMin

	eastMat := [3]float64{-y0, x0, 0}
	northMat := [3]float64{
		(y0*eastMat[2] - eastMat[1]*z0),
		(z0*eastMat[0] - eastMat[2]*x0),
		(x0*eastMat[1] - eastMat[0]*y0),
	}

	eastNormal := math.Sqrt(eastMat[0]*eastMat[0] + eastMat[1]*eastMat[1] + eastMat[2]*eastMat[2])
	northNormal := math.Sqrt(northMat[0]*northMat[0] + northMat[1]*northMat[1] + northMat[2]*northMat[2])

	return [16]float64{
		eastMat[0] / eastNormal,
		eastMat[1] / eastNormal,
		eastMat[2] / eastNormal,
		0,
		northMat[0] / northNormal,
		northMat[1] / northNormal,
		northMat[2] / northNormal,
		0,
		xn,
		yn,
		zn,
		0,
		px + dx,
		py + dy,
		pz + dz,
		1,
	}
}

// IsBoundingBoxesOverlap
// 这里判断两个aabb是否重叠
func IsBoundingBoxesOverlap(box1, box2 *tileset.BoundingVolume, tolerance float64) bool {
	r1 := box1.GetRegion()
	r2 := box2.GetRegion()
	check1 := (r1[0] - r2[2]) < tolerance // w1 < e2
	check2 := (r1[2] - r2[0]) > tolerance // e1 > w2
	check3 := (r1[1] - r2[3]) < tolerance // s1 < n2
	check4 := (r1[3] - r2[1]) > tolerance // n1 > s2
	check5 := (r1[4] - r2[5]) < tolerance // minZ1 < maxZ2
	check6 := (r1[5] - r2[4]) > tolerance // maxZ1 > minZ2

	return check1 && check2 && check3 && check4 && check5 && check6
}
