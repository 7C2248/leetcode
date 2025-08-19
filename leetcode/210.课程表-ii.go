/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {

	sign := true
	result := []int{}
	m := make(map[int]bool)

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

			if _, ok := m[course]; !ok {
				result = append(result, course)
				m[course] = true
			}

			return true
		}

		state[course] = 1
		for _, pre := range graph[course] {
			if !dfs(pre) {
				return false
			}
		}

		if _, ok := m[course]; !ok {
			result = append(result, course)
			m[course] = true
		}
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

