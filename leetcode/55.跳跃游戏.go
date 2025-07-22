/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {

	n, max := len(nums), 0
	for i := 0; i < n; i++ {
		if i <= max {
			if i+nums[i] > max {
				max = i + nums[i]
			}
		}
	}

	return max >= n-1
}

// @lc code=end

