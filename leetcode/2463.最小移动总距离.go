/*
 * @lc app=leetcode.cn id=2463 lang=golang
 *
 * [2463] 最小移动总距离
 */

// @lc code=start
import (
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func minimumTotalDistance(robot []int, factory [][]int) int64 {

	/*
		由于需要总路径最短，
		所以在一个工厂中修理的机器人在排序后的队列中一定是相邻的，
		从而对于排序后的工厂和机器人可以转化为边长区间划分的问题
	*/
	nr, nf := len(robot), len(factory)
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })

	distence := make([][]int, nf)
	for i := 0; i < nf; i++ {
		distence[i] = make([]int, nr+1)
		for j := 1; j <= nr; j++ {
			distence[i][j] += abs(robot[j-1]-factory[i][0]) + distence[i][j-1]
		}
	}
	//前i个工厂修理前j个机器人时的最小移动总距离
	dp := make([][]int, nf+1)
	for i := 0; i <= nf; i++ {
		dp[i] = make([]int, nr+1)
		for j := 1; j <= nr; j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0

	for i := 0; i < nf; i++ {

		// 根据前一行进行转移
		for j, v := range dp[i] {
			if v != math.MaxInt {
				for k := 0; k <= factory[i][1] && j+k <= nr; k++ {
					dp[i+1][j+k] = min(dp[i+1][j+k], v+distence[i][j+k]-distence[i][j])
				}
			}
		}
	}
	return int64(dp[nf][nr])
}

// @lc code=end

