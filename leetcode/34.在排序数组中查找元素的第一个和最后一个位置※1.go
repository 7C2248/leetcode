/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
		if nums[mid] == target {
			ans = mid
		}
	}
	return ans
}

func findLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		if nums[mid] == target {
			ans = mid
		}
	}
	return ans
}

// @lc code=end

