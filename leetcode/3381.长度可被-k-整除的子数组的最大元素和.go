/*
 * @lc app=leetcode.cn id=3381 lang=golang
 *
 * [3381] 长度可被 K 整除的子数组的最大元素和
 */

// @lc code=start
func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	dp := make([]int64, n+1)
	var ans int64 = -1e18

	// 初始窗口 [0, k-1] 的和
	var curKSum int64 = 0
	for i := 0; i < k; i++ {
		curKSum += int64(nums[i])
	}
	dp[k] = curKSum
	ans = dp[k]

	// i 从 k+1 到 n，窗口每次右移
	for i := k + 1; i <= n; i++ {
		// 窗口右移：去掉 nums[i-k-1]，加入 nums[i-1]
		curKSum = curKSum - int64(nums[i-k-1]) + int64(nums[i-1])
		// 状态转移
		dp[i] = max(0, dp[i-k]) + curKSum
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

