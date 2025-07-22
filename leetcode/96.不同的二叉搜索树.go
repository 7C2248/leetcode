/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start

func numTrees(n int) int {

	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for k := 3; k <= n; k++ {
		sum := 0
		for i := 0; i < k; i++ {
			small := i
			big := k - i - 1
			lefttreenum := dp[small]
			righttreenum := dp[big]
			if (lefttreenum > 0) && (righttreenum > 0) {
				sum += lefttreenum * righttreenum
			} else {
				if lefttreenum > 0 {
					sum += lefttreenum
				} else {
					sum += righttreenum
				}
			}
		}
		dp[k] = sum
	}
	return dp[n]
}

// @lc code=end

