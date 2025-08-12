/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {

	//vioSearch(board)
	bounderSearch(board)
}

func vioSearch(board [][]byte) {
	x0, y0 := len(board[0]), len(board)
	sign := make([][]bool, y0)
	for i := 0; i < y0; i++ {
		sign[i] = make([]bool, x0)
	}

	var region [][]int
	var bounderSign bool

	var dfsSearch func(board [][]byte, y int, x int)
	dfsSearch = func(board [][]byte, y int, x int) {

		if y < 0 || x < 0 || y >= y0 || x >= x0 {
			bounderSign = true
			return
		}

		if sign[y][x] == true || board[y][x] == 'X' {
			return
		}
		sign[y][x] = true
		region = append(region, []int{x, y})
		dfsSearch(board, y+1, x)
		dfsSearch(board, y, x+1)
		dfsSearch(board, y, x-1)
		dfsSearch(board, y-1, x)
	}

	for i := 0; i < y0; i++ {
		for j := 0; j < x0; j++ {
			if board[i][j] == 'O' && !sign[i][j] {
				region = make([][]int, 0)
				bounderSign = false
				dfsSearch(board, i, j)
				if !bounderSign {
					for _, p := range region {
						board[p[1]][p[0]] = 'X'
					}
				}

			}
		}
	}
}

func bounderSearch(board [][]byte) {

	var dfs func(board [][]byte, x, y int)
	var m, n int

	dfs = func(board [][]byte, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(board, x+1, y)
		dfs(board, x-1, y)
		dfs(board, x, y+1)
		dfs(board, x, y-1)
	}

	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n = len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	for i := 1; i < n-1; i++ {
		dfs(board, 0, i)
		dfs(board, m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// @lc code=end

