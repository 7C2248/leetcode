/*
 * @lc app=leetcode.cn id=3196 lang=golang
 *
 * [3196] 最大化子数组的总成本
 */

// @lc code=start
func maximumTotalCost(nums []int) int64 {
	n := len(nums)
	dp0, dp1 := 0, nums[0]
	for i := 1; i < n; i++ {
		cur := max(nums[i]+dp1, nums[i-1]-nums[i]+dp0)
		dp0, dp1 = dp1, cur
	}
	return int64(dp1)
}

// @lc code=end