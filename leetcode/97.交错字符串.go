/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	ls1, ls2, ls3 := len(s1), len(s2), len(s3)
	if ls1+ls2 != ls3 {
		return false
	}
	dp := make([]bool, ls2+1)

	dp[0] = true
	for i := 0; i <= ls1; i++ {
		for j := 0; j <= ls2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			dp[j] = (i > 0 && dp[j] && s1[i-1] == s3[i+j-1]) ||
				(j > 0 && dp[j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[ls2]
}

// @lc code=end