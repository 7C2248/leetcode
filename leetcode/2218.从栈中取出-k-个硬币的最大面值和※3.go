/*
 * @lc app=leetcode.cn id=2218 lang=golang
 *
 * [2218] 从栈中取出 K 个硬币的最大面值和
 */

// @lc code=start
func maxValueOfCoins(piles [][]int, k int) int {
	// dp[j] = 取 j 枚硬币的最大价值，-1 表示不可达
	dp := make([]int, k+1)
	for j := 1; j <= k; j++ {
		dp[j] = -1
	}

	for _, pile := range piles {
		// 最多能从当前栈取多少枚
		m := len(pile)
		if m > k {
			m = k
		}
		// 计算前缀和（1-based，prefix[t] 表示取 t 枚的价值）
		prefix := make([]int, m+1)
		for t := 1; t <= m; t++ {
			prefix[t] = prefix[t-1] + pile[t-1]
		}

		// 分组背包：倒序更新 dp
		for j := k; j >= 0; j-- {
			// 遍历当前栈可取的数量 t
			for t := 1; t <= m; t++ {
				if j >= t && dp[j-t] != -1 {
					if dp[j] < dp[j-t]+prefix[t] {
						dp[j] = dp[j-t] + prefix[t]
					}
				}
			}
		}
	}

	// 题目保证 k <= 总硬币数，所以 dp[k] 一定可达
	return dp[k]
}

// @lc code=end
