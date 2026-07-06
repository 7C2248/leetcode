/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {

	dp := make([]int, k*2)
	for i := 0; i < k; i++ {
		dp[i*2] = -prices[0]
	}
	for _, v := range prices {
		for i := range dp {
			if i == 0 {
				dp[i] = max(dp[i], -v)
				continue
			}
			if i%2 == 0 {
				dp[i] = max(dp[i], dp[i-1]-v)
			} else {
				dp[i] = max(dp[i], dp[i-1]+v)
			}
		}
	}
	return dp[k*2-1]
}

// @lc code=end

