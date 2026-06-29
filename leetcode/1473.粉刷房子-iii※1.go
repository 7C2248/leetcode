/*
 * @lc app=leetcode.cn id=1473 lang=golang
 *
 * [1473] 粉刷房子 III
 */

// @lc code=start
func minCost(houses []int, cost [][]int, m int, n int, target int) int {

	if n == 1 && target == 1 {
		sum := 0
		for i := 0; i < m; i++ {
			if houses[i] == 0 {
				sum += cost[i][0]
			}
		}
		return sum
	}

	pre := make([][]int, m+1)
	now := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		pre[i] = make([]int, n)
		now[i] = make([]int, n)
		for j := 0; j < n; j++ {
			pre[i][j] = math.MaxInt
			now[i][j] = math.MaxInt
		}
	}
	for i := 0; i < n; i++ {
		pre[0][i] = 0
	}

	for i := 0; i < target; i++ {
		fmt.Println(pre)
		for j := 1; j <= m; j++ {

			for cc := 0; cc < n; cc++ {
				curCost := 0
				for k := j - 1; k > -1; k-- {
					if houses[k] == 0 {
						curCost += cost[k][cc]
						for pc := 0; pc < n; pc++ {
							if pc == cc {
								continue
							}
							if pre[k][pc] != math.MaxInt {
								now[j][cc] = min(now[j][cc], pre[k][pc]+curCost)
							}
						}
					} else if houses[k] == cc+1 {
						for pc := 0; pc < n; pc++ {
							if pc == cc {
								continue
							}
							if pre[k][pc] != math.MaxInt {
								now[j][cc] = min(now[j][cc], pre[k][pc]+curCost)
							}
						}
					} else {
						for pc := 0; pc < n; pc++ {
							if pc == houses[k]-1 {
								continue
							}
							if pre[k+1][pc] != math.MaxInt {
								now[j][houses[k]-1] = min(now[j][houses[k]-1], pre[k+1][pc]+curCost)
							}
						}
						break
					}
				}
			}

		}
		pre, now = now, pre
		for i := 0; i <= m; i++ {
			for j := 0; j < n; j++ {
				now[i][j] = math.MaxInt
			}
		}
	}
	fmt.Println(pre)
	minimum := math.MaxInt
	for i := 0; i < n; i++ {
		minimum = min(minimum, pre[m][i])
	}
	if minimum == math.MaxInt {
		return -1
	}
	return minimum
}

// @lc code=end

/*
	// 由于多了颜色的维度，枚举区间的复杂度会很高。
	// 此外由于无论如何都要枚举颜色，所以不如将区间扫描换位单位房子扫描
	// 若当前房子颜色和之前不同，则记为新段，否则从旧段更新。
	// 时间复杂度从O(target*m^2 * n^2)降低到O(target*m*n)
func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	const INF = math.MaxInt32 / 2

	// prev[c][k] 滚动数组
	prev := make([][]int, n)
	cur := make([][]int, n)
	for c := 0; c < n; c++ {
		prev[c] = make([]int, target+1)
		cur[c] = make([]int, target+1)
		for k := 0; k <= target; k++ {
			prev[c][k] = INF
			cur[c][k] = INF
		}
	}

	// 初始化：0 个房子，0 个街区，成本 0（可以把“前0个房子”作为基准）
	// 虚拟一个颜色 -1，但这里用另一种方式：
	// 处理第 1 个房子时，从“0 个房子，街区数 0”出发。
	// 我们可以在循环里特殊处理 i == 1 的情况。
	for i := 1; i <= m; i++ {
		// 重置 cur
		for c := 0; c < n; c++ {
			for k := 0; k <= target; k++ {
				cur[c][k] = INF
			}
		}

		// 若当前房子已涂色，则只允许固定颜色
		colors := []int{}
		if houses[i-1] != 0 {
			colors = append(colors, houses[i-1]-1)
		} else {
			for c := 0; c < n; c++ {
				colors = append(colors, c)
			}
		}

		// 提前计算 prev 中每个 k 的 (最小值，次小值) 及其颜色
		type minPair struct {
			min1, min2 int
			c1, c2     int
		}
		minForK := make([]minPair, target+1)
		for k := 0; k <= target; k++ {
			m1, m2 := INF, INF
			c1, c2 := -1, -1
			for c := 0; c < n; c++ {
				v := prev[c][k]
				if v < m1 {
					m2, c2 = m1, c1
					m1, c1 = v, c
				} else if v < m2 {
					m2, c2 = v, c
				}
			}
			minForK[k] = minPair{m1, m2, c1, c2}
		}

		for _, c := range colors {
			paintCost := 0
			if houses[i-1] == 0 {
				paintCost = cost[i-1][c]
			}

			// 街区数至少为 1
			for k := 1; k <= target && k <= i; k++ {
				// 情况 1：和前一房子同色，街区数不变
				same := prev[c][k]

				// 情况 2：和前一房子不同色，街区数 k-1
				diff := INF
				if k-1 >= 0 {
					pair := minForK[k-1]
					if pair.c1 != c {
						diff = pair.min1
					} else {
						diff = pair.min2
					}
				}
				// 如果是第 1 个房子且 k == 1，不存在“前一个房子”，
				// 我们可以认为 prev 的 k=0 时成本为 0，且街区从 0 到 1 视为不同色触发。
				// 这里统一：当 k == 1 时，从“0 个街区”转移。
				if k == 1 {
					// 第一个街区，前面没有房子，可以认为是“颜色不同，从 0 街区来”
					// 我们只使用 diff 并修正为从 0 开始
					if i == 1 {
						diff = 0 // 前 0 个房子，0 街区，成本 0
					}
				}

				best := same
				if diff < best {
					best = diff
				}
				if best != INF {
					cur[c][k] = best + paintCost
				}
			}
		}
		prev, cur = cur, prev
	}

	ans := INF
	for c := 0; c < n; c++ {
		if prev[c][target] < ans {
			ans = prev[c][target]
		}
	}
	if ans == INF {
		return -1
	}
	return ans
}
*/