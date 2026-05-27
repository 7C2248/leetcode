/*
 * @lc app=leetcode.cn id=2466 lang=golang
 *
 * [2466] 统计构造好字符串的方案数
 */

// @lc code=start
func countGoodStrings(low int, high int, zero int, one int) int {

	dp := make([]int, high+1)
	dp[0] = 1

	l := min(zero, one)
	for i := l; i < high+1; i++ {
		temp := 0
		if i-zero >= 0 {
			temp += dp[i-zero]
		}
		if i-one >= 0 {
			temp += dp[i-one]
		}
		dp[i] += temp
		dp[i] %= int(1e9 + 7)
	}
	result := 0
	for i := low; i <= high; i++ {
		result += dp[i]
		result %= int(1e9 + 7)
	}

	return result
}

// @lc code=end

