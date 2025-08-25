/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

func generateParenthesis(n int) []string {
	//层次遍历
	result := []string{}
	temp := make([]int, 0, n)

	var dfs func(index int, maxlevel int)
	dfs = func(index int, maxlevel int) {
		if index == n {
			r, level := "", 0
			for j := 0; j < n; j++ {
				if temp[j] == level {
					level++
				} else {
					r += strings.Repeat(")", level-temp[j])
					level = temp[j] + 1
				}
				r += "("
			}
			r += strings.Repeat(")", level)
			result = append(result, r)
			return
		}

		for i := 0; i <= maxlevel; i++ {
			temp = append(temp, i)
			dfs(index+1, i+1)
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0, 0)
	return result
}

/*
func generateParenthesis(n int) []string {
	//DFS回溯

	var result []string

	var dfs func(path string, left, right int)
	dfs = func(path string, left, right int) {
		if left == n && right == n {
			result = append(result, path)
			return
		}

		if left < n {
			dfs(path+"(", left+1, right)
		}

		if right < left {
			dfs(path+")", left, right+1)
		}
	}

	dfs("", 0, 0)
	return result
}
*/
// @lc code=end

