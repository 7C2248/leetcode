/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	m, n := len(matrix), len(matrix[0])

	li, ri := 0, m*n-1

	for li < ri {
		hi := (li + ri) >> 1
		y, x := hi/n, hi%n

		if matrix[y][x] > target {
			ri = hi - 1
			continue
		}
		if matrix[y][x] < target {
			li = hi + 1
			continue
		}
		return true
	}

	if li == ri {
		y, x := li/n, li%n
		if matrix[y][x] == target {
			return true
		}
	}
	return false
}

// @lc code=end

