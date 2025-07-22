/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {

	n := len(nums)
	l, r := make([]int, n), make([]int, n)
	l[0] = 1
	r[n-1] = 1
	for i := 1; i < n; i++ {
		l[i] = l[i-1] * nums[i-1]
		r[n-1-i] = r[n-i] * nums[n-i]
	}
	res := make([]int, n)
	for i := 1; i < n-1; i++ {
		res[i] = l[i] * r[i]
	}
	res[0] = r[0]
	res[n-1] = l[n-1]

	return res
}

// @lc code=end

