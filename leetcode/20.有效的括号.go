/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {

	stack := make([]int, 0)
	kk := map[rune]int{'{': 0, '}': 1, '(': 2, ')': 3, '[': 4, ']': 5}

	for _, c := range s {
		if kk[c]%2 == 0 {
			stack = append(stack, kk[c])
		}
		if kk[c]%2 == 1 {
			if len(stack) == 0 {
				return false
			}
			for len(stack) > 0 {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if v != kk[c]-1 {
					return false
				}
				if v == kk[c]-1 {
					break
				}
			}
		}
	}

	return len(stack) == 0
}

// @lc code=end

