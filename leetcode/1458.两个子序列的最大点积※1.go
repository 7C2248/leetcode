/*
 * @lc app=leetcode.cn id=1458 lang=golang
 *
 * [1458] 两个子序列的最大点积
 */

// @lc code=start
func maxDotProduct(nums1 []int, nums2 []int) int {
	const MIN = math.MinInt / 2
	l1, l2 := len(nums1), len(nums2)
	if l1 < l2 {
		l1, l2 = l2, l1
		nums1, nums2 = nums2, nums1
	}
	dp := make([]int, l2)
	for i := 0; i < l1; i++ {
		// 必须标记不可达，否则会污染迭代值
		prev := MIN
		for j := 0; j < l2; j++ {
			var maximum int = MIN
			if prev != MIN {
				maximum = max(maximum, prev+nums1[i]*nums2[j])
			}
			if i > 0 {
				maximum = max(maximum, dp[j])
			}
			if j > 0 {
				maximum = max(maximum, dp[j-1])
			}
			temp := dp[j]
			dp[j] = max(maximum, nums1[i]*nums2[j])
			prev = temp
		}
	}
	return dp[l2-1]
}

// @lc code=end

