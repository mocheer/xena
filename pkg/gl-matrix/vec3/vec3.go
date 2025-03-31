package vec3

import (
	"math"

	"github.com/mocheer/xena/pkg/gl-matrix/cst"
)

// https://github.com/toji/gl-matrix/blob/master/src/vec3.js

// Vec3
type Vec3 [3]float64

// Clone 创建一个新的 Vec3，并使用给定向量的值初始化。
func Clone(a Vec3) Vec3 {
	return Vec3{a[0], a[1], a[2]}
}

// Length  返回向量的长度。
func Length(a Vec3) float64 {
	return math.Sqrt(a[0]*a[0] + a[1]*a[1] + a[2]*a[2])
}

// Copy 将一个 Vec3 的值复制到另一个 Vec3。
func Copy(out, a Vec3) Vec3 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	return out
}

// Set 将 Vec3 的分量设置为给定的值。
func Set(out Vec3, x, y, z float64) Vec3 {
	out[0] = x
	out[1] = y
	out[2] = z
	return out
}

// Add 将两个 Vec3 相加。
func Add(out, a, b Vec3) Vec3 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	return out
}

// Subtract 从向量 a 中减去向量 b。
func Subtract(out, a, b Vec3) Vec3 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	return out
}

// Multiply 将两个 Vec3 相乘。
func Multiply(a, b Vec3) Vec3 {
	return Vec3{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

// Divide 将向量 a 除以向量 b。
func Divide(out, a, b Vec3) Vec3 {
	out[0] = a[0] / b[0]
	out[1] = a[1] / b[1]
	out[2] = a[2] / b[2]
	return out
}

// Scale  将 Vec3 按标量值缩放。
func Scale(out, a Vec3, b float64) Vec3 {
	out[0] = a[0] * b
	out[1] = a[1] * b
	out[2] = a[2] * b
	return out
}

// ScaleAndAdd 将第二个向量缩放后与第一个向量相加。
// out: 接收结果的向量
// a: 第一个向量
// b: 第二个向量
// scale: 缩放系数
// 返回值: 结果向量 out
func ScaleAndAdd(out, a, b Vec3, scale float64) Vec3 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	return out
}

// Distance 计算两个 Vec3 之间的距离。
func Distance(a, b Vec3) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	z := b[2] - a[2]
	return math.Sqrt(x*x + y*y + z*z)
}

// SquaredLength  返回 Vec3 的长度的平方。
func SquaredLength(a Vec3) float64 {
	return a[0]*a[0] + a[1]*a[1] + a[2]*a[2]
}

// Negate 对 Vec3 的分量取反。
func Negate(out, a Vec3) Vec3 {
	out[0] = -a[0]
	out[1] = -a[1]
	out[2] = -a[2]
	return out
}

// Inverse 返回 Vec3 分量的倒数。
func Inverse(out, a Vec3) Vec3 {
	out[0] = 1.0 / a[0]
	out[1] = 1.0 / a[1]
	out[2] = 1.0 / a[2]
	return out
}

// Normalize 对 Vec3 进行归一化。
func Normalize(out, a Vec3) Vec3 {
	len := Length(a)
	if len > 0 {
		out[0] = a[0] / len
		out[1] = a[1] / len
		out[2] = a[2] / len
	}
	return out
}

// Dot 计算两个 Vec3 的点积。
func Dot(a, b Vec3) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// Cross 计算两个 Vec3 的叉积。
func Cross(out, a, b Vec3) Vec3 {
	out[0] = a[1]*b[2] - a[2]*b[1]
	out[1] = a[2]*b[0] - a[0]*b[2]
	out[2] = a[0]*b[1] - a[1]*b[0]
	return out
}

// Lerp 在两个 Vec3 之间进行线性插值。
func Lerp(out, a, b Vec3, t float64) Vec3 {
	out[0] = a[0] + t*(b[0]-a[0])
	out[1] = a[1] + t*(b[1]-a[1])
	out[2] = a[2] + t*(b[2]-a[2])
	return out
}

// Slerp 执行球面线性插值。
// out: 接收插值结果的向量
// a: 第一个向量
// b: 第二个向量
// t: 插值系数，范围 [0, 1]
// 返回值: 插值后的向量 out
func Slerp(out, a, b Vec3, t float64) Vec3 {
	// 计算两个向量之间的夹角
	angle := math.Acos(math.Min(math.Max(Dot(a, b), -1), 1))
	sinTotal := math.Sin(angle)

	// 避免除以零
	if sinTotal == 0 {
		out[0] = a[0]
		out[1] = a[1]
		out[2] = a[2]
		return out
	}

	// 计算插值系数
	ratioA := math.Sin((1-t)*angle) / sinTotal
	ratioB := math.Sin(t*angle) / sinTotal

	// 计算插值结果
	out[0] = ratioA*a[0] + ratioB*b[0]
	out[1] = ratioA*a[1] + ratioB*b[1]
	out[2] = ratioA*a[2] + ratioB*b[2]

	return out
}

// Hermite 执行 Hermite 插值。
// out: 接收插值结果的向量
// a: 第一个向量
// b: 第二个向量
// c: 第三个向量
// d: 第四个向量
// t: 插值系数，范围 [0, 1]
// 返回值: 插值后的向量 out
func Hermite(out, a, b, c, d Vec3, t float64) Vec3 {
	factorTimes2 := t * t
	factor1 := factorTimes2*(2*t-3) + 1
	factor2 := factorTimes2*(t-2) + t
	factor3 := factorTimes2 * (t - 1)
	factor4 := factorTimes2 * (3 - 2*t)

	// 计算插值结果
	out[0] = a[0]*factor1 + b[0]*factor2 + c[0]*factor3 + d[0]*factor4
	out[1] = a[1]*factor1 + b[1]*factor2 + c[1]*factor3 + d[1]*factor4
	out[2] = a[2]*factor1 + b[2]*factor2 + c[2]*factor3 + d[2]*factor4

	return out
}

// Bezier 执行贝塞尔插值。
// out: 接收插值结果的向量
// a: 第一个向量
// b: 第二个向量
// c: 第三个向量
// d: 第四个向量
// t: 插值系数，范围 [0, 1]
// 返回值: 插值后的向量 out
func Bezier(out, a, b, c, d Vec3, t float64) Vec3 {
	inverseFactor := 1 - t
	inverseFactorTimesTwo := inverseFactor * inverseFactor
	factorTimes2 := t * t
	factor1 := inverseFactorTimesTwo * inverseFactor
	factor2 := 3 * t * inverseFactorTimesTwo
	factor3 := 3 * factorTimes2 * inverseFactor
	factor4 := factorTimes2 * t

	// 计算插值结果
	out[0] = a[0]*factor1 + b[0]*factor2 + c[0]*factor3 + d[0]*factor4
	out[1] = a[1]*factor1 + b[1]*factor2 + c[1]*factor3 + d[1]*factor4
	out[2] = a[2]*factor1 + b[2]*factor2 + c[2]*factor3 + d[2]*factor4

	return out
}

// Random  生成一个具有给定比例的随机向量。
func Random(out Vec3, scale float64) Vec3 {
	// 占位符，用于随机数生成
	r := 1.0 //   替换为实际的随机数生成逻辑
	out[0] = r * scale
	out[1] = r * scale
	out[2] = r * scale
	return out
}

// TransformMat4 使用 4x4 矩阵变换 Vec3。
func TransformMat4(out, a Vec3, m [16]float64) Vec3 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := m[3]*x + m[7]*y + m[11]*z + m[15]
	w = 1.0 / w
	out[0] = (m[0]*x + m[4]*y + m[8]*z + m[12]) * w
	out[1] = (m[1]*x + m[5]*y + m[9]*z + m[13]) * w
	out[2] = (m[2]*x + m[6]*y + m[10]*z + m[14]) * w
	return out
}

// TransformQuat 使用四元数变换 Vec3。
func TransformQuat(out, a Vec3, q [4]float64) Vec3 {
	x := a[0]
	y := a[1]
	z := a[2]
	qx := q[0]
	qy := q[1]
	qz := q[2]
	qw := q[3]

	// Calculate quat * vec
	ix := qw*x + qy*z - qz*y
	iy := qw*y + qz*x - qx*z
	iz := qw*z + qx*y - qy*x
	iw := -qx*x - qy*y - qz*z

	// Calculate result * inverse quat
	out[0] = ix*qw + iw*-qx + iy*-qz - iz*-qy
	out[1] = iy*qw + iw*-qy + iz*-qx - ix*-qz
	out[2] = iz*qw + iw*-qz + ix*-qy - iy*-qx
	return out
}

// Angle
// 计算两个 Vec3 之间的夹角。
func Angle(a, b Vec3) float64 {
	sqrtMagA := math.Sqrt(SquaredLength(a))
	sqrtMagB := math.Sqrt(SquaredLength(b))
	magnitude := sqrtMagA * sqrtMagB
	cosine := Dot(a, b) / magnitude
	return math.Acos(math.Min(math.Max(cosine, -1), 1))
}

// RotateX 围绕 X 轴旋转 3D 向量。
// out: 接收旋转结果的向量
// a: 需要旋转的向量
// b: 旋转中心点
// rad: 旋转角度（弧度）
// 返回值: 旋转后的向量 out
func RotateX(out, a, b Vec3, rad float64) Vec3 {
	var p, r Vec3

	// 将点平移到原点
	p[0] = a[0] - b[0]
	p[1] = a[1] - b[1]
	p[2] = a[2] - b[2]

	// 执行旋转
	r[0] = p[0]
	r[1] = p[1]*math.Cos(rad) - p[2]*math.Sin(rad)
	r[2] = p[1]*math.Sin(rad) + p[2]*math.Cos(rad)

	// 平移回正确的位置
	out[0] = r[0] + b[0]
	out[1] = r[1] + b[1]
	out[2] = r[2] + b[2]

	return out
}

// RotateY 围绕 Y 轴旋转 3D 向量。
// out: 接收旋转结果的向量
// a: 需要旋转的向量
// b: 旋转中心点
// rad: 旋转角度（弧度）
// 返回值: 旋转后的向量 out
func RotateY(out, a, b Vec3, rad float64) Vec3 {
	var p, r Vec3

	// 将点平移到原点
	p[0] = a[0] - b[0]
	p[1] = a[1] - b[1]
	p[2] = a[2] - b[2]

	// 执行旋转
	r[0] = p[2]*math.Sin(rad) + p[0]*math.Cos(rad)
	r[1] = p[1]
	r[2] = p[2]*math.Cos(rad) - p[0]*math.Sin(rad)

	// 平移回正确的位置
	out[0] = r[0] + b[0]
	out[1] = r[1] + b[1]
	out[2] = r[2] + b[2]

	return out
}

// RotateZ 围绕 Z 轴旋转 3D 向量。
// out: 接收旋转结果的向量
// a: 需要旋转的向量
// b: 旋转中心点
// rad: 旋转角度（弧度）
// 返回值: 旋转后的向量 out
func RotateZ(out, a, b Vec3, rad float64) Vec3 {
	var p, r Vec3

	// 将点平移到原点
	p[0] = a[0] - b[0]
	p[1] = a[1] - b[1]
	p[2] = a[2] - b[2]

	// 执行旋转
	r[0] = p[0]*math.Cos(rad) - p[1]*math.Sin(rad)
	r[1] = p[0]*math.Sin(rad) + p[1]*math.Cos(rad)
	r[2] = p[2]

	// 平移回正确的位置
	out[0] = r[0] + b[0]
	out[1] = r[1] + b[1]
	out[2] = r[2] + b[2]

	return out
}

// Zero 将 Vec3 的分量设置为零。
func Zero(out Vec3) Vec3 {
	out[0] = 0
	out[1] = 0
	out[2] = 0
	return out
}

// ExactEquals 如果两个 Vec3 的分量完全相等，则返回 true。
func ExactEquals(a, b Vec3) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}

// Equals  如果两个 Vec3 的分量近似相等，则返回 true。
// 主要是为了防止浮点数计算错误
func Equals(a, b Vec3) bool {
	return math.Abs(a[0]-b[0]) <= cst.EPSILON &&
		math.Abs(a[1]-b[1]) <= cst.EPSILON &&
		math.Abs(a[2]-b[2]) <= cst.EPSILON
}
