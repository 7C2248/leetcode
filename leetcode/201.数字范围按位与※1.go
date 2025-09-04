/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */

// @lc code=start

func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		//最右位1赋0
		n &= (n - 1)
	}
	return n
}

/*
func rangeBitwiseAnd(left int, right int) int {

	differ := right - left
	result := left & right
	i := 1
	for differ > 0 {
		result &= 0xFFFFFFFF << i
		differ = differ >> 1
		i++
	}
	return result & 0xFFFFFFFF
}
*/
// @lc code=end

