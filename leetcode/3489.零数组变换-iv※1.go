/*
 * @lc app=leetcode.cn id=3489 lang=golang
 *
 * [3489] 零数组变换 IV
 */

// @lc code=start
func minZeroArray(nums []int, queries [][]int) int {

	result := 0
	for index, num := range nums {
		pre := make([]bool, num+1)
		now := make([]bool, num+1)
		pre[0] = true
		tempR := 0
		for k, query := range queries {
			l, r, val := query[0], query[1], query[2]
			if l > index || r < index || val > num {
				continue
			}
			copy(now, pre)
			for j := 0; j+val <= num; j++ {
				if pre[j] {
					now[j+val] = true
				}
			}
			now, pre = pre, now
			if pre[num] {
				tempR = k + 1
				break
			}
		}
		if pre[num] {
			result = max(result, tempR)
		} else {
			return -1
		}
	}
	return result
}

// @lc code=end

