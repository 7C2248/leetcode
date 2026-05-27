/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}
	// 思路使用了3840的思路
	dp := make([][2]int, n)

	// 偷第一个
	dp[0][1] = nums[0]

	// 不偷第一个
	dp[0][0] = 0

	for i := 1; i < n; i++ {

		// 状态0：没偷第一个
		if i == 1 {
			dp[i][0] = nums[1]
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-2][0]+nums[i])
		}

		// 状态1：偷了第一个
		if i == n-1 {
			// 最后一个不能偷
			dp[i][1] = dp[i-1][1]
		} else if i == 1 {
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][1] = max(dp[i-1][1], dp[i-2][1]+nums[i])
		}
	}

	return max(dp[n-1][0], dp[n-1][1])
}

// @lc code=end

