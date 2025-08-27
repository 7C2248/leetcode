/*
 * @lc app=leetcode.cn id=918 lang=golang
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	type pair struct{ idx, val int }
	n := len(nums)
	pre, res := nums[0], nums[0]
	q := []pair{{0, pre}}
	for i := 1; i < 2*n; i++ {
		for len(q) > 0 && q[0].idx < i-n {
			q = q[1:]
		}
		pre += nums[i%n]
		res = max(res, pre-q[0].val)
		for len(q) > 0 && q[len(q)-1].val >= pre {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, pre})
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

