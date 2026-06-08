/*
 * @lc app=leetcode.cn id=3181 lang=golang
 *
 * [3181] 执行操作可获得的最大总奖励 II
 */

// @lc code=start
import (
	"sort"
)

func getMax(v uint64) int {
	n := 0
	for v > 0 {
		n += 1
		v >>= 1
	}
	return n
}

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	maxVal := rewardValues[len(rewardValues)-1]

	maxRange := int((2*maxVal+1)/64 + 1)
	dp := make([]uint64, maxRange)
	dp[0] = 1
	maximum := -1
	for i := range rewardValues {
		val := rewardValues[i]
		effectRange := val / 64
		leftMove := val % 64

		for k := effectRange; k > -1; k-- {
			if k+effectRange >= maxRange {
				continue
			}

			var needToMove uint64
			if k == effectRange {
				needToMove = dp[k] & (1<<leftMove - 1)
			} else {
				needToMove = dp[k]
			}

			if needToMove == 0 {
				continue
			}

			dp[k+effectRange] |= needToMove << leftMove

			if leftMove != 0 {
				overflow := needToMove >> (64 - leftMove)
				if overflow != 0 {
					if k+effectRange+1 >= maxRange {
						dp = append(dp, 0)
						maxRange += 1
					}
					dp[k+effectRange+1] |= overflow
				}
			}

			/*
				// 初始实现代码，虽然能过，都能非常慢约1800ms
				// 主要性能瓶颈，由于大量调用循环判断最高位，导致开销巨大
				offset0 := getMax(needToMove)
				dp[k+effectRange] |= needToMove << leftMove
				if offset0-1+leftMove >= 64{
					if k + effectRange + 1 >= maxRange{
						dp = append(dp,0)
						maxRange += 1
					}
					nextRight:= 64 - leftMove
					dp[k+effectRange+1] |= needToMove >> nextRight

				}
			*/
		}
	}

	for i := maxRange - 1; i > -1; i-- {
		// 可以用bits标准库加快查找
		offset := getMax(dp[i])
		if offset != 0 {
			maximum = offset + 64*i - 1
			break
		}
	}
	return maximum
}

// @lc code=end

/*
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
*/