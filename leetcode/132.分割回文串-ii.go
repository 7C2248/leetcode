/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start
func minCut(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = 2001
	}

	for i := 0; i < n+1; i++ {
		for j := 0; i+j < n && i-j > -1; j++ {
			if s[i+j] == s[i-j] {
				dp[i+j+1] = min(dp[i+j+1], dp[i-j]+1)
			} else {
				break
			}
		}
		for j := 0; i+j+1 < n && i-j > -1; j++ {
			if s[i+j+1] == s[i-j] {
				dp[i+j+2] = min(dp[i+j+2], dp[i-j]+1)
			} else {
				break
			}
		}
	}
	return dp[n] - 1
}

// @lc code=end