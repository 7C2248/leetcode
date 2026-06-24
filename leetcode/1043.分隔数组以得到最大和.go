/*
 * @lc app=leetcode.cn id=1043 lang=golang
 *
 * [1043] 分隔数组以得到最大和
 */

// @lc code=start
func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		maximum := 0
		limit := i
		if limit >= k {
			limit = k - 1
		}
		for j := 0; j <= limit; j++ {
			if arr[i-j] > maximum {
				maximum = arr[i-j]
			}
			dp[i+1] = max(dp[i-j]+maximum*(j+1), dp[i+1])
		}
	}
	return dp[n]
}

// @lc code=end

