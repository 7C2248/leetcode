/*
 * @lc app=leetcode.cn id=2472 lang=golang
 *
 * [2472] 不重叠回文子字符串的最大数目
 */

// @lc code=start
func maxPalindromes(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		limit := min(n-i, i+1)
		dp[i+1] = max(dp[i], dp[i+1])
		// 奇数
		for j := 0; j < limit && s[i-j] == s[i+j]; j++ {
			if j*2+1 >= k {
				dp[i+j+1] = max(dp[i+j+1], dp[i-j]+1)
			}
		}
		//偶数
		limit = min(n-i-1, i+1)
		for j := 0; j < limit && s[i+j+1] == s[i-j]; j++ {
			if (j+1)*2 >= k {
				dp[i+j+2] = max(dp[i+j+2], dp[i-j]+1)
			}
		}
	}

	return dp[n]
}

// @lc code=end
