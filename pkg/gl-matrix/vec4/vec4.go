package vec4

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/mocheer/xena/pkg/gl-matrix/cst"
)

// https://github.com/toji/gl-matrix/blob/master/src/vec4.js
// https://github.com/flywave/go3d/blob/master/vec4/vec4.go

// Vec4 表示一个四维向量
type Vec4 [4]float64

// Create 创建一个新的零向量
func Create() Vec4 {
	return Vec4{0, 0, 0, 0}
}

// Clone 克隆一个向量
func Clone(a Vec4) Vec4 {
	return Vec4{a[0], a[1], a[2], a[3]}
}

// FromValues 根据给定的值创建一个新的向量
func FromValues(x, y, z, w float64) Vec4 {
	return Vec4{x, y, z, w}
}

// Copy 将一个向量的值复制到另一个向量
func Copy(out *Vec4, a Vec4) Vec4 {
	out[0] = a[0]
	out[1] = a[1]
	out[2] = a[2]
	out[3] = a[3]
	return *out
}

// Set 设置向量的分量
func Set(out *Vec4, x, y, z, w float64) Vec4 {
	out[0] = x
	out[1] = y
	out[2] = z
	out[3] = w
	return *out
}

// Add 两个向量相加
func Add(out *Vec4, a, b Vec4) Vec4 {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	return *out
}

// Subtract 两个向量相减
func Subtract(out *Vec4, a, b Vec4) Vec4 {
	out[0] = a[0] - b[0]
	out[1] = a[1] - b[1]
	out[2] = a[2] - b[2]
	out[3] = a[3] - b[3]
	return *out
}

// Multiply 两个向量相乘
func Multiply(out *Vec4, a, b Vec4) Vec4 {
	out[0] = a[0] * b[0]
	out[1] = a[1] * b[1]
	out[2] = a[2] * b[2]
	out[3] = a[3] * b[3]
	return *out
}

// Divide 两个向量相除
func Divide(out *Vec4, a, b Vec4) Vec4 {
	out[0] = a[0] / b[0]
	out[1] = a[1] / b[1]
	out[2] = a[2] / b[2]
	out[3] = a[3] / b[3]
	return *out
}

// Ceil 对向量的每个分量进行向上取整
func Ceil(out *Vec4, a Vec4) Vec4 {
	out[0] = math.Ceil(a[0])
	out[1] = math.Ceil(a[1])
	out[2] = math.Ceil(a[2])
	out[3] = math.Ceil(a[3])
	return *out
}

// Floor 对向量的每个分量进行向下取整
func Floor(out *Vec4, a Vec4) Vec4 {
	out[0] = math.Floor(a[0])
	out[1] = math.Floor(a[1])
	out[2] = math.Floor(a[2])
	out[3] = math.Floor(a[3])
	return *out
}

// Min 返回两个向量每个分量的最小值
func Min(out *Vec4, a, b Vec4) Vec4 {
	out[0] = math.Min(a[0], b[0])
	out[1] = math.Min(a[1], b[1])
	out[2] = math.Min(a[2], b[2])
	out[3] = math.Min(a[3], b[3])
	return *out
}

// Max 返回两个向量每个分量的最大值
func Max(out *Vec4, a, b Vec4) Vec4 {
	out[0] = math.Max(a[0], b[0])
	out[1] = math.Max(a[1], b[1])
	out[2] = math.Max(a[2], b[2])
	out[3] = math.Max(a[3], b[3])
	return *out
}

// Round 对向量的每个分量进行四舍五入
func Round(out *Vec4, a Vec4) Vec4 {
	out[0] = math.Round(a[0])
	out[1] = math.Round(a[1])
	out[2] = math.Round(a[2])
	out[3] = math.Round(a[3])
	return *out
}

// Scale 向量缩放
func Scale(out *Vec4, a Vec4, b float64) Vec4 {
	out[0] = a[0] * b
	out[1] = a[1] * b
	out[2] = a[2] * b
	out[3] = a[3] * b
	return *out
}

// ScaleAndAdd 向量缩放后相加
func ScaleAndAdd(out *Vec4, a, b Vec4, scale float64) Vec4 {
	out[0] = a[0] + b[0]*scale
	out[1] = a[1] + b[1]*scale
	out[2] = a[2] + b[2]*scale
	out[3] = a[3] + b[3]*scale
	return *out
}

// Distance 计算两个向量之间的欧几里得距离
func Distance(a, b Vec4) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	z := b[2] - a[2]
	w := b[3] - a[3]
	return math.Sqrt(x*x + y*y + z*z + w*w)
}

// SquaredDistance 计算两个向量之间的平方欧几里得距离
func SquaredDistance(a, b Vec4) float64 {
	x := b[0] - a[0]
	y := b[1] - a[1]
	z := b[2] - a[2]
	w := b[3] - a[3]
	return x*x + y*y + z*z + w*w
}

// Length 计算向量的长度
func Length(a Vec4) float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := a[3]
	return math.Sqrt(x*x + y*y + z*z + w*w)
}

// SquaredLength 计算向量的平方长度
func SquaredLength(a Vec4) float64 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := a[3]
	return x*x + y*y + z*z + w*w
}

// Negate 向量取反
func Negate(out *Vec4, a Vec4) Vec4 {
	out[0] = -a[0]
	out[1] = -a[1]
	out[2] = -a[2]
	out[3] = -a[3]
	return *out
}

// Inverse 向量取倒数
func Inverse(out *Vec4, a Vec4) Vec4 {
	out[0] = 1.0 / a[0]
	out[1] = 1.0 / a[1]
	out[2] = 1.0 / a[2]
	out[3] = 1.0 / a[3]
	return *out
}

// Normalize 向量归一化
func Normalize(out *Vec4, a Vec4) Vec4 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := a[3]
	len := x*x + y*y + z*z + w*w
	if len > 0 {
		len = 1 / math.Sqrt(len)
	}
	out[0] = x * len
	out[1] = y * len
	out[2] = z * len
	out[3] = w * len
	return *out
}

// Dot 计算两个向量的点积
func Dot(a, b Vec4) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

// Cross 计算三个向量的叉积
func Cross(out *Vec4, u, v, w Vec4) Vec4 {
	A := v[0]*w[1] - v[1]*w[0]
	B := v[0]*w[2] - v[2]*w[0]
	C := v[0]*w[3] - v[3]*w[0]
	D := v[1]*w[2] - v[2]*w[1]
	E := v[1]*w[3] - v[3]*w[1]
	F := v[2]*w[3] - v[3]*w[2]
	G := u[0]
	H := u[1]
	I := u[2]
	J := u[3]

	out[0] = H*F - I*E + J*D
	out[1] = -(G * F) + I*C - J*B
	out[2] = G*E - H*C + J*A
	out[3] = -(G * D) + H*B - I*A

	return *out
}

// Lerp 两个向量之间的线性插值
func Lerp(out *Vec4, a, b Vec4, t float64) Vec4 {
	ax := a[0]
	ay := a[1]
	az := a[2]
	aw := a[3]
	out[0] = ax + t*(b[0]-ax)
	out[1] = ay + t*(b[1]-ay)
	out[2] = az + t*(b[2]-az)
	out[3] = aw + t*(b[3]-aw)
	return *out
}

// Random 生成一个随机向量
func Random(out *Vec4, scale float64) Vec4 {
	if scale == 0 {
		scale = 1.0
	}

	rand.Seed(time.Now().UnixNano())
	v1 := rand.Float64()*2 - 1
	v2 := (4*rand.Float64() - 2) * math.Sqrt(rand.Float64()*-rand.Float64()+rand.Float64())
	s1 := v1*v1 + v2*v2

	v3 := rand.Float64()*2 - 1
	v4 := (4*rand.Float64() - 2) * math.Sqrt(rand.Float64()*-rand.Float64()+rand.Float64())
	s2 := v3*v3 + v4*v4

	d := math.Sqrt((1 - s1) / s2)
	out[0] = scale * v1
	out[1] = scale * v2
	out[2] = scale * v3 * d
	out[3] = scale * v4 * d
	return *out
}

// TransformMat4 用矩阵变换向量
func TransformMat4(out *Vec4, a Vec4, m [16]float64) Vec4 {
	x := a[0]
	y := a[1]
	z := a[2]
	w := a[3]
	out[0] = m[0]*x + m[4]*y + m[8]*z + m[12]*w
	out[1] = m[1]*x + m[5]*y + m[9]*z + m[13]*w
	out[2] = m[2]*x + m[6]*y + m[10]*z + m[14]*w
	out[3] = m[3]*x + m[7]*y + m[11]*z + m[15]*w
	return *out
}

// TransformQuat 用四元数变换向量
func TransformQuat(out *Vec4, a Vec4, q Vec4) Vec4 {
	x := a[0]
	y := a[1]
	z := a[2]
	qx := q[0]
	qy := q[1]
	qz := q[2]
	qw := q[3]

	// 计算 quat * vec
	ix := qw*x + qy*z - qz*y
	iy := qw*y + qz*x - qx*z
	iz := qw*z + qx*y - qy*x
	iw := -qx*x - qy*y - qz*z

	// 计算 result * inverse quat
	out[0] = ix*qw + iw*-qx + iy*-qz - iz*-qy
	out[1] = iy*qw + iw*-qy + iz*-qx - ix*-qz
	out[2] = iz*qw + iw*-qz + ix*-qy - iy*-qx
	out[3] = a[3]
	return *out
}

// Zero 将向量分量设置为零
func Zero(out *Vec4) Vec4 {
	out[0] = 0.0
	out[1] = 0.0
	out[2] = 0.0
	out[3] = 0.0
	return *out
}

// Str 返回向量的字符串表示
func Str(a Vec4) string {
	return fmt.Sprintf("vec4(%v, %v, %v, %v)", a[0], a[1], a[2], a[3])
}

// ExactEquals 判断两个向量是否完全相等
func ExactEquals(a, b Vec4) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2] && a[3] == b[3]
}

// Equals 判断两个向量是否近似相等
func Equals(a, b Vec4) bool {
	a0 := a[0]
	a1 := a[1]
	a2 := a[2]
	a3 := a[3]
	b0 := b[0]
	b1 := b[1]
	b2 := b[2]
	b3 := b[3]
	return math.Abs(a0-b0) <= cst.EPSILON*math.Max(1.0, math.Max(math.Abs(a0), math.Abs(b0))) &&
		math.Abs(a1-b1) <= cst.EPSILON*math.Max(1.0, math.Max(math.Abs(a1), math.Abs(b1))) &&
		math.Abs(a2-b2) <= cst.EPSILON*math.Max(1.0, math.Max(math.Abs(a2), math.Abs(b2))) &&
		math.Abs(a3-b3) <= cst.EPSILON*math.Max(1.0, math.Max(math.Abs(a3), math.Abs(b3)))
}

// ForEach 对数组中的每个向量执行操作
func ForEach(a []float64, stride, offset, count int, fn func(out, arg Vec4), arg Vec4) []float64 {
	if stride == 0 {
		stride = 4
	}

	if offset == 0 {
		offset = 0
	}

	var l int
	if count > 0 {
		l = int(math.Min(float64(count*stride+offset), float64(len(a))))
	} else {
		l = len(a)
	}

	for i := offset; i < l; i += stride {
		vec := Vec4{a[i], a[i+1], a[i+2], a[i+3]}
		fn(vec, arg)
		a[i] = vec[0]
		a[i+1] = vec[1]
		a[i+2] = vec[2]
		a[i+3] = vec[3]
	}

	return a
}
