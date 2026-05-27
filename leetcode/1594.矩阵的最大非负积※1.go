/*
 * @lc app=leetcode.cn id=1594 lang=golang
 *
 * [1594] 矩阵的最大非负积
 */

// @lc code=start
func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]map[int]bool, n)

	var mod int = 1e9 + 7
	dp[0] = map[int]bool{grid[0][0]: true}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			var up, left map[int]bool = nil, nil
			if i > 0 {
				up = dp[j]
			}
			if j > 0 {
				left = dp[j-1]
			}

			temp := map[int]bool{}
			if up != nil {
				for k, _ := range up {
					temp[k*grid[i][j]] = true
				}
			}
			if left != nil {
				for k, _ := range left {
					temp[k*grid[i][j]] = true
				}
			}
			dp[j] = temp
		}
	}
	maximum := -1
	for k, _ := range dp[n-1] {
		maximum = max(maximum, k)
	}
	return maximum % mod
}

/*
//只存最值优化
func maxProductPath(grid [][]int) int {
    const MOD = 1_000_000_007
    m, n := len(grid), len(grid[0])

    // dpMax[j], dpMin[j] 表示到达当前行第 j 列的最大/最小乘积
    dpMax := make([]int64, n)
    dpMin := make([]int64, n)

    // 初始化第一行
    dpMax[0] = int64(grid[0][0])
    dpMin[0] = int64(grid[0][0])
    for j := 1; j < n; j++ {
        dpMax[j] = dpMax[j-1] * int64(grid[0][j])
        dpMin[j] = dpMax[j]  // 第一行只能向右，最大最小相同
    }

    for i := 1; i < m; i++ {
        // 保存上一行的值，用于第一列更新
        prevMax := dpMax[0]
        prevMin := dpMin[0]
        // 当前行第一列只能从上方来
        val := int64(grid[i][0])
        candidates := []int64{
            prevMax * val,
            prevMin * val,
        }
        dpMax[0] = max(candidates...)
        dpMin[0] = min(candidates...)

        for j := 1; j < n; j++ {
            val := int64(grid[i][j])
            // 合并上方和左方的来源
            upMax, upMin := dpMax[j], dpMin[j]
            leftMax, leftMin := dpMax[j-1], dpMin[j-1]

            // 所有可能的乘积候选
            candidates := []int64{
                upMax * val,
                upMin * val,
                leftMax * val,
                leftMin * val,
            }
            dpMax[j] = max(candidates...)
            dpMin[j] = min(candidates...)
        }
    }

    if dpMax[n-1] < 0 {
        return -1
    }
    return int(dpMax[n-1] % MOD)
}

// 辅助函数
func max(vals ...int64) int64 {
    ans := vals[0]
    for _, v := range vals[1:] {
        if v > ans {
            ans = v
        }
    }
    return ans
}

func min(vals ...int64) int64 {
    ans := vals[0]
    for _, v := range vals[1:] {
        if v < ans {
            ans = v
        }
    }
    return ans
}
*/

// @lc code=end

