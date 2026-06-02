/*
 * @lc app=leetcode.cn id=2787 lang=golang
 *
 * [2787] 将一个数字表示成幂的和的方案数
 */

// @lc code=start
func numberOfWays(n int, x int) int {
	const mod = 1_000_000_000 + 7
	maximum := 1
	for ; intPow(maximum, x) <= n; maximum++ {
	}
	maximum -= 1
	pre := make([]int, n+1)
	now := make([]int, n+1)

	pre[0] = 1

	for i := 1; i <= maximum; i++ {
		temp := intPow(i, x)
		copy(now, pre)
		for k, v := range pre {
			if v != 0 {
				if k+temp <= n {
					now[k+temp] += v
					now[k+temp] %= mod
				}
			}
		}
		now, pre = pre, now
	}
	return pre[n]
}

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

// @lc code=end

