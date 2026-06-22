/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {

	// 排序后由于宽度相等时高度降序排列
	// 从而高度的最长递增子序列长度一定小于等于宽度的不同元素个数
	// 也即宽度的最长递增子序列

	dph := make([]int, 0)

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	for _, envelope := range envelopes {
		ch := envelope[1]
		pos := sort.Search(len(dph), func(r int) bool {
			return dph[r] >= ch
		})
		if pos == len(dph) {
			dph = append(dph, ch)
		} else {
			dph[pos] = ch
		}
	}
	return len(dph)
}

// @lc code=end