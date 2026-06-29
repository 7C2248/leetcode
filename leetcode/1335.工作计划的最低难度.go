/*
 * @lc app=leetcode.cn id=1335 lang=golang
 *
 * [1335] 工作计划的最低难度
 */

// @lc code=start
func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	pre := make([]int, n+1)
	now := make([]int, n+1)
	for i := 0; i <= n; i++ {
		pre[i] = math.MaxInt
		now[i] = math.MaxInt
	}
	pre[0] = 0
	for i := 0; i < d; i++ {
		for j := 1; j <= n; j++ {
			maximum := 0
			for k := j - 1; k > -1; k-- {
				maximum = max(maximum, jobDifficulty[k])
				if pre[k] != math.MaxInt {
					now[j] = min(now[j], pre[k]+maximum)
				}
			}
		}
		pre, now = now, pre
		for i := 0; i <= n; i++ {
			now[i] = math.MaxInt
		}
	}
	return pre[n]
}

// @lc code=end

