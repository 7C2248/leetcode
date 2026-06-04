/*
 * @lc app=leetcode.cn id=1774 lang=golang
 *
 * [1774] 最接近目标价格的甜点成本
 */

// @lc code=start
import (
	"math"
)

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	dp := map[int]bool{0: true}
	for _, vt := range toppingCosts {
		for i := 0; i < 2; i++ {
			keys := make([]int, 0, len(dp))
			for k := range dp {
				keys = append(keys, k)
			}
			for _, k := range keys {
				dp[k+vt] = true
			}
		}
	}
	result := math.MaxInt
	for _, vb := range baseCosts {
		for k, _ := range dp {
			if abs(k+vb-target) < abs(result-target) {
				result = k + vb
			} else if abs(k+vb-target) == abs(result-target) {
				if k+vb < result {
					result = k + vb
				}
			}
		}
	}
	return result
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end