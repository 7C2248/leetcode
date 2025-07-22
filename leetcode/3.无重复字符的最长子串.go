/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	maxLen := 0
	i := 0
	for j := 0; j < len(s); j++ {
		pos := strings.IndexByte(s[i:j], s[j])
		if pos != -1 {
			i += pos + 1
		}
		if j-i+1 > maxLen {
			maxLen = j - i + 1
		}
	}
	return maxLen
}

// @lc code=end

