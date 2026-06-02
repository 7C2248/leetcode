/*
 * @lc app=leetcode.cn id=2328 lang=golang
 *
 * [2328] 网格图中递增路径的数目
 */

// @lc code=start
import (
	"sort"
)

func countPaths(grid [][]int) int {

	// 由于递归路径，这道题使用dfs比排序dp更快

	const mod = 1_000_000_000 + 7
	m, n := len(grid), len(grid[0])
	type cell struct{ r, c, val int }
	cells := []*cell{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			r := &cell{r: i, c: j, val: grid[i][j]}
			cells = append(cells, r)
		}
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1
		}
	}
	sort.Slice(cells, func(i, j int) bool { return cells[i].val < cells[j].val })
	dirc := [][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{1, 0}, [2]int{-1, 0}}
	result := 0
	for _, cur := range cells {
		r, c := cur.r, cur.c
		for _, delta := range dirc {
			nr, nc := r+delta[0], c+delta[1]
			if nr > -1 && nr < m && nc > -1 && nc < n && grid[nr][nc] > grid[r][c] {
				dp[nr][nc] += dp[r][c]
				dp[nr][nc] %= mod
			}
		}
		result += dp[r][c]
		result %= mod
	}
	return result
}

// @lc code=end

