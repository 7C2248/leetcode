/*
 * @lc app=leetcode.cn id=3144 lang=golang
 *
 * [3144] 分割字符频率相等的最少子字符串
 */

// @lc code=start
func minimumSubstringsInPartition(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1001
	}
	for i := 0; i < n; i++ {
		freq := [26]int{}
		maxFreq := 0
		unique := 0

		for j := i; j >= 0; j-- {
			idx := s[j] - 'a'
			if freq[idx] == 0 {
				unique++
			}
			freq[idx]++
			maxFreq = max(freq[idx], maxFreq)

			length := i - j + 1
			// maxFreq*unique>=length,只有当平衡时才能取等
			if maxFreq*unique == length {
				dp[i+1] = min(dp[i+1], dp[j]+1)
			}
		}
	}
	return dp[n]
}

// @lc code=end
