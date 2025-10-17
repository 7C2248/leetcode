/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1, s2, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	dp := make([]bool, l2+1)
	dp[0] = true

	for j := 1; j <= l2; j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= l1; i++ {
		dp[0] = dp[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= l2; j++ {
			dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
				(dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[l2]
}

// @lc code=end

