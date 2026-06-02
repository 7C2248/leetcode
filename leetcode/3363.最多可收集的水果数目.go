/*
 * @lc app=leetcode.cn id=3363 lang=golang
 *
 * [3363] 最多可收集的水果数目
 */

// @lc code=start
func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)
	p0, p1, p2 := 0, 0, 0
	for i := 0; i < n; i++ {
		p0 += fruits[i][i]
		fruits[i][i] = 0
	}

	dp := make([][][2]int, n)
	dp[0] = [][2]int{[2]int{fruits[0][n-1], -1}}
	for i := 1; i < n; i++ {
		curLen := min(i+1, n-i)
		dp[i] = make([][2]int, curLen)
		for k, v := range dp[i-1] {
			if k < curLen {
				if v[0] > dp[i][k][0] {
					dp[i][k][0] = v[0]
					dp[i][k][1] = k
				}
			}
			if k > 0 {
				if v[0] > dp[i][k-1][0] {
					dp[i][k-1][0] = v[0]
					dp[i][k-1][1] = k
				}
			}
			if k < curLen-1 {
				if v[0] > dp[i][k+1][0] {
					dp[i][k+1][0] = v[0]
					dp[i][k+1][1] = k
				}
			}
		}
		for k := range dp[i] {
			dp[i][k][0] += fruits[i][n-1-k]
		}
	}
	p1 = dp[n-1][0][0]
	cur := dp[n-1][0][1]
	for i := n - 2; i > -1; i-- {
		fruits[i][n-1-cur] = 0
		cur = dp[i][cur][1]
	}

	dp = make([][][2]int, n)
	dp[0] = [][2]int{[2]int{fruits[n-1][0], -1}}
	for i := 1; i < n; i++ {
		curLen := min(i+1, n-i)
		dp[i] = make([][2]int, curLen)
		for k, v := range dp[i-1] {
			if k < curLen {
				if v[0] > dp[i][k][0] {
					dp[i][k][0] = v[0]
					dp[i][k][1] = k
				}
			}
			if k > 0 {
				if v[0] > dp[i][k-1][0] {
					dp[i][k-1][0] = v[0]
					dp[i][k-1][1] = k
				}
			}
			if k < curLen-1 {
				if v[0] > dp[i][k+1][0] {
					dp[i][k+1][0] = v[0]
					dp[i][k+1][1] = k
				}
			}
		}
		for k := range dp[i] {
			dp[i][k][0] += fruits[n-1-k][i]
		}
	}
	p2 = dp[n-1][0][0]
	return p0 + p1 + p2
}

// @lc code=end