/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		// 必须倒序，防止同一 num 被重复使用
		for s := target; s >= num; s-- {
			if dp[s-num] {
				dp[s] = true
			}
		}
	}
	return dp[target]
}

// @lc code=end

