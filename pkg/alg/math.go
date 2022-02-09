package alg

import "math"

// Factorial 阶乘
func Factorial(i int) int {
	n := 1
	for j := 1; j <= i; j++ {
		n *= j
	}
	return n
}

// C 组合排序
func C(n int, k int) float64 {
	son := Factorial(n)
	mother := Factorial(k) * Factorial(n-k)
	return float64(son) / float64(mother)
}

// BezierFormula 贝塞尔曲线的基函数、公式
func BezierFormula(n int, k int, t float64) float64 {
	fk := float64(k)
	fn := float64(n)
	return C(n, k) * math.Pow(t, fk) * math.Pow(1-t, fn-fk)
}
