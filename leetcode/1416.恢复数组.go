/*
 * @lc app=leetcode.cn id=1416 lang=golang
 *
 * [1416] 恢复数组
 */

// @lc code=start
func numberOfArrays(s string, k int) int {
	const mod = 1_000_000_007
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1

	maxLen := 0
	for temp := k; temp > 0; temp /= 10 {
		maxLen++
	}

	for i := 1; i <= n; i++ {
		num := 0
		base := 1
		for j := i - 1; j >= 0 && i-j <= maxLen; j-- {
			d := int(s[j] - '0')
			num = d*base + num
			base *= 10

			if s[j] != '0' && num <= k {
				dp[i] = (dp[i] + dp[j]) % mod
			}

			if num > k && s[j] != '0' {
				break
			}
		}
	}
	return dp[n]
}

// @lc code=end

