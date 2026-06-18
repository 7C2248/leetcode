/*
 * @lc app=leetcode.cn id=2826 lang=golang
 *
 * [2826] 将三个组排序
 */

// @lc code=start
func minimumOperations(nums []int) int {
	// 分别为以1，2，3结尾的最长递增子序列长度
	dp1, dp2, dp3 := 0, 0, 0
	for _, x := range nums {
		switch x {
		case 1:
			dp1++
		case 2:
			dp2 = max(dp1, dp2) + 1
		case 3:
			dp3 = max(dp1, max(dp2, dp3)) + 1
		}
	}
	L := max(dp1, max(dp2, dp3))
	return len(nums) - L
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// 初始解法，求出最长递增子序列然后计算最短删除长度，时间为O(n^2)
func increasingTriplet(nums []int) bool {


	n := len(nums)
	m0,m1,m2 := nums[0],math.MaxInt,math.MaxInt
	for i := 1; i < n; i++ {
		if m1 < nums[i] {
			m2 = nums[i]
		}
		if m0 < nums[i] && nums[i] < m1 {
			m1 = nums[i]
		}
		if m0 > nums[i] {
			m0 = nums[i]
		}
	}

	fmt.Println(m0,m1,m2)

	if m2 != math.MaxInt {
		return true
	}
	return false

}
*/