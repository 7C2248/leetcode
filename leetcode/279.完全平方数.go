/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func pow(x, n int) int {
	base := x
	result := 1
	for n > 0 {
		if n%2 == 1 {
			result *= base
		}
		base *= base
		n >>= 1
	}
	return result
}
func numSquares(n int) int {
	const INF = 1<<31 - 1
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = INF
	}
	for i := 1; pow(i, 2) <= n; i++ {
		val := pow(i, 2)
		for j := val; j <= n; j++ {
			if dp[j-val] != INF {
				dp[j] = min(dp[j], dp[j-val]+1)
			}
		}
	}
	return dp[n]
}

// @lc code=end

