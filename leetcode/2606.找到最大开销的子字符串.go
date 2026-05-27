/*
 * @lc app=leetcode.cn id=2606 lang=golang
 *
 * [2606] 找到最大开销的子字符串
 */

// @lc code=start
func maximumCostSubstring(s string, chars string, vals []int) int {
	// 固定范围映射
	value := [26]int{}
	for i := 0; i < 26; i++ {
		value[i] = i + 1
	}
	for i := 0; i < len(chars); i++ {
		value[chars[i]-'a'] = vals[i]
	}

	// 滚动变量
	cur, result := 0, 0
	for _, ch := range s {
		cur = max(cur, 0) + value[ch-'a']
		result = max(result, cur)
	}
	return result
}

/*
func maximumCostSubstring(s string, chars string, vals []int) int {
	n := len(chars)
	hmap := map[byte]int{}
	for i := 0; i < n; i++ {
		hmap[chars[i]] = vals[i]
	}

	n = len(s)
	dp := make([]int, n)
	result := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			if v, ok := hmap[s[i]]; ok {
				dp[i] = v
			} else {
				dp[i] = int(s[i] - 'a' + 1)
			}
		} else {
			if v, ok := hmap[s[i]]; ok {
				dp[i] = max(dp[i-1], 0) + v
			} else {
				dp[i] = max(dp[i-1], 0) + int(s[i]-'a'+1)
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}
*/
// @lc code=end

