/*
 * @lc app=leetcode.cn id=1301 lang=golang
 *
 * [1301] 最大得分的路径数目
 */

// @lc code=start
func pathsWithMaxScore(board []string) []int {
	const mod = 1_000_000_007
	n := len(board)

	now := make([][]int, n)
	pre := make([][]int, n)
	for i := 0; i < n; i++ {
		now[i] = make([]int, 2)
		pre[i] = []int{-1, 0}
	}

	now[n-1][1] = 1

	// 辅助函数：获取格子的数字得分
	getScore := func(ch byte) int {
		if ch == 'E' || ch == 'S' {
			return 0
		}
		return int(ch - '0')
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 起点已经初始化过
			if i == n-1 && j == n-1 {
				continue
			}
			// 障碍物不可达
			if board[i][j] == 'X' {
				now[j][0] = -1 // 标记为不可达
				continue
			}

			// 从三个方向收集候选
			var down, right, downR []int = nil, nil, nil
			if i < n-1 {
				down = pre[j]
			}
			if j < n-1 {
				right = now[j+1]
			}
			if i < n-1 && j < n-1 {
				downR = pre[j+1]
			}

			best := -1
			totalWays := 0
			if down != nil {
				if down[0] > best {
					best = down[0]
					totalWays = down[1]
				} else if down[0] == best {
					totalWays += down[1]
					totalWays %= mod
				}
			}
			if right != nil {
				if right[0] > best {
					best = right[0]
					totalWays = right[1]
				} else if right[0] == best {
					totalWays += right[1]
					totalWays %= mod
				}
			}
			if downR != nil {
				if downR[0] > best {
					best = downR[0]
					totalWays = downR[1]
				} else if downR[0] == best {
					totalWays += downR[1]
					totalWays %= mod
				}
			}
			if best == -1 {
				now[j][0] = -1
				continue
			}

			now[j][0] = best + getScore(board[i][j])
			now[j][1] = totalWays
		}
		now, pre = pre, now
	}

	if pre[0][0] == -1 {
		return []int{0, 0}
	}
	return pre[0]
}

// @lc code=end
