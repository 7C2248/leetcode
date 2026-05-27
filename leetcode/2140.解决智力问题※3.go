/*
 * @lc app=leetcode.cn id=2140 lang=golang
 *
 * [2140] 解决智力问题
 */

// @lc code=start
/*
// 会超时
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([][2]int, n)
	dp[0][1] = questions[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		maximum := 0
		for j := 0; j < i; j++ {
			if questions[j][1]+j < i {
				if dp[j][1] > maximum {
					maximum = dp[j][1]
				}
			}
		}
		dp[i][1] = maximum + questions[i][0]

	}
	return int64(max(dp[n-1][0], dp[n-1][1]))
}*/
// 正向dp解法
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([][2]int, n)
	futureMax := make([]int, n) // futureMax[i] 表示位置 i 可用的最大值
	curMax := 0

	for i := 0; i < n; i++ {
		// curMax 代表当前位置 i 之前所有可用值的最大值
		// 之前记录的“未来事件”在 i 位置生效
		curMax = max(curMax, futureMax[i])

		if i > 0 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		}
		dp[i][1] = curMax + questions[i][0]

		// 把当前选择的影响注册到未来
		availDay := i + questions[i][1] + 1
		if availDay < n {
			futureMax[availDay] = max(futureMax[availDay], dp[i][1])
		}
	}
	return int64(max(dp[n-1][0], dp[n-1][1]))
}

/*
// 反向dp解法
func mostPoints(questions [][]int) int64 {
    n := len(questions)
    dp := make([]int64, n + 1) // 解决每道题及以后题目的最高分数
    for i := n - 1; i >= 0; i-- {
        dp[i] = max(dp[i + 1], int64(questions[i][0]) + dp[min(n, i + questions[i][1] + 1)])
    }
    return dp[0]
}
*/
// @lc code=end
