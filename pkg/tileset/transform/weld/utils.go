package weld

import (
	"math"
)

// 找到大于或等于给定值 x 的最小的 2 的幂
func ceilPowerOfTwo(x int) int {
	// math.Pow(2, math.Ceil(math.Log(value)/math.Log(2)))
	return 1 << uint(math.Ceil(math.Log2(float64(x))))
}
