/*
 * @lc app=leetcode.cn id=2369 lang=golang
 *
 * [2369] 检查数组是否存在有效划分
 */

// @lc code=start
func validPartition(nums []int) bool {

	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			dp[i+1] = dp[i-1] || dp[i+1]
		}
		if i > 1 && ((nums[i] == nums[i-1] && nums[i-1] == nums[i-2]) || (nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1)) {
			dp[i+1] = dp[i-2] || dp[i+1]
		}
	}
	return dp[n]
}

// @lc code=end

