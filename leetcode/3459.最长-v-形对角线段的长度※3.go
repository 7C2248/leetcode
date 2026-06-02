/*
 * @lc app=leetcode.cn id=3459 lang=golang
 *
 * [3459] 最长 V 形对角线段的长度
 */

// @lc code=start
func lenOfVDiagonal(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	// 方向：顺时针顺序
	dirs := [4][2]int{
		{-1, -1}, // 0: 左上
		{-1, 1},  // 1: 右上
		{1, 1},   // 2: 右下
		{1, -1},  // 3: 左下
	}

	// dp[i][j][d][turn] 初始化为 -1 表示未计算
	// 关键状态设计，i,j表示位置，d表示前进方向，turn记载当前位置是否转向。
	dp := make([][][4][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][4][2]int, m)
		for j := 0; j < m; j++ {
			for d := 0; d < 4; d++ {
				dp[i][j][d][0] = -1
				dp[i][j][d][1] = -1
			}
		}
	}

	var dfs func(i, j, d, turn int) int
	dfs = func(i, j, d, turn int) int {
		if dp[i][j][d][turn] != -1 {
			return dp[i][j][d][turn]
		}

		cur := grid[i][j]
		// 只有当前格子的值在序列里合法才会被调用，所以这里只关心 next
		nextVal := -1
		if cur == 1 {
			nextVal = 2
		} else if cur == 2 {
			nextVal = 0
		} else if cur == 0 {
			nextVal = 2
		}

		maxLen := 1 // 至少包含自己

		// 1. 直走
		ni, nj := i+dirs[d][0], j+dirs[d][1]
		if ni >= 0 && ni < n && nj >= 0 && nj < m && grid[ni][nj] == nextVal {
			maxLen = max(maxLen, 1+dfs(ni, nj, d, turn))
		}

		// 2. 顺时针转弯（只能转一次）
		if turn == 0 {
			nd := (d + 1) % 4
			ni2, nj2 := i+dirs[nd][0], j+dirs[nd][1]
			if ni2 >= 0 && ni2 < n && nj2 >= 0 && nj2 < m && grid[ni2][nj2] == nextVal {
				maxLen = max(maxLen, 1+dfs(ni2, nj2, nd, 1))
			}
		}

		dp[i][j][d][turn] = maxLen
		return maxLen
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				for d := 0; d < 4; d++ {
					ans = max(ans, dfs(i, j, d, 0))
				}
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

