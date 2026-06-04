/*
 * @lc app=leetcode.cn id=1744 lang=golang
 *
 * [1744] 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
 */

// @lc code=start
func canEat(candiesCount []int, queries [][]int) []bool {
	for i := 1; i < len(candiesCount); i++ {
		candiesCount[i] += candiesCount[i-1]
	}

	ans := make([]bool, len(queries))
	for k, q := range queries {
		ft, fd, dc := q[0], q[1], q[2]
		least := 1
		if ft > 0 {
			least = candiesCount[ft-1] + 1
		}
		most := candiesCount[ft]

		ans[k] = (least <= (fd+1)*dc) && (most >= (fd + 1))
	}
	return ans
}

// @lc code=end

