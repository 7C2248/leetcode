/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	var p int = (sum + target) / 2
	if p < 0 {
		return 0
	}
	pre := make([]int, p+1)
	now := make([]int, p+1)
	pre[0] = 1

	for i := 0; i < n; i++ {
		copy(now, pre)
		for index, v := range pre {
			if v != 0 {
				if index+nums[i] <= p {
					now[index+nums[i]] += v
				}
			}
		}
		pre, now = now, pre
		for i := 0; i <= p; i++ {
			now[i] = 0
		}
	}

	return max(0, pre[p])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

