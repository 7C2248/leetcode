/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
import (
	"math"
)

func lastStoneWeightII(stones []int) int {
	result := math.MaxInt
	sum := getSum(stones)
	var target int = sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, v := range stones {
		for s := target; s >= v; s-- {
			if dp[s-v] {
				dp[s] = true
			}
		}
	}
	for k, v := range dp {
		if v {
			result = min(result, abs(sum-k*2))
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
func getSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// @lc code=end