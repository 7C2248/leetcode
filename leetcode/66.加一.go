/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {

	up := true
	i := len(digits) - 1
	for up {
		if i < 0 {
			digits = append([]int{1}, digits...)
			break
		}
		digits[i]++
		if digits[i] > 9 {
			digits[i] -= 10
			up = true
			i--
		} else {
			up = false
		}
	}
	return digits
}

// @lc code=end

