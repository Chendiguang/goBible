package isprime

import (
	"math"
)

// O(Sqrt(n))
func IsPrimeSimple(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 判断2之后，就可以判断从3到sqrt(n)之间的奇数了
// 无需再判断之间的偶数，时间复杂度O(sqrt(n)/2)
func IsPrimeImprove(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i++ {
		if n%2 != 0 && n%i == 0 {
			return false
		}
	}
	return true
}

// 质数分布的规律: 大于等于5的质数一定和6的倍数相邻
// 时间复杂度O(sqrt(n)/3)
func IsPrimeBest(n int) bool {
	// 处理特殊点
	if n == 2 || n == 3 {
		return true
	}

	// 不在6的倍数两边必定不是素数
	if t := n % 6; t != 1 && t != 5 {
		return false
	}

	// 在6的倍数两边也可能不是素数，根据判断：必须是6x-1 或 6x+1这种格式
	// 所以在步长为6的循坏中，不需要考虑6i, 6i+2, 6i+3, 6i+4这种四种情况
	tmp := int(math.Sqrt(float64(n)))
	for i := 5; i <= tmp; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
