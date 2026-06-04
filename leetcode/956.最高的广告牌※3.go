/*
 * @lc app=leetcode.cn id=956 lang=golang
 *
 * [956] 最高的广告牌
 */

// @lc code=start
func tallestBillboard(rods []int) int {
	sum := 0
	for _, v := range rods {
		sum += v
	}
	// dp[diff] = 在差为 diff 时，较高支架的最大高度
	dp := make([]int, sum+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for _, h := range rods {
		// 拷贝上一轮状态，防止本轮更新干扰
		prev := make([]int, sum+1)
		copy(prev, dp)

		for diff := 0; diff <= sum; diff++ {
			if prev[diff] == -1 {
				continue
			}
			taller := prev[diff]

			// 1. 放到高侧
			if diff+h <= sum {
				dp[diff+h] = max(dp[diff+h], taller+h)
			}
			// 2. 放到低侧 (h <= diff)
			if diff >= h {
				dp[diff-h] = max(dp[diff-h], taller)
			} else {
				// h > diff, 低侧变高侧
				if h-diff <= sum {
					dp[h-diff] = max(dp[h-diff], taller-diff+h)
				}
			}
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
