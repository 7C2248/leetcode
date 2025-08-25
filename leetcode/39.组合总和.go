/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	n := len(candidates)
	temp := []int{}

	sort.Ints(candidates)

	var dfs func(target int, index int)
	dfs = func(target int, index int) {
		if target == 0 {
			r := make([]int, len(temp))
			copy(r, temp)
			result = append(result, r)
			return
		}

		for i := index; i < n; i++ {
			if candidates[i] > target {
				break
			}

			temp = append(temp, candidates[i])
			dfs(target-candidates[i], i)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(target, 0)
	return result
}

// @lc code=end

