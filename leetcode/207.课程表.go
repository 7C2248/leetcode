/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start

func canFinish(numCourses int, prerequisites [][]int) bool {
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
		state[course] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

// @lc code=end

