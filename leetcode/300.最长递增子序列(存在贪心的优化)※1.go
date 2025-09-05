/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {

	n := len(nums)
	dp := make([]int, n)
	mmax := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > mmax {
					mmax = dp[i]
				}
			}
		}
	}

	return mmax + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

