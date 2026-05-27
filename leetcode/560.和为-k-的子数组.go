/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	n := len(nums)
	pre := 0
	result := 0
	for i := 0; i < n; i++ {
		pre += nums[i]
		if count, ok := m[pre-k]; ok {
			result += count
		}
		m[pre] += 1
	}
	return result
}

// @lc code=end

