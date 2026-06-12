/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	if l1 < l2 {
		l1, l2 = l2, l1
		nums1, nums2 = nums2, nums1
	}
	dp := make([]int, l2+1) // n1[0:i]和n2[0:j] 为止，不交叉的线的数量

	for i := 1; i <= l1; i++ {
		prev := dp[0]
		for j := 1; j <= l2; j++ {
			temp := dp[j]
			maximum := max(dp[j-1], dp[j])
			if nums1[i-1] == nums2[j-1] {
				maximum = max(maximum, prev+1)
			}
			dp[j] = maximum
			prev = temp
		}
	}
	return dp[l2]
}

// @lc code=end