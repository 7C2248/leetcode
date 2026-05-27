func maxPathScore(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	dp := make([][]int, n)
	for j := 0; j < n; j++ {
		dp[j] = make([]int, k+1)
		for c := 0; c <= k; c++ {
			dp[j][c] = -1
		}
	}

	// 初始化起点
	if grid[0][0] == 0 {
		dp[0][0] = 0
	} else if k >= 1 {
		dp[0][1] = grid[0][0]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := grid[i][j]
			// 获取上方和左方的状态引用
			var up, left []int
			if i > 0 {
				up = dp[j] // 上一行同一列
			}
			if j > 0 {
				left = dp[j-1] // 当前行左侧列
			}

			if val == 0 {
				for c := 0; c <= k; c++ {
					best := -1
					if up != nil && up[c] > best {
						best = up[c]
					}
					if left != nil && left[c] > best {
						best = left[c]
					}
					dp[j][c] = best
				}
			} else {
				// 非零值：必须收集，c从1..k倒序更新
				for c := k; c >= 1; c-- {
					best := -1
					if up != nil && up[c-1] > best {
						best = up[c-1]
					}
					if left != nil && left[c-1] > best {
						best = left[c-1]
					}
					if best != -1 {
						dp[j][c] = best + val
					} else {
						dp[j][c] = -1
					}
				}
				dp[j][0] = -1 // 当前元素非0，必须收集
			}
		}
	}

	ans := -1
	for c := 0; c <= k; c++ {
		if dp[n-1][c] > ans {
			ans = dp[n-1][c]
		}
	}
	return ans
}