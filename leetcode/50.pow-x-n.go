/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if x == 0.0 {
		return 0.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}

	result := 1.0
	base := x
	for n > 0 {
		if n&1 == 1 {
			result *= base
		}
		base *= base
		n >>= 1
	}

	return result
}

// @lc code=end

