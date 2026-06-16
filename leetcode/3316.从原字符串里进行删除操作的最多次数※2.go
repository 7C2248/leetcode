/*
 * @lc app=leetcode.cn id=3316 lang=golang
 *
 * [3316] 从原字符串里进行删除操作的最多次数
 */

// @lc code=start
func maxRemovals(source, pattern string, targetIndices []int) int {
	m := len(pattern)
	//恰好匹配 pattern 的前 j 个字符的前提下，能删除的最大字符数
	f := make([]int, m+1)
	for i := 1; i <= m; i++ {
		f[i] = math.MinInt
	}
	k := 0
	for i := range source {
		if k < len(targetIndices) && targetIndices[k] < i {
			k++
		}
		isDel := 0
		if k < len(targetIndices) && targetIndices[k] == i {
			isDel = 1
		}
		for j := min(i, m-1); j >= 0; j-- {
			f[j+1] += isDel
			if source[i] == pattern[j] {
				f[j+1] = max(f[j+1], f[j])
			}
		}
		f[0] += isDel
	}
	return f[m]
}

// @lc code=end