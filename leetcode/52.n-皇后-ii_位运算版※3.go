/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N 皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {
	if n <= 0 {
		return 0
	}
	limit := uint((1 << n) - 1) // 低 n 位有效
	var dfs func(cols, diagL, diagR uint) int

	dfs = func(cols, diagL, diagR uint) int {
		if cols == limit { // n 个皇后都放好了（每行正好放 1 个 → 列数也到位）
			return 1
		}
		count := 0
		// 本行所有可用列位（1 表示可放）
		available := limit &^ (cols | diagL | diagR)

		for available != 0 {
			// 取最低位的 1（选择一个列位置）
			pos := available & -available
			// 从候选中移除这个位置（相当于“试过这个分支”）
			available -= pos

			// 进入下一行：列占用累加；两条对角线分别左移/右移一位
			count += dfs(
				cols|pos,
				(diagL|pos)<<1,
				(diagR|pos)>>1,
			)
			// 不需要回滚：因为 cols/diagL/diagR 是值传递，上一层的值没变
		}
		return count
	}
	return dfs(0, 0, 0)
}

// @lc code=end

