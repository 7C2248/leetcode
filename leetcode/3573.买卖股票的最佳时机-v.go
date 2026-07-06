/*
 * @lc app=leetcode.cn id=3573 lang=golang
 *
 * [3573] 买卖股票的最佳时机 V
 */

// @lc code=start
func maximumProfit(prices []int, k int) int64 {

	l := make([]int, k*2)
	s := make([]int, k*2)

	for i := 0; i < k; i++ {
		l[i*2] = math.MinInt / 2
		s[i*2+1] = math.MinInt / 2
	}

	for _, v := range prices {
		// 同一天买入再卖出或卖出再买入的利润为0，不会影响到最终的结果
		for i := 2*k - 1; i > -1; i-- {
			if i == 0 {
				l[i] = max(l[i], -v)
				s[i] = max(s[i], v)
				continue
			}
			if i%2 == 1 {
				l[i] = max(l[i], l[i-1]+v)
				s[i] = max(s[i], s[i-1]-v)
			} else {
				l[i] = max(l[i], max(l[i-1], s[i-1])-v)
				s[i] = max(s[i], max(l[i-1], s[i-1])+v)
			}
		}
	}
	return int64(max(l[2*k-1], s[2*k-1]))
}

// @lc code=end

