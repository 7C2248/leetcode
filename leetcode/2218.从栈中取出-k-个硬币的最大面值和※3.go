/*
 * @lc app=leetcode.cn id=2218 lang=golang
 *
 * [2218] 从栈中取出 K 个硬币的最大面值和
 */

// @lc code=start
func maxValueOfCoins(piles [][]int, k int) int {

	// 顺序遍历每个栈，记录对应栈下进行k次选择的最大值。
	dp := make([]int, k+1)
	for j := 1; j <= k; j++ {
		dp[j] = -1
	}

	for _, pile := range piles {
		// 单个栈中的前缀和等价于不同重量和价值的货物。
		pl := len(pile)
		preSum := make([]int, pl)
		for i, v := range pile {
			if i == 0 {
				preSum[0] = v
				continue
			}
			preSum[i] = v + preSum[i-1]
		}
		for i := k; i > 0; i-- {
			for j := min(pl, i) - 1; j > -1; j-- {
				if dp[i-j-1] != -1 {
					dp[i] = max(dp[i], dp[i-j-1]+preSum[j])
				}
			}
		}
	}
	return dp[k]
}

// @lc code=end