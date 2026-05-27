/*
 * @lc app=leetcode.cn id=2320 lang=golang
 *
 * [2320] 统计放置房子的方式数
 */

// @lc code=start
func countHousePlacements(n int) int {

	if n == 1 {
		return 4
	}
	if n == 2 {
		return 9
	}

	mod := int(1e9 + 7)

	dp := make([]int, n)

	dp[0] = 1
	dp[1] = 1
	result := 2 //sum to dp[i-1]
	for i := 2; i < n; i++ {
		dp[i] = (dp[i] + result - dp[i-1]) % mod
		dp[i] += 1
		result = (result + dp[i]) % mod
	}
	result += 1
	return result * result % mod
}

// @lc code=end

