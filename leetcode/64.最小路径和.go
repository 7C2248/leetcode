/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if j == 0 && i == 0 {
				continue
			}
			if j == 0 {
				grid[j][i] += grid[j][i-1]
				continue
			}
			if i == 0 {
				grid[j][i] += grid[j-1][i]
				continue
			}
			grid[j][i] += min(grid[j-1][i], grid[j][i-1])
		}
	}

	return grid[m-1][n-1]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

