/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func min(a int, b int) (int, int) {
	if a > b {
		return b, 1
	} else {
		return a, 0
	}
}
func maxArea(height []int) int {

	i, j, max := 0, len(height)-1, 0
	for i <= j {
		x, poi := min(height[i], height[j])
		if x*(j-i) > max {
			max = x * (j - i)
		}
		if poi == 1 {
			j--
		} else {
			i++
		}

	}
	return max
}

// @lc code=end

