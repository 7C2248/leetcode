/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {

	sign := true
	result := []int{}

	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}

	state := make([]int, numCourses)

	var dfs func(int) bool
	dfs = func(course int) bool {
		if state[course] == 1 {
			return false
		}
		if state[course] == 2 {
			return true
		}

		state[course] = 1
		for _, pre := range graph[course] {
			if !dfs(pre) {
				return false
			}
		}

		// 首次遍历时如果能完成该课程则说明该课程之前的都已经完成，所以直接添加到结果后面。
		result = append(result, course)
		state[course] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			sign = false
		}
	}

	if sign {
		return result
	} else {
		return []int{}
	}

}

// @lc code=end

