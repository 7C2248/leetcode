/*
 * @lc app=leetcode.cn id=3603 lang=golang
 *
 * [3603] 交替方向的最小路径代价 II
 */

// @lc code=start
func minCost(m int, n int, waitCost [][]int) int64 {
	Cost := make([]int, n)
	Cost[0] = 1
	for i := 1; i < n; i++ {
		Cost[i] += Cost[i-1] + i + 1 + waitCost[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				Cost[j] += (i+1)*(j+1) + waitCost[i][j]
				continue
			}
			Cost[j] = (i+1)*(j+1) + waitCost[i][j] + min(Cost[j], Cost[j-1])
		}
	}
	result := Cost[n-1] - waitCost[m-1][n-1]
	return int64(result)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

