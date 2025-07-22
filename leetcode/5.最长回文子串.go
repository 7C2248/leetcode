/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {

	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	max, start := 1, 0
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			if s[i] == s[i+length-1] {
				if i+1 <= i+length-2 {
					dp[i][i+length-1] = dp[i+1][i+length-2]
				} else {
					dp[i][i+length-1] = true
				}
			} else {
				dp[i][i+length-1] = false
			}

			if dp[i][i+length-1] == true {
				if length > max {
					max = length
					start = i
				}
			}
		}
	}

	return s[start : start+max]
}

// @lc code=end

