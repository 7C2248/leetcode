/*
 * @lc app=leetcode.cn id=2707 lang=golang
 *
 * [2707] 字符串中的额外字符
 */

// @lc code=start
func minExtraChar(s string, dictionary []string) int {
	n := len(s)
	hmap := map[string]bool{}
	for _, word := range dictionary {
		hmap[word] = true
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 51
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if hmap[s[j:i]] {
				dp[i] = min(dp[i], dp[j])
			} else {
				dp[i] = min(dp[i], dp[j]+i-j)
			}
		}
	}
	return dp[n]
}

// @lc code=end

