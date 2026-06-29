/*
 * @lc app=leetcode.cn id=813 lang=golang
 *
 * [813] 最大平均值和的分组
 */

// @lc code=start
func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	dp := make([][]float64, n+1)
	sumN := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sumN[i] += sumN[i-1] + nums[i-1]
	}
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1.0
		}
	}
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j < k; j++ {
			for m := i - 1; m > -1; m-- {
				if dp[m][j] == -1.0 {
					continue
				}
				dp[i][j+1] = max(dp[i][j+1], dp[m][j]+float64(sumN[i]-sumN[m])/float64(i-m))
			}
		}
	}
	return dp[n][k]
}

// @lc code=end
