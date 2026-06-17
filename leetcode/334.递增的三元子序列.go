/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
func increasingTriplet(nums []int) bool {

	n := len(nums)
	m0, m1, m2 := nums[0], math.MaxInt, math.MaxInt
	for i := 1; i < n; i++ {
		if m1 < nums[i] {
			m2 = nums[i]
		}
		if m0 < nums[i] && nums[i] < m1 {
			m1 = nums[i]
		}
		if m0 > nums[i] {
			m0 = nums[i]
		}
	}
	if m2 != math.MaxInt {
		return true
	}
	return false

}

// @lc code=end

