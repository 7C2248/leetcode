/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {

	minimum := math.MaxInt
	maximum := 0
	result := 0
	for _, v := range prices {

		if v > minimum && v > maximum {
			maximum = v
		}
		if v < minimum {
			minimum = v
			maximum = 0
			continue
		}
		result = max(result, maximum-minimum)
	}

	if result < 0 {
		return 0
	}
	return result
}

// @lc code=end

