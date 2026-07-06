/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	buy, sell := -prices[0], 0
	n := len(prices)
	for i := 1; i < n; i++ {
		v := prices[i]
		buy = max(buy, sell-v)
		sell = max(sell, buy+v-fee)
	}
	return sell
}

// @lc code=end

