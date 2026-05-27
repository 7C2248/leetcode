/*
 * @lc app=leetcode.cn id=740 lang=golang
 *
 * [740] 删除并获得点数
 */

// @lc code=start
import (
	"sort"
)

func deleteAndEarn(nums []int) int {
	hmap := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if _, ok := hmap[nums[i]]; !ok {
			hmap[nums[i]] = 1
		} else {
			hmap[nums[i]] += 1
		}
	}
	points := []int{}
	for k, _ := range hmap {
		points = append(points, k)
	}
	sort.Slice(points, func(i, j int) bool { return points[i] < points[j] })

	n = len(points)
	if n == 1 {
		return points[0] * hmap[points[0]]
	}
	dp := make([]int, n)
	dp[0] = hmap[points[0]] * points[0]
	for i := 1; i < n; i++ {
		if i == 1 {
			if points[i] == points[i-1]+1 {
				dp[i] = max(dp[i-1], points[i]*hmap[points[i]])
			} else {
				dp[i] = dp[i-1] + points[i]*hmap[points[i]]
			}
		} else {
			if points[i] == points[i-1]+1 {
				dp[i] = max(dp[i-2]+points[i]*hmap[points[i]], dp[i-1])
			} else {
				dp[i] = dp[i-1] + points[i]*hmap[points[i]]
			}
		}
	}
	return dp[n-1]
}

// @lc code=end
