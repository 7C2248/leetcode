/*
 * @lc app=leetcode.cn id=1824 lang=golang
 *
 * [1824] 最少侧跳次数
 */

// @lc code=start
import (
	"math"
)

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	now := []int{1, 0, 1}
	next := make([]int, 3)
	for i := 1; i < n; i++ {
		copy(next, now)
		if obstacles[i] > 0 {
			next[obstacles[i]-1] = math.MaxInt - 2
		}
		for j := 0; j < 3; j++ {
			if obstacles[i]-1 == j {
				continue
			}
			if now[j] == math.MaxInt-2 {
				temp := math.MaxInt
				for k := 0; k < 3; k++ {
					if k == j {
						continue
					}
					temp = min(temp, next[k]+1)
				}
				next[j] = temp
			}
		}
		now, next = next, now

	}
	minimum := math.MaxInt
	for _, v := range now {
		minimum = min(v, minimum)
	}
	return minimum
}

// @lc code=end

