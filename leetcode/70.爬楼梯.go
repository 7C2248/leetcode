/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {

	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)

	dp[n-1], dp[n-2] = 1, 1

	for i := n - 1; i > 0; i-- {
		if i >= 2 {
			dp[i-1] += dp[i]
			dp[i-2] += dp[i]
		} else {
			dp[i-1] += dp[i]
		}
	}
	return dp[0]
}

// @lc code=end

