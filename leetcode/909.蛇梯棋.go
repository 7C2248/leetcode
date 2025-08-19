/*
 * @lc app=leetcode.cn id=909 lang=golang
 *
 * [909] 蛇梯棋
 */

// @lc code=start
func snakesAndLadders(board [][]int) int {
	n := len(board)
	end := n*n - 1

	for i := 0; i < n/2; i++ {
		board[i], board[n-i-1] = board[n-i-1], board[i]
	}

	queue := [][]int{{0, 0}}
	visited := make([]bool, n*n)
	visited[0] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		pos, steps := current[0], current[1]

		if pos == end {
			return steps
		}

		for j := 1; j <= 6; j++ {
			target := pos + j

			if target > end {
				continue
			}

			y, xm := target/n, target%n
			x := xm
			if y%2 == 1 {
				x = n - xm - 1
			}

			if board[y][x] != -1 {
				target = board[y][x] - 1
			}

			if !visited[target] {
				visited[target] = true
				queue = append(queue, []int{target, steps + 1})
			}
		}
	}

	return -1
}

// @lc code=end

