/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	//先滤除无法找到的情况，性能的最大提升点 200ms->25ms
	boardCnt := map[byte]int{}
	for _, r := range board {
		for _, c := range r {
			boardCnt[c]++
		}
	}
	wordBytes := []byte(word)
	wordCnt := map[byte]int{}
	for _, c := range wordBytes {
		wordCnt[c]++
		if wordCnt[c] > boardCnt[c] {
			return false
		}
	}

	//先从少的开始搜，些微性能提升 25ms->0ms
	if boardCnt[wordBytes[len(wordBytes)-1]] < boardCnt[wordBytes[0]] {
		slices.Reverse(wordBytes)
		word = string(wordBytes)
	}

	nw, n, m := len(word), len(board), len(board[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				var dfs func(index, x, y int) bool

				dfs = func(index, x, y int) bool {
					if x < 0 || x >= n || y < 0 || y >= m || board[x][y] == '#' || board[x][y] != word[index] {
						return false
					}
					if index == nw-1 {
						return true
					}
					for _, v := range directions {
						char := board[x][y]
						board[x][y] = '#'
						if dfs(index+1, x+v[0], y+v[1]) {
							return true
						}
						board[x][y] = char
					}
					return false
				}

				if dfs(0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

// @lc code=end

