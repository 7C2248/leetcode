/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 买卖股票的最佳时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {

	pb, ps, pc := -prices[0], 0, 0
	for _, v := range prices {
		cb := max(pb, pc-v)
		cs := pb + v
		cc := max(pc, ps)
		pb, ps, pc = cb, cs, cc
	}
	return max(pc, ps)
}

// @lc code=end

