/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
func splitArray(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	sumN := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sumN[i] += sumN[i-1] + nums[i-1]
	}
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j < k; j++ {
			for m := i - 1; m > -1; m-- {
				if dp[m][j] == math.MaxInt {
					continue
				}
				curVal := max(sumN[i]-sumN[m], dp[m][j])
				dp[i][j+1] = min(dp[i][j+1], curVal)
			}
		}
	}
	return dp[n][k]
}

// @lc code=end

