/*
 * @lc app=leetcode.cn id=1671 lang=golang
 *
 * [1671] 得到山形数组的最少删除次数
 */

// @lc code=start
func lengthOfLIS(nums []int) int {

	n := len(nums)
	// dp保存长度为i时结尾的最大值
	dp := make([]int, 0)
	for i := 0; i < n; i++ {

		// 二分查找，找到第一个大于等于当前元素的位置。
		// 如果等于len(dp)说明dp中所有长度的结尾元素都小于当前元素
		// 从而当前元素可以直接追加到尾部。
		// 否则，用当前元素替换对应位置的元素。
		pos := sort.Search(len(dp), func(r int) bool { return dp[r] >= nums[i] })
		if pos == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[pos] = nums[i]
		}
	}

	return len(dp)
}
func ReverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func minimumMountainRemovals(nums []int) int {

	n := len(nums)
	maximum := 0
	for i := 1; i < n-1; i++ {
		left := lengthOfLIS(nums[:i+1])
		if left == 1 {
			continue
		}
		temp := make([]int, max(n-i, 0))
		copy(temp, nums[i:])
		ReverseSlice(temp)
		right := lengthOfLIS(temp)
		if right == 1 {
			continue
		}
		if left+right-1 > maximum {
			maximum = left + right - 1
		}
	}
	return n - maximum
}

// @lc code=end





