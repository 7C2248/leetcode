/*
 * @lc app=leetcode.cn id=2267 lang=golang
 *
 * [2267] 检查是否有合法括号字符串路径
 */

// @lc code=start
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, m+n+1)
	}
	turnToInt := func(a byte) int {
		if a == '(' {
			return 1
		}
		if a == ')' {
			return -1
		}
		return 0
	}
	start := turnToInt(grid[0][0])
	if start < 0 {
		return false
	}
	dp[0][start] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			temp := make([]bool, m+n+1)
			cur := turnToInt(grid[i][j])
			if i > 0 {
				for index, val := range dp[j] {
					if val && index+cur > -1 {
						temp[index+cur] = true
					}
				}
			}
			if j > 0 {
				for index, val := range dp[j-1] {
					if val && index+cur > -1 {
						temp[index+cur] = true
					}
				}
			}
			dp[j] = temp
		}
	}
	return dp[n-1][0]
}

// @lc code=end

/*

// 位运算优化版本
func hasValidPath(grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    maxBal := m + n                      // 最大平衡值索引
    words := (maxBal + 64) / 64          // 确保容纳位 0..maxBal

    prev := make([][]uint64, n)          // 上一行，初始全零
    cur := make([][]uint64, n)           // 当前行
    for j := 0; j < n; j++ {
        prev[j] = make([]uint64, words)
        cur[j] = make([]uint64, words)
    }

    if grid[0][0] == ')' {
        return false
    }
    setBit(cur[0], 1)                    // 起点状态放入 cur[0]

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                continue
            }
            bits := make([]uint64, words)
            if i > 0 {
                orBitsInPlace(bits, prev[j])   // 上方
            }
            if j > 0 {
                orBitsInPlace(bits, cur[j-1])  // 左方
            }

            if grid[i][j] == '(' {
                shiftLeftInPlace(bits)
            } else {
                shiftRightInPlace(bits)
            }
            copy(cur[j], bits)
        }
        // 滚动：当前行变上一行，清零下一行
        prev, cur = cur, prev
        for j := 0; j < n; j++ {
            for w := 0; w < words; w++ {
                cur[j][w] = 0
            }
        }
    }
    return hasBit(prev[n-1], 0)
}

// ---------- 辅助函数（原地操作，减少分配）----------
func setBit(bits []uint64, pos int) {
    bits[pos/64] |= 1 << (pos % 64)
}

func hasBit(bits []uint64, pos int) bool {
    return bits[pos/64]&(1<<(pos%64)) != 0
}

func orBitsInPlace(a, b []uint64) {
    for i := range a {
        a[i] |= b[i]
    }
}

func shiftLeftInPlace(bits []uint64) {
    carry := uint64(0)
    for i := 0; i < len(bits); i++ {
        newCarry := bits[i] >> 63
        bits[i] = (bits[i] << 1) | carry
        carry = newCarry
    }
}

func shiftRightInPlace(bits []uint64) {
    carry := uint64(0)
    for i := len(bits) - 1; i >= 0; i-- {
        newCarry := bits[i] & 1
        bits[i] = (bits[i] >> 1) | (carry << 63)
        carry = newCarry
    }
}
*/