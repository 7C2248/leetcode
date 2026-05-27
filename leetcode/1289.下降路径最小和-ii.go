/*
 * @lc app=leetcode.cn id=1289 lang=golang
 *
 * [1289] 下降路径最小和  II
 */

// @lc code=start
import (
	"math"
)

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}

	// 初始化第一行的三个值
	prev_min1, prev_min2 := math.MaxInt, math.MaxInt
	prev_idx1 := -1
	for j, val := range grid[0] {
		if val < prev_min1 {
			prev_min2 = prev_min1
			prev_min1 = val
			prev_idx1 = j
		} else if val < prev_min2 {
			prev_min2 = val
		}
	}

	var cur_min1, cur_min2, cur_idx1 int
	// 逐行转移，只维护三个变量
	for i := 1; i < n; i++ {
		cur_min1, cur_min2 = math.MaxInt, math.MaxInt
		cur_idx1 = -1

		for j := 0; j < n; j++ {
			// 计算当前列候选值
			var candidate int
			if j == prev_idx1 {
				candidate = grid[i][j] + prev_min2
			} else {
				candidate = grid[i][j] + prev_min1
			}
			// 更新当前行的最小值、次小值及索引
			if candidate < cur_min1 {
				cur_min2 = cur_min1
				cur_min1 = candidate
				cur_idx1 = j
			} else if candidate < cur_min2 {
				cur_min2 = candidate
			}
		}
		prev_min1, prev_min2, prev_idx1 = cur_min1, cur_min2, cur_idx1
	}

	return prev_min1
}

// @lc code=end

