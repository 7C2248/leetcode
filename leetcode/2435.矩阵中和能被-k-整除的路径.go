/*
 * @lc app=leetcode.cn id=2435 lang=golang
 *
 * [2435] 矩阵中和能被 K 整除的路径
 */

// @lc code=start
func numberOfPaths(grid [][]int, k int) int {
	const mod = 1_000_000_000 + 7
	m, n := len(grid), len(grid[0])
	dp := make([][50]int, n)
	dp[0][grid[0][0]%k] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var temp [50]int
			if i > 0 {
				for index, value := range dp[j] {
					temp[(index+grid[i][j])%k] += value
				}
			}
			if j > 0 {
				for index, value := range dp[j-1] {
					temp[(index+grid[i][j])%k] += value
				}
			}
			for idx := range temp {
				if temp[idx] >= mod {
					temp[idx] -= mod
				}
			}
			dp[j] = temp
		}
	}
	return dp[n-1][0]
}

// @lc code=end
