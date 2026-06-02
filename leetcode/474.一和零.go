/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	result := 0
	for _, v := range strs {
		cm, cn := 0, 0
		for ci := range v {
			if v[ci] == '0' {
				cm += 1
			} else {
				cn += 1
			}
		}
		if cm > m || cn > n {
			continue
		}
		for i := m; i >= cm; i-- {
			for j := n; j >= cn; j-- {
				if dp[i-cm][j-cn] != -1 {
					dp[i][j] = max(dp[i-cm][j-cn]+1, dp[i][j])
					result = max(dp[i][j], result)
				}
			}
		}
	}
	return result
}

// @lc code=end

