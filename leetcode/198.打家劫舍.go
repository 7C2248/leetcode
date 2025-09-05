/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
/*
func rob(nums []int) int {
	n := len(nums)
	mmax := make([]int, n+2)

	var dfs func(now int)
	dfs = func(now int) {
		if now > n-1 {
			return
		}
		dfs(now + 1)
		mmax[now] = max(mmax[now+2]+nums[now], mmax[now+1])
		return

	}
	dfs(0)
	return mmax[0]
}
*/
func rob(nums []int) int {
	n := len(nums)
	mmax := make([]int, n+2)

	for i := 2; i < n+2; i++ {
		mmax[i] = max(mmax[i-2]+nums[i-2], mmax[i-1])
	}
	return mmax[len(mmax)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

