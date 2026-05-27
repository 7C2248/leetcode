/*
 * @lc app=leetcode.cn id=2684 lang=golang
 *
 * [2684] 矩阵中移动的最大次数
 */

// @lc code=start
func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = []int{grid[i][0], 0}
	}
	best := make([][]int, m)
	maximum := 0
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			var up, same, down []int = nil, nil, nil
			if j > 0 {
				up = dp[j-1]
			}
			same = dp[j]
			if j < m-1 {
				down = dp[j+1]
			}
			// 注意行列顺序
			best[j] = []int{grid[j][i], -1}
			if up != nil && up[0] < best[j][0] && up[1] != -1 && up[1]+1 > best[j][1] {
				best[j][1] = up[1] + 1
			}
			if same[0] < best[j][0] && same[1] != -1 && same[1]+1 > best[j][1] {
				best[j][1] = same[1] + 1
			}
			if down != nil && down[0] < best[j][0] && down[1] != -1 && down[1]+1 > best[j][1] {
				best[j][1] = down[1] + 1
			}

			maximum = max(maximum, best[j][1])
		}
		best, dp = dp, best
	}
	return maximum
}

// @lc code=end

