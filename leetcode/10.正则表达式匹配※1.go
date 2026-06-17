/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)

	k := 0
	mutiluse := make([]bool, lp)
	np := []byte{}
	for i := 0; i < lp; i++ {
		if p[i] != '*' {
			np = append(np, p[i])
			k++
		} else {
			mutiluse[k-1] = true
		}
	}
	lnp := len(np)
	dp := make([]bool, lnp+1)
	dp[0] = true
	for i := 1; i <= lnp; i++ {
		if mutiluse[i-1] {
			dp[i] = dp[i-1]
		}
	}

	for i := 1; i <= ls; i++ {
		prev := dp[0]
		// 对于 i > 0 时使用0个np中的字符一定无法匹配。
		dp[0] = false
		for j := 1; j <= lnp; j++ {
			temp := dp[j]
			if mutiluse[j-1] {
				dp[j] = ((s[i-1] == np[j-1] || np[j-1] == '.') && dp[j]) || (dp[j-1])
			} else {
				dp[j] = (s[i-1] == np[j-1] || np[j-1] == '.') && prev
			}
			prev = temp
		}
	}
	return dp[lnp]
}

// @lc code=end

