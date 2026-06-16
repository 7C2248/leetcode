/*
 * @lc app=leetcode.cn id=1639 lang=golang
 *
 * [1639] 通过给定词典构造目标字符串的方案数
 */

// @lc code=start
func numWays(words []string, target string) int {
	const mod = 1_000_000_007
	lw, lt := len(words[0]), len(target)
	if lw < lt {
		return 0
	}

	dp := make([]int, lt+1)
	dp[0] = 1

	for i := 0; i < lw; i++ {
		// 统计第 i 列各字母的出现次数
		// 内层循环优化
		cnt := [26]int{}
		for _, w := range words {
			cnt[w[i]-'a']++
		}

		for k := min(i, lt-1); k >= 0; k-- {
			c := target[k] - 'a'
			if cnt[c] > 0 {
				dp[k+1] = (dp[k+1] + dp[k]*cnt[c]) % mod
			}
		}
	}
	return dp[lt]
}

// @lc code=end

/*

func numWays(words []string, target string) int {
	const mod = 1_000_000_000 + 7

    lws,lw,lt := len(words),len(words[0]),len(target)
	if lw < lt{
		return 0
	}
	dp := make([]int,lt+1)
	dp[0]=1
	// 三层循环超时
	for i:=0;i<lw;i++{
		for k:=min(i,lt-1);k>-1;k--{
	// 对于同一个单词高度，每次查询target时是相同的，不需要每次重复计算
			for j:=0;j<lws;j++{
				if words[j][i] == target[k] {
					dp[k+1] += dp[k]
					dp[k+1] %= mod
				}
			}
		}
	}

	return dp[lt]
}
*/