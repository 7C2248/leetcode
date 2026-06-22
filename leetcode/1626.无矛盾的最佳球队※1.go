/*
 * @lc app=leetcode.cn id=1626 lang=golang
 *
 * [1626] 无矛盾的最佳球队
 */

// @lc code=start
func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	type player struct{ score, age int }
	ps := make([]player, n)
	for i := range scores {
		ps[i] = player{scores[i], ages[i]}
	}
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].age != ps[j].age {
			return ps[i].age < ps[j].age
		}
		return ps[i].score < ps[j].score
	})

	// 贪心+二分查找无法保存分数信息。加上数据总量不大，应该使用传统dp

	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = ps[i].score
		for j := 0; j < i; j++ {
			if ps[j].score <= ps[i].score {
				dp[i] = max(dp[i], dp[j]+ps[i].score)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// @lc code=end
