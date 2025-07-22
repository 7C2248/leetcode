/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {

	l := len(matrix)

	for i := 0; 2*i <= l; i++ {
		ne := l - i*2

		for k := 0; k < ne-1; k++ {
			m, n := i, i
			for j := 0; j < 2*ne+2*(ne-2)-1; j++ {

				var di, dj int = 0, 0
				if j < ne-1 {
					di = 1
					dj = 0
				} else if j < 2*(ne-1) {
					di = 0
					dj = 1
				} else if j < 3*(ne-1) {
					di = -1
					dj = 0
				} else {
					di = 0
					dj = -1
				}
				matrix[m][n], matrix[m+di][n+dj] = matrix[m+di][n+dj], matrix[m][n]
				m += di
				n += dj
			}
		}
	}
}

// @lc code=end

