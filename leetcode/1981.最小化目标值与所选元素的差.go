/*
 * @lc app=leetcode.cn id=1981 lang=golang
 *
 * [1981] 最小化目标值与所选元素的差
 */

// @lc code=start
func minimizeTheDifference(mat [][]int, target int) int {

	maxSum := 0
	for _, vi := range mat {
		temp := 0
		for _, vj := range vi {
			temp = max(temp, vj)
		}
		maxSum += temp
	}
	pre := make([]bool, maxSum+1)

	pre[0] = true
	for _, vi := range mat {
		now := make([]bool, maxSum+1)
		for _, vj := range vi {
			for i := vj; i <= maxSum; i++ {
				if pre[i-vj] {
					now[i] = true
				}
			}
		}
		pre = now

	}
	maximum := math.MaxInt
	for k, v := range pre {
		if v {
			maximum = min(maximum, abs(k-target))
		}
	}

	return maximum
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

