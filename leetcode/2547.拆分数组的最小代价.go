/*
 * @lc app=leetcode.cn id=2547 lang=golang
 *
 * [2547] 拆分数组的最小代价
 */

// @lc code=start
func minCost(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < n; i++ {
		// 每次反向遍历时需要利用已经遍历过的数据，不要每次都重新计算
		freq := make(map[int]int, i+1)
		dupSum := 0
		for j := i; j > -1; j-- {
			x := nums[j]
			if freq[x] == 0 {
				freq[x] = 1
			} else if freq[x] == 1 {
				freq[x] = 2
				dupSum += 2
			} else {
				freq[x]++
				dupSum++
			}
			dp[i+1] = min(dp[i+1], dp[j]+dupSum+k)
		}
	}
	return dp[n]
}

// @lc code=end
