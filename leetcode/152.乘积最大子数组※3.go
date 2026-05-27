/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	n := len(nums)
	curMin, curMax, maximum := nums[0], nums[0], nums[0]

	for i := 1; i < n; i++ {

		// 负数元素会导致当前的最大值可能来自之前的最小值，所以需要同时存储
		curMax, curMin = max(max(curMax*nums[i], curMin*nums[i]), nums[i]), min(min(curMax*nums[i], curMin*nums[i]), nums[i])

		if curMax > maximum {
			maximum = curMax
		}
	}

	return maximum
}

// @lc code=end
