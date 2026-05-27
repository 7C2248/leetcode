/*
 * @lc app=leetcode.cn id=1749 lang=golang
 *
 * [1749] 任意子数组和的绝对值的最大值
 */

// @lc code=start

// 前缀和
func maxAbsoluteSum(nums []int) int {
	nums = append([]int{0}, nums...)
	max, min := nums[0], nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max - min
}

/*
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func maxAbsoluteSum(nums []int) int {

	n := len(nums)

	curMin, curMax := nums[0], nums[0]
	min, max := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if curMax+nums[i] > nums[i] {
			curMax = curMax + nums[i]
		} else {
			curMax = nums[i]
		}
		if curMin+nums[i] < nums[i] {
			curMin = curMin + nums[i]
		} else {
			curMin = nums[i]
		}
		if curMax > max {
			max = curMax
		}
		if curMin < min {
			min = curMin
		}
	}
	if abs(min) > abs(max) {
		return abs(min)
	}

	return abs(max)
}

*/

// @lc code=end

