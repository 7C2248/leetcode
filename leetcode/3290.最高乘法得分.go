/*
 * @lc app=leetcode.cn id=3290 lang=golang
 *
 * [3290] 最高乘法得分
 */

// @lc code=start
func maxScore(a []int, b []int) int64 {

	const INF = math.MaxInt / 2

	dp := make([]int, 5)

	for i := 1; i <= 4; i++ {
		dp[i] = INF
	}

	for _, v := range b {
		prev := 0
		for i := 1; i <= 4; i++ {
			best := math.MinInt
			temp := dp[i]
			if dp[i] != INF {
				best = max(best, dp[i])
			}
			if prev != INF {
				best = max(best, prev+v*a[i-1])
			}
			if best != math.MinInt {
				dp[i] = best
			}
			prev = temp
		}
	}

	return int64(dp[4])
}

// @lc code=end

