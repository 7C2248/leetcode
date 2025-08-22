/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	n := len(nums)
	result := [][]int{}
	temp := []int{}
	visited := make([]bool, n)

	var dfs func()
	dfs = func() {
		if len(temp) == n {
			r := make([]int, n)
			copy(r, temp)
			result = append(result, r)
			return
		}
		for i := 0; i < n; i++ {
			if !visited[i] {
				visited[i] = true
				temp = append(temp, nums[i])
				dfs()
				temp = temp[:len(temp)-1]
				visited[i] = false
			}
		}
	}

	dfs()
	return result
}

// @lc code=end

