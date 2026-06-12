/*
 * @lc app=leetcode.cn id=712 lang=golang
 *
 * [712] 两个字符串的最小ASCII删除和
 */

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
	l1, l2 := len(s1), len(s2)
	if l1 < l2 {
		l1, l2 = l2, l1
		s1, s2 = s2, s1
	}
	sum := 0
	for i := 0; i < l1; i++ {
		sum += int(s1[i])
	}
	for i := 0; i < l2; i++ {
		sum += int(s2[i])
	}
	dp := make([]int, l2+1)
	for i := 1; i <= l1; i++ {
		// 左上角元素
		prev := 0
		for j := 1; j <= l2; j++ {
			temp := dp[j]
			if s1[i-1] == s2[j-1] {
				// 假设val = s1[i-1] = s2[i-1]
				// 对于上一行，dp[j]有两种可能，要么使用了s2[i-1]，要么没用
				// 没用的情况下，这个值等于上一行的dp[j]
				// 如果使用了s2[i-1]，这个值为s1的前i-1个字符中最长公共子序列的
				// ASCII和加上val，由于dp是非减的，所以这个和一定小于等于dp[j-1]+val
				// 从而prev 一定大于等于左侧或上侧的元素。
				dp[j] = prev + int(s2[j-1])
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = temp
		}
	}
	return sum - 2*dp[l2]
}

// @lc code=end