/*
 * @lc app=leetcode.cn id=1155 lang=golang
 *
 * [1155] 掷骰子等于目标和的方法数
 */

// @lc code=start
func numRollsToTarget(n int, k int, target int) int {
	const mod = 1_000_000_000 + 7

	pre := make([]int, target+1)
	now := make([]int, target+1)
	pre[n], now[n] = 1, 1

	for i := 0; i < n; i++ {
		copy(now, pre)
		for j := 1; j < k; j++ {
			for index := j; index <= target; index++ {
				now[index] += pre[index-j]
				now[index] %= mod
			}
		}
		pre, now = now, pre
	}
	return pre[target]
}

// @lc code=end

