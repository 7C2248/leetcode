/*
 * @lc app=leetcode.cn id=2304 lang=golang
 *
 * [2304] 网格中的最小路径代价
 */

// @lc code=start
import (
	"math"
)

func minPathCost(grid [][]int, moveCost [][]int) int {

	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = grid[0][i]
	}
	temp := make([]int, n)
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			temp[j] = math.MaxInt
			for k := 0; k < n; k++ {
				temp[j] = min(dp[k]+moveCost[grid[(i - 1)][k]][j], temp[j])
			}
			temp[j] += grid[i][j]
		}
		// 滚动数组，空间优化
		dp, temp = temp, dp
	}
	minimum := math.MaxInt
	for i := 0; i < n; i++ {
		minimum = min(minimum, dp[i])
	}
	return minimum
}

// @lc code=end

