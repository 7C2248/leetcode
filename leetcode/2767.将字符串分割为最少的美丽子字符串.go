/*
 * @lc app=leetcode.cn id=2767 lang=golang
 *
 * [2767] 将字符串分割为最少的美丽子字符串
 */

// @lc code=start
func minimumBeautifulSubstrings(s string) int {
	n := len(s)
	var five []string
	for num := int64(1); ; num *= 5 {
		bin := strconv.FormatInt(num, 2)
		if len(bin) > n {
			break
		}
		five = append(five, bin)
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 16
	}
	for i := 1; i <= n; i++ {
		for _, w := range five {
			l := len(w)
			if i >= l && s[i-l] == '1' && s[i-l:i] == w {
				dp[i] = min(dp[i], dp[i-l]+1)
			}
		}
	}
	if dp[n] == 16 {
		return -1
	}
	return dp[n]
}

// @lc code=end
