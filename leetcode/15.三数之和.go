/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {

	result := make([][]int, 0)
	n := len(nums)
	sort(nums, 0, n)
	for i := 0; i <= n-3; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		for j, k := i+1, n-1; j < k; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for nums[i]+nums[j]+nums[k] > 0 && k > j+1 {
				k--
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return result
}

func sort(aa []int, left int, right int) {

	if right-left <= 1 {
		return
	}
	i, j := left, right-1
	for i < j {
		for aa[j] > aa[left] && i < j {
			j--
		}
		for aa[i] <= aa[left] && i < j {
			i++
		}
		aa[i], aa[j] = aa[j], aa[i]
	}
	aa[left], aa[i] = aa[i], aa[left]
	sort(aa, left, i)
	sort(aa, i+1, right)
	return
}

// @lc code=end

