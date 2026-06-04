/*
 * @lc app=leetcode.cn id=879 lang=golang
 *
 * [879] 盈利计划
 */

// @lc code=start
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	const mod = 1_000_000_000 + 7
	// 产生利润k时使用i个员工的方案数
	// 空间优化到minProfit
	dp := make([][]int, minProfit+1)
	for i := 0; i <= minProfit; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1

	for k := range group {
		for k0 := minProfit; k0 >= 0; k0-- {
			np := min(minProfit, k0+profit[k])
			for i := n; i >= group[k]; i-- {
				if dp[k0][i-group[k]] != 0 {
					dp[np][i] += dp[k0][i-group[k]]
					dp[np][i] %= mod
				}
			}
		}
	}
	result := 0
	for _, v := range dp[minProfit] {
		result += v
		result %= mod
	}
	return result
}

// @lc code=end
