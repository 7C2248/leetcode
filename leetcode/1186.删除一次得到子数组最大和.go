/*
 * @lc app=leetcode.cn id=1186 lang=golang
 *
 * [1186] 删除一次得到子数组最大和
 */

// @lc code=start
func maximumSum(arr []int) int {

	// 记住规划状态转移问题

	n := len(arr)
	noDel, haveDel := arr[0], arr[0]
	result := arr[0]
	for i := 1; i < n; i++ {
		noDel, haveDel = max(arr[i], arr[i]+noDel), max(noDel, arr[i]+haveDel)
		if noDel > result {
			result = noDel
		}
		if haveDel > result {
			result = haveDel
		}
	}
	return result
}

// @lc code=end