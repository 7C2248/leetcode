/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {

	result := 0
	x0, y0 := len(grid[0]), len(grid)

	var dfs func(grid [][]byte, y int, x int)
	dfs = func(grid [][]byte, y int, x int) {
		if y < 0 || x < 0 || y >= y0 || x >= x0 || grid[y][x] == '0' {
			return
		}

		grid[y][x] = '0'
		dfs(grid, y+1, x)
		dfs(grid, y, x+1)
		dfs(grid, y, x-1)
		dfs(grid, y-1, x)
	}

	for i := 0; i < y0; i++ {
		for j := 0; j < x0; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				result++
			}
		}
	}
	return result
}

// @lc code=end

