func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(nums []int, colors []int) int64 {

	n := len(nums)
	// 维护两个动态数组，分别存储位置x时，是否选择元素x对应的最大和
	dp := make([][2]int, n)
	dp[0] = [2]int{0, nums[0]}
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		if colors[i] == colors[i-1] {
			dp[i][1] = dp[i-1][0] + nums[i]
		} else {
			dp[i][1] = dp[i][0] + nums[i]
		}
	}
	return int64(max(dp[n-1][0], dp[n-1][1]))
}
