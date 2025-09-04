/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(n int) int {

	result := 0
	for n > 0 {
		if n&1 == 1 {
			result++
		}
		n = n >> 1
	}
	return result
}

// @lc code=end

