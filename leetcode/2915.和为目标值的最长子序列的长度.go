/*
 * @lc app=leetcode.cn id=2915 lang=golang
 *
 * [2915] 和为目标值的最长子序列的长度
 */

// @lc code=start
func lengthOfLongestSubsequence(nums []int, target int) int {
	n := len(nums)
	pre := make([]int, target+1)
	now := make([]int, target+1)
	for i := 0; i <= target; i++ {
		pre[i] = -1
		now[i] = -1
	}
	pre[0] = 0

	for i := 0; i < n; i++ {
		for index, v := range pre {
			if v != -1 {
				now[index] = max(now[index], v)
				temp := index + nums[i]
				if temp <= target {
					now[temp] = max(now[temp], max(pre[temp], v+1))
				}
			}
		}
		now, pre = pre, now
		for index := range now {
			now[index] = -1
		}
	}
	return pre[target]
}

// @lc code=end

