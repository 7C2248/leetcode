func uniquePaths(grid [][]int) int {
	const mod = 1_000_000_000 + 7
	m, n := len(grid), len(grid[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var leftState, upState int = -1, -1
			if i > 0 {
				upState = grid[i-1][j]
			}
			if j > 0 {
				leftState = grid[i][j-1]
			}
			var left, up int = 0, 0
			if leftState == 0 {
				left = (dp[j-1][0] + dp[j-1][1]) % mod
			} else if leftState == 1 {
				left = dp[j-1][1]
			}
			if upState == 0 {
				up = (dp[j][0] + dp[j][1]) % mod
			} else if upState == 1 {
				up = dp[j][0]
			}
			// 注意内存复用降低GC开销
			dp[j][0] = left
			dp[j][1] = up
		}
	}
	return (dp[n-1][0] + dp[n-1][1]) % mod
}