/*
 * @lc app=leetcode.cn id=3418 lang=golang
 *
 * [3418] 机器人可以获得的最大金币数
 */

// @lc code=start
import (
	"math"
)

// 每次转移时对不同状态单独处理。
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	const inf = math.MinInt / 2

	dp := make([][3]int, n)
	for j := range dp {
		dp[j] = [3]int{inf, inf, inf}
	}

	dp[0][0] = coins[0][0]
	dp[0][1] = max(0, coins[0][0])
	dp[0][2] = inf

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			// 收集上方和左方的三个状态
			var from [2][3]int
			for k := 0; k < 3; k++ {
				from[0][k] = inf // 上方
				from[1][k] = inf // 左方
			}
			if i > 0 {
				from[0] = dp[j] // dp 此时还保留着上一行的状态（滚动前）
			}
			if j > 0 {
				from[1] = dp[j-1] // 当前行左侧已经更新过的状态
			}

			// 合并两个来源，得到每个次数下的最优前驱
			bestPrev := [3]int{inf, inf, inf}
			for k := 0; k < 3; k++ {
				bestPrev[k] = max(from[0][k], from[1][k])
			}

			val := coins[i][j]
			// 不感化
			if bestPrev[0] != inf {
				dp[j][0] = bestPrev[0] + val
			} else {
				dp[j][0] = inf
			}

			// 1 次感化
			dp[j][1] = inf
			if bestPrev[1] != inf {
				dp[j][1] = bestPrev[1] + val // 不感化
			}
			if val < 0 && bestPrev[0] != inf {
				dp[j][1] = max(dp[j][1], bestPrev[0]) // 感化当前负值
			}

			// 2 次感化
			dp[j][2] = inf
			if bestPrev[2] != inf {
				dp[j][2] = bestPrev[2] + val // 不感化
			}
			if val < 0 && bestPrev[1] != inf {
				dp[j][2] = max(dp[j][2], bestPrev[1]) // 感化当前负值
			}
		}
	}

	ans := inf
	for k := 0; k < 3; k++ {
		ans = max(ans, dp[n-1][k])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end