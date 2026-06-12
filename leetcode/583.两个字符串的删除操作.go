/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 < l2 {
		l1, l2 = l2, l1
		word1, word2 = word2, word1
	}

	dp := make([]int, l2+1) // 只用一个一维数组

	for i := 1; i <= l1; i++ {
		prev := 0 // 对应 dp[i-1][0]，总是 0
		for j := 1; j <= l2; j++ {
			temp := dp[j] // 保存上一行的 dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev + 1
			} else {
				// dp[j-1] 已经是本行新值, temp 是上一行的 dp[j]
				if dp[j-1] > temp {
					dp[j] = dp[j-1]
				} else {
					dp[j] = temp
				}
			}
			prev = temp // 为下一次迭代准备“左上角”
		}
	}
	return l1 + l2 - 2*dp[l2]
}

// @lc code=end

