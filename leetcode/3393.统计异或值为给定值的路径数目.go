/*
 * @lc app=leetcode.cn id=3393 lang=golang
 *
 * [3393] 统计异或值为给定值的路径数目
 */

// @lc code=start
func countPathsWithXorValue(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dp[i] = map[int]int{}
	}
	dp[0][grid[0][0]] = 1
	var mod int = 1e9 + 7
	for i := 1; i < n; i++ {
		for key, v := range dp[i-1] {
			dp[i][grid[0][i]^key] = v % mod
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			hmap := map[int]int{}
			if j == 0 {
				for key, v := range dp[j] {
					hmap[grid[i][j]^key] = v % mod
				}
				dp[j] = hmap
				continue
			}
			for key, v := range dp[j-1] {
				hmap[grid[i][j]^key] += v
				hmap[grid[i][j]^key] %= mod
			}
			for key, v := range dp[j] {
				hmap[grid[i][j]^key] += v
				hmap[grid[i][j]^key] %= mod
			}
			dp[j] = hmap
		}
	}
	return dp[n-1][k]
}

// @lc code=end

