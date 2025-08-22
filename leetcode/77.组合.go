/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	result := [][]int{}

	temp := []int{0}
	var dfs func(i int)
	dfs = func(i int) {
		if i == k {
			r := make([]int, i)
			copy(r, temp[1:])
			result = append(result, r)
			return
		}

		for j := temp[len(temp)-1] + 1; j <= n-(k-i)+1; j++ {
			temp = append(temp, j)
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0)
	return result
}

// @lc code=end

