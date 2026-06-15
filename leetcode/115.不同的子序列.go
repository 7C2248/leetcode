/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	dp := make([]int, lt+1)
	dp[0] = 1

	for i := 0; i < ls; i++ {
		c := s[i]
		for j := lt; j >= 1; j-- {
			if c == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[lt]
}

// @lc code=end

