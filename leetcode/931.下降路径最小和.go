/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
import (
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = matrix[0][i]
	}
	for i := 1; i < n; i++ {
		temp := make([]int, n)
		for j := 0; j < n; j++ {
			// 比较优化
			tminimum := dp[j]
			if j > 0 {
				tminimum = min(tminimum, dp[j-1])
			}
			if j < n-1 {
				tminimum = min(tminimum, dp[j+1])
			}
			temp[j] = tminimum + matrix[i][j]
		}
		dp = temp
	}
	minimum := math.MaxInt
	for i := 0; i < n; i++ {
		minimum = min(dp[i], minimum)
	}
	return minimum
}

// @lc code=end

