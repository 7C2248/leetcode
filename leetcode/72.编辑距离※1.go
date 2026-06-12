/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 < l2 {
		l1, l2 = l2, l1
		word1, word2 = word2, word1
	}
	if l2 == 0 {
		return l1
	}

	dp := make([]int, l2+1) // word1[0:i] 和 word2[0:j]之间的编辑距离
	for i:=1;i<=l2;i++{
		dp[i]=i
	}
	for i := 1; i <= l1; i++ {
		dp[0] = i - 1
		prev := dp[0]
		for j := 1; j <= l2; j++ {
			temp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = min(min(dp[j],dp[j-1])+1, prev)
			} else {
				dp[j] = min(min(dp[j],dp[j-1]), prev) + 1
			}
			prev = temp
		}
	}
	return dp[l2]
}

// @lc code=end

  s e a
a 1 
t
e