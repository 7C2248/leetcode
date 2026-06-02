/*
 * @lc app=leetcode.cn id=3180 lang=golang
 *
 * [3180] 执行操作可获得的最大总奖励 I
 */

// @lc code=start
import (
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	maxVal := rewardValues[len(rewardValues)-1]
	dp := make([]bool, 2*maxVal+1)
	dp[0] = true

	maximum := -1
	for i := range rewardValues {
		for k := 0; k < rewardValues[i]; k++ {
			if dp[k] {
				dp[k+rewardValues[i]] = true
				if k+rewardValues[i] > maximum {
					maximum = k + rewardValues[i]
				}
			}
		}
	}
	return maximum
}

// @lc code=end

