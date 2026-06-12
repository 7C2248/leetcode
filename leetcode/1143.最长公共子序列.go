/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	if l1 < l2 {
		text1, text2 = text2, text1
		l1, l2 = l2, l1
	}
	// 附带第0行，减少判断分支
	pre := make([]int, l2+1)
	now := make([]int, l2+1)

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				now[j] = pre[j-1] + 1
			} else {
				now[j] = max(pre[j], now[j-1])
			}
		}
		now, pre = pre, now
	}
	return pre[l2]
}

// @lc code=end