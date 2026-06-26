/*
 * @lc app=leetcode.cn id=3579 lang=golang
 *
 * [3579] 字符串转换需要的最小操作数
 */

// @lc code=start
func minOperations(word1 string, word2 string) int {
	// 预处理操作数可以将检索降低到O(1)，从而整体复杂度为O(n^2)
	n := len(word2)

	// 前i个字符相等的最小操作数
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 101
	}

	for i := 1; i <= n; i++ {
		cur1 := word1[:i]
		cur2 := word2[:i]
		// 整体划分
		dp[i] = min(dp[i], min(trans(cur1, cur2), trans(reverse(cur1), cur2)+1))
		for j := i + 1; j <= n; j++ {
			cur1 = word1[i:j]
			cur2 = word2[i:j]
			// 后继划分
			dp[j] = min(dp[j], dp[i]+min(trans(cur1, cur2), trans(reverse(cur1), cur2)+1))
		}
	}
	return dp[n]
}
func trans(s1, s2 string) int {
	var cnt [26][26]int
	m := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			cnt[s1[i]-'a'][s2[i]-'a']++
			m++
		}
	}
	pairs := 0
	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			pairs += min(cnt[i][j], cnt[j][i])
		}
	}
	return m - pairs
}
func reverse(s string) string {
	res := []byte(s)
	for left, right := 0, len(res)-1; left < right; left, right = left+1, right-1 {
		res[left], res[right] = res[right], res[left]
	}
	return string(res)
}

// @lc code=end




