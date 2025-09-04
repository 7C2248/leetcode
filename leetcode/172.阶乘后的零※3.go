/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) int {

	result := 0
	for n > 0 {
		n /= 5
		result += n
	}
	return result
}

// @lc code=end

