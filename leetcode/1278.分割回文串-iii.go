/*
 * @lc app=leetcode.cn id=1278 lang=golang
 *
 * [1278] 分割回文串 III
 */

// @lc code=start
func palindromePartition(s string, k int) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			countO := 0
			// 奇数串
			for m := 0; i+m < n && i-m > -1; m++ {
				if s[i+m] != s[i-m] {
					countO += 1
				}
				if dp[i-m][j] == math.MaxInt {
					continue
				}
				dp[i+m+1][j+1] = min(dp[i+m+1][j+1], dp[i-m][j]+countO)
			}
			countM := 0
			// 偶数串
			for m := 0; i+m+1 < n && i-m > -1; m++ {
				if s[i+m+1] != s[i-m] {
					countM += 1
				}
				if dp[i-m][j] == math.MaxInt {
					continue
				}
				dp[i+m+2][j+1] = min(dp[i+m+2][j+1], dp[i-m][j]+countM)
			}
		}
	}
	return dp[n][k]
}

// @lc code=end
