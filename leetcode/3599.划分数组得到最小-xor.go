/*
 * @lc app=leetcode.cn id=3599 lang=golang
 *
 * [3599] 划分数组得到最小 XOR
 */

// @lc code=start
func minXor(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	xorN := make([]int, n+1)
	for i := 1; i <= n; i++ {
		xorN[i] = xorN[i-1] ^ nums[i-1]
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
				curval := max(dp[m][j], xorN[i]^xorN[m])
				dp[i][j+1] = min(dp[i][j+1], curval)
			}
		}
	}
	return dp[n][k]
}

// @lc code=end

