/*
 * @lc app=leetcode.cn id=2742 lang=golang
 *
 * [2742] 给墙壁刷油漆
 */

// @lc code=start
import (
	"math"
)

func paintWalls(cost []int, time []int) int {

	n := len(cost)
	dp := make([]int, n+1) // 粉刷n堵墙需要的最小cost
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := range cost {
		ToBrush := time[i] + 1
		if ToBrush >= n {
			dp[n] = min(dp[n], cost[i])
			continue
		}
		for j := n; j >= 0; j-- {
			if dp[j] != math.MaxInt {
				maxBrush := min(n, j+ToBrush)
				dp[maxBrush] = min(dp[maxBrush], dp[j]+cost[i])
			}
		}
	}
	return dp[n]
}

// @lc code=end
