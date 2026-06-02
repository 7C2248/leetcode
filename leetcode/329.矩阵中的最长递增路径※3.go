/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	var dfs func(x, y, lastval int) int
	dfs = func(x, y, lastval int) int {
		if x >= n || x < 0 || y >= m || y < 0 || matrix[y][x] >= lastval || matrix[y][x] == -1 {
			return 0
		}
		if dp[y][x] != 0 {
			return dp[y][x]
		}
		num := matrix[y][x]
		matrix[y][x] = -1

		dp[y][x] = max(dp[y][x], dfs(x+1, y, num)+1)
		dp[y][x] = max(dp[y][x], dfs(x-1, y, num)+1)
		dp[y][x] = max(dp[y][x], dfs(x, y+1, num)+1)
		dp[y][x] = max(dp[y][x], dfs(x, y-1, num)+1)

		matrix[y][x] = num
		return dp[y][x]
	}
	maximum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				maximum = max(dfs(j, i, math.MaxInt), maximum)
			} else {
				maximum = max(dp[i][j], maximum)
			}
		}
	}
	return maximum
}

// @lc code=end

/*

// 迭代排序dp性能更优
func longestIncreasingPath(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])

	// 取网格值并排序
    type node struct {
        r, c, val int
    }
    cells := make([]node, 0, m*n)
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            cells = append(cells, node{r, c, matrix[r][c]})
        }
    }
    // 按值升序
    sort.Slice(cells, func(i, j int) bool {
        return cells[i].val < cells[j].val
    })

    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = 1  // 初始长度至少为 1
        }
    }

    dirs := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
    ans := 1
    for _, cell := range cells {
        r, c, val := cell.r, cell.c, cell.val
        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && matrix[nr][nc] < val {
                // 可以从邻居 nr,nc 走过来
                if dp[nr][nc]+1 > dp[r][c] {
                    dp[r][c] = dp[nr][nc] + 1
                }
            }
        }
        if dp[r][c] > ans {
            ans = dp[r][c]
        }
    }
    return ans
}
*/