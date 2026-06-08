/*
 * @lc app=leetcode.cn id=3592 lang=golang
 *
 * [3592] 硬币面值还原
 */

// @lc code=start
func findCoins(numWays []int) []int {
	n := len(numWays)
	//dp := make([]int, n+1)
	//dp[0] = 1
	result := []int{}
	for value, ways := range numWays {
		if ways == 1 {
			result = append(result, value+1)
			// 逆序更新，空间优化
			for i := n; i >= value+1; i-- {
				if i == value+1 {
					numWays[i-1] -= 1
				} else {
					numWays[i-1] -= numWays[i-value-2]
				}
			}
			/*
				for i := value + 1; i <= n; i++ {
					dp[i] += dp[i-value-1]
					numWays[i-1] -= dp[i-value-1]
				}
			*/
		}
	}
	for _, v := range numWays {
		if v != 0 {
			return []int{}
		}
	}
	return result
}

// @lc code=end
