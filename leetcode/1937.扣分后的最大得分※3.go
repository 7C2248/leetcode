/*
 * @lc app=leetcode.cn id=1937 lang=golang
 *
 * [1937] 扣分后的最大得分
 */

// @lc code=start
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	pre := make([]int, n)
	cur := make([]int, n)
	cur[0] = points[0][0]
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	for i := 0; i < m; i++ {
		// 关键优化，通过讨论绝对值，分别计算第j列对应的上一行的
		// 左右最大值，将当前行的位置从最值比较中剔除，
		// 从而只需要O(n)便可以计算当前列的最大值。
		leftMax[0] = pre[0] + 0
		for j := 1; j < n; j++ {
			leftMax[j] = max(leftMax[j-1], pre[j]+j)
		}
		rightMax[n-1] = pre[n-1] - n + 1
		for j := n - 2; j >= 0; j-- {
			rightMax[j] = max(rightMax[j+1], pre[j]-j)
		}
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			cur[j] = points[i][j] + max(leftMax[j]-j, rightMax[j]+j)
		}
		cur, pre = pre, cur
	}
	maximum := 0
	for _, v := range pre {
		maximum = max(maximum, v)
	}
	return int64(maximum)
}

// @lc code=end

