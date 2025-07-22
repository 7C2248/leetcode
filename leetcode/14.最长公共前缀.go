/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {

	minlen := 201
	tempi := 0
	for i, str := range strs {
		if str == "" {
			return ""
		}
		if len(str) < minlen {
			minlen = len(str)
			tempi = i
		}
	}

	pre := ""
	k := 0
	for {
		if k == len(strs[tempi]) {
			return pre[0:k]
		}
		temp := strs[tempi][k]
		for i := 0; i < len(strs); i++ {
			if strs[i][k] != temp {
				return pre[0:k]
			}

		}
		k += 1
		pre += string(temp)
	}
}

// @lc code=end

