/*
 * @lc app=leetcode.cn id=3186 lang=golang
 *
 * [3186] 施咒的最大总伤害
 */

// @lc code=start
import (
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	hmap := map[int]int{}
	n := len(power)
	for i := 0; i < n; i++ {
		if _, ok := hmap[power[i]]; !ok {
			hmap[power[i]] = 1
		} else {
			hmap[power[i]] += 1
		}
	}

	points := []int{}
	for k, _ := range hmap {
		points = append(points, k)
	}
	n = len(points)
	if n == 1 {
		return int64(points[0] * hmap[points[0]])
	}
	//sort.Ints(points)
	sort.Slice(points, func(i, j int) bool { return points[i] < points[j] })

	dp := make([]int, n)
	dp[0] = hmap[points[0]] * points[0]

	pre := -1 // 最后一个 points[pre] <= points[i] - 3
	for i := 1; i < n; i++ {
		// 移动 pre 指针，找到最靠右的不冲突位置
		for pre+1 < i && points[pre+1] <= points[i]-3 {
			pre++
		}

		curVal := points[i] * hmap[points[i]]
		// 不选当前值
		take := dp[i-1]
		// 选当前值：如果有不冲突的前驱，则加上其 dp 值
		if pre >= 0 {
			take = max(take, dp[pre]+curVal)
		} else {
			take = max(take, curVal)
		}
		dp[i] = take
	}
	return int64(dp[n-1])
}

// @lc code=end

