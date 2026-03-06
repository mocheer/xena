package cesium

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

// @see https://cesium.com/learn/cesiumjs/ref-doc/Cartesian3.html?classFilter=Cartesian3

type Cartesian3 struct {
	X float64
	Y float64
	Z float64
}

// MagnitudeSquared
// https://cesium.com/learn/cesiumjs/ref-doc/Cartesian3.html?classFilter=Cartesian3#.MagnitudeSquared
// 计算给定的笛卡儿坐标XYZ的平方。
func (cartesian *Cartesian3) MagnitudeSquared() float64 {
	return (cartesian.X*cartesian.X +
		cartesian.Y*cartesian.Y +
		cartesian.Z*cartesian.Z)
}

// Magnitude
// https://cesium.com/learn/cesiumjs/ref-doc/Cartesian3.html?classFilter=Cartesian3#.Magnitude
// 计算笛卡儿坐标的长度
func (cartesian *Cartesian3) Magnitude() float64 {
	return math.Sqrt(cartesian.MagnitudeSquared())
}

// Normalize
// https://cesium.com/learn/cesiumjs/ref-doc/Cartesian3.html?classFilter=Cartesian3#.normalize
// 归一化笛卡尔坐标
func Normalize(cartesian *Cartesian3) *Cartesian3 {
	res := Cartesian3{}
	magnitude := cartesian.Magnitude()

	res.X = cartesian.X / magnitude
	res.Y = cartesian.Y / magnitude
	res.Z = cartesian.Z / magnitude

	return &res
}

func cartesianFromSlice(sl []float64) *Cartesian3 {
	cs := Cartesian3{}
	cs.X = float64(sl[0])
	cs.Y = sl[1]
	cs.Z = sl[2]
	return &cs
}

// 计算两个笛卡尔坐标的分量差。
func subtract(left *Cartesian3, right *Cartesian3) *Cartesian3 {
	res := Cartesian3{}

	res.X = left.X - right.X
	res.Y = left.Y - right.Y
	res.Z = left.Z - right.Z

	return &res
}

// 计算两个笛卡儿坐标的点积
func dot(left *Cartesian3, right *Cartesian3) float64 {
	return left.X*right.X + left.Y*right.Y + left.Z*right.Z
}

func sign(value float64) float64 {
	if value == 0 {
		return 0
	}

	if value > 0 {
		return 1
	}

	return 0
}

func multiplyMat4ByPoint(matrix *mat.Dense, cartesian *Cartesian3) *Cartesian3 {
	res := Cartesian3{}

	vX := cartesian.X
	vY := cartesian.Y
	vZ := cartesian.Z

	x := matrix.At(0, 0)*vX + matrix.At(1, 0)*vY + matrix.At(2, 0)*vZ + matrix.At(3, 0)
	y := matrix.At(0, 1)*vX + matrix.At(1, 1)*vY + matrix.At(2, 1)*vZ + matrix.At(3, 1)
	z := matrix.At(0, 2)*vX + matrix.At(1, 2)*vY + matrix.At(2, 2)*vZ + matrix.At(3, 2)

	res.X = x
	res.Y = y
	res.Z = z
	return &res
}

// 将提供的笛卡尔坐标分量乘以提供的标量，返回新的笛卡尔坐标对象
func multiplyByScalar(cartesian *Cartesian3, scalar float64) *Cartesian3 {
	res := Cartesian3{}
	res.X = cartesian.X * scalar
	res.Y = cartesian.Y * scalar
	res.Z = cartesian.Z * scalar

	return &res
}

func multiplyCartesian3Components(left *Cartesian3, right *Cartesian3) *Cartesian3 {
	res := Cartesian3{}

	res.X = left.X * right.X
	res.Y = left.Y * right.Y
	res.Z = left.Z * right.Z

	return &res
}

var (
	ScaleToGeodeticSurfaceGradient = &Cartesian3{}
)

func scaleToGeodeticSurface(cartesian *Cartesian3, oneOverRadii *Cartesian3, oneOverRadiiSquared *Cartesian3, centerToleranceSquared float64) (*Cartesian3, error) {

	if cartesian == nil {
		return nil, fmt.Errorf("cartesian is required")
	}

	positionX := cartesian.X
	positionY := cartesian.Y
	positionZ := cartesian.Z

	oneOverRadiiX := oneOverRadii.X
	oneOverRadiiY := oneOverRadii.Y
	oneOverRadiiZ := oneOverRadii.Z

	x2 := positionX * positionX * oneOverRadiiX * oneOverRadiiX
	y2 := positionY * positionY * oneOverRadiiY * oneOverRadiiY
	z2 := positionZ * positionZ * oneOverRadiiZ * oneOverRadiiZ

	// Compute the squared ellipsoid norm.
	squaredNorm := x2 + y2 + z2
	ratio := math.Sqrt(1.0 / squaredNorm)

	// As an initial approximation, assume that the radial intersection is the projection point.
	intersection := multiplyByScalar(cartesian, ratio)

	// If the position is near the center, the iteration will not converge.
	if squaredNorm < centerToleranceSquared {
		if !math.IsNaN(ratio) {
			return intersection, nil
		} else {
			return nil, fmt.Errorf("RATIO IS NOT A NUMBER")
		}
	}

	oneOverRadiiSquaredX := oneOverRadiiSquared.X
	oneOverRadiiSquaredY := oneOverRadiiSquared.Y
	oneOverRadiiSquaredZ := oneOverRadiiSquared.Z

	// Use the gradient at the intersection point in place of the true unit normal.
	// The difference in magnitude will be absorbed in the multiplier.
	gradient := ScaleToGeodeticSurfaceGradient
	gradient.X = intersection.X * oneOverRadiiSquaredX * 2.0
	gradient.Y = intersection.Y * oneOverRadiiSquaredY * 2.0
	gradient.Z = intersection.Z * oneOverRadiiSquaredZ * 2.0

	// Compute the initial guess at the normal vector multiplier, lambda.
	lambda := ((1.0 - ratio) * cartesian.Magnitude()) / (0.5 * gradient.Magnitude())
	correction := 0.0

	var funcX float64
	var denominator float64
	var xMultiplier float64
	var yMultiplier float64
	var zMultiplier float64
	var xMultiplier2 float64
	var yMultiplier2 float64
	var zMultiplier2 float64
	var xMultiplier3 float64
	var yMultiplier3 float64
	var zMultiplier3 float64

	for ok := true; ok; ok = (math.Abs(funcX) > Epsilon12) {
		lambda -= correction

		xMultiplier = 1.0 / (1.0 + lambda*oneOverRadiiSquaredX)
		yMultiplier = 1.0 / (1.0 + lambda*oneOverRadiiSquaredY)
		zMultiplier = 1.0 / (1.0 + lambda*oneOverRadiiSquaredZ)

		xMultiplier2 = xMultiplier * xMultiplier
		yMultiplier2 = yMultiplier * yMultiplier
		zMultiplier2 = zMultiplier * zMultiplier

		xMultiplier3 = xMultiplier2 * xMultiplier
		yMultiplier3 = yMultiplier2 * yMultiplier
		zMultiplier3 = zMultiplier2 * zMultiplier

		funcX = x2*xMultiplier2 + y2*yMultiplier2 + z2*zMultiplier2 - 1.0

		// "denominator" here refers to the use of this expression in the velocity and acceleration
		// computations in the sections to follow.
		denominator =
			x2*xMultiplier3*oneOverRadiiSquaredX +
				y2*yMultiplier3*oneOverRadiiSquaredY +
				z2*zMultiplier3*oneOverRadiiSquaredZ

		derivative := -2.0 * denominator

		correction = funcX / derivative
	}

	res := Cartesian3{}
	res.X = positionX * xMultiplier
	res.Y = positionY * yMultiplier
	res.Z = positionZ * zMultiplier

	return &res, nil
}

// 计算两个点之间的欧氏距离
func (v Cartesian3) DistanceTo(other Cartesian3) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
