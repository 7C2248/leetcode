/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	n := len(s)
	if s[0] == '0' {
		return 0
	}

	dp0, dp1 := 1, 1
	for i := 1; i < n; i++ {
		cur := s[i]
		prev := s[i-1]
		curdp := 0
		if cur != '0' {
			curdp += dp1
		}
		if (prev == '1') || (prev == '2' && cur >= '0' && cur < '7') {
			curdp += dp0
		}
		dp0, dp1 = dp1, curdp
	}
	return dp1
}

// @lc code=end


