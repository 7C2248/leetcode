/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	l1, l2 := len(nums1), len(nums2)
	if l1 < l2 {
		nums1, nums2 = nums2, nums1
		l1, l2 = l2, l1
	}
	// 附带第0行，减少判断分支
	pre := make([]int, l2+1)
	now := make([]int, l2+1)
	maximum := 0
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				now[j] = pre[j-1] + 1
			} else {
				now[j] = 0
			}
			maximum = max(maximum, now[j])
		}
		now, pre = pre, now
	}
	return maximum
}

// @lc code=end

