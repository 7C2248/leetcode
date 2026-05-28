/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
import (
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {

	// 根据转移判断无后效性方向
	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n)
	dp[n-1] = min(0, dungeon[m-1][n-1])
	for i := m - 1; i > -1; i-- {
		for j := n - 1; j > -1; j-- {
			if i == m-1 && j == n-1 {
				continue
			}
			best := math.MinInt
			if i < m-1 {
				if dp[j] > best {
					best = dp[j]
				}
			}
			if j < n-1 {
				if dp[j+1] > best {
					best = dp[j+1]
				}
			}
			dp[j] = min(0, best+dungeon[i][j])

		}
	}
	if dp[0] > 0 {
		return 1
	} else {
		return -dp[0] + 1
	}
}

// @lc code=end

/*
[-2,-2],[-3,-3],[-3, 0]
[-5,-5],[-5,-5],[-3,-3]
[-9,-9],[-6,-6],[-3,-3]

[ -2, -3,  3] [-6,-6],[-4,-4],[-1,-1]
[ -5,-10,  1] [30,-5],[15,-10],[-4,-4]
[ 10, 30, -5] [35,0],[25,0],[-5,-5]
*/