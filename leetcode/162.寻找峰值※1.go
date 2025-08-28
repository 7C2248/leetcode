/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
/*
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	if n == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}
	var i, j int
	for i, j = 0, n-1; j > i; {
		mid := (i + j) >> 1
		m, r := nums[mid], nums[mid+1]

		if r > m {
			i = mid + 1
		} else {
			if mid == 0 {
				return 0
			}
			l := nums[mid-1]
			if l < m {
				return mid
			} else {
				j = mid - 1
			}
		}
	}
	return i
}*/
func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] > nums[mid+1] {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

// @lc code=end

