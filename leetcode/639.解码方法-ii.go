/*
 * @lc app=leetcode.cn id=639 lang=golang
 *
 * [639] 解码方法 II
 */

// @lc code=start
func numDecodings(s string) int {
	const mod = 1_000_000_000 + 7
	n := len(s)
	if s[0] == '0' {
		return 0
	}
	dp0, dp1 := 1, 1
	if s[0] == '*' {
		dp1 = 9
	}
	for i := 1; i < n; i++ {
		cur := s[i]
		prev := s[i-1]
		curdp := 0
		if cur != '0' {
			if cur == '*' {
				curdp = dp1 * 9
			} else {
				curdp = dp1
			}
		}
		if prev == '1' {
			if cur != '*' {
				curdp += dp0
			} else {
				curdp += dp0 * 9
			}
		} else if prev == '2' {
			if cur >= '0' && cur <= '6' {
				curdp += dp0
			} else if cur == '*' {
				curdp += dp0 * 6
			}
		} else if prev == '*' {
			if cur >= '0' && cur <= '6' {
				curdp += dp0 * 2
			} else if cur >= '7' && cur <= '9' {
				curdp += dp0
			} else {
				curdp += dp0 * 15
			}
		}
		curdp %= mod
		dp0, dp1 = dp1, curdp
	}
	return dp1
}

// @lc code=end
