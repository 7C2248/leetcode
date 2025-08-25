/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N 皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {

	result := 0
	cmap := make([][]int, n)
	for i := 0; i < n; i++ {
		cmap[i] = make([]int, n)
	}
	var dfs func(level int)

	dfs = func(level int) {
		if level == n {
			result += 1
			return
		}

		for i := 0; i < n; i++ {

			if cmap[level][i] == 0 {
				for j := 0; j < n; j++ {
					if cmap[level][j] == 0 {
						cmap[level][j] = level + 1
					}
				}
				for count := 0; level+count < n; count++ {
					if cmap[level+count][i] == 0 {
						cmap[level+count][i] = level + 1
					}

					if i+count < n && cmap[level+count][i+count] == 0 {
						cmap[level+count][i+count] = level + 1
					}
					if i-count >= 0 && cmap[level+count][i-count] == 0 {
						cmap[level+count][i-count] = level + 1
					}
				}

				dfs(level + 1)

				for j := 0; j < n; j++ {
					if cmap[level][j] == level+1 {
						cmap[level][j] = 0
					}
				}
				for count := 0; level+count < n; count++ {
					if cmap[level+count][i] == level+1 {
						cmap[level+count][i] = 0
					}

					if i+count < n && cmap[level+count][i+count] == level+1 {
						cmap[level+count][i+count] = 0
					}
					if i-count >= 0 && cmap[level+count][i-count] == level+1 {
						cmap[level+count][i-count] = 0
					}
				}
			}
		}
	}
	dfs(0)
	return result
}

// @lc code=end

