/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start

func minimumTotal(triangle [][]int) int {

	var mmax int = 10e4 + 1
	n := len(triangle)
	mem := make([]int, n)
	for i := 0; i < n; i++ {
		mem[i] = mmax
	}
	mem[0] = triangle[0][0]

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			mem[j] = min(mem[j], mem[j-1]) + triangle[i][j]
		}
		mem[0] = mem[0] + triangle[i][0]
	}

	mmin := mem[0]
	for i := 1; i < n; i++ {
		if mem[i] < mmin {
			mmin = mem[i]
		}
	}

	return mmin
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

