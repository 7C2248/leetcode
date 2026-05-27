import (
	"math"
)

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 和hashmap有近乎100倍的性能差距
	dp := make([][1024]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = [1024]bool{}
	}
	dp[0][grid[0][0]] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp := [1024]bool{}
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				for index, is := range dp[j-1] {
					if is {
						temp[index^grid[i][j]] = true
					}
				}
				dp[j] = temp
				continue
			}
			if j == 0 {
				for index, is := range dp[j] {
					if is {
						temp[index^grid[i][j]] = true
					}
				}
				dp[j] = temp
				continue
			}
			for index, is := range dp[j] {
				if is {
					temp[index^grid[i][j]] = true
				}
			}
			for index, is := range dp[j-1] {
				if is {
					temp[index^grid[i][j]] = true
				}
			}

			dp[j] = temp
		}
	}
	minimum := math.MaxInt
	for index, is := range dp[n-1] {
		if is {
			minimum = min(minimum, index)
		}

	}
	return minimum
}