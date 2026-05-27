/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost) + 1
	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n-1; i++ {
		dp[i] = min(dp[i-2], dp[i-1]) + cost[i]
	}
	dp[n-1] = min(dp[n-3], dp[n-2])
	return dp[n-1]
}

// @lc code=end

