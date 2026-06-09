/*
 * @lc app=leetcode.cn id=1449 lang=golang
 *
 * [1449] 数位成本和为目标值的最大数字
 */

// @lc code=start
func maxStr(a, b string) string {
	la, lb := len(a), len(b)
	if la > lb {
		return a
	} else if lb > la {
		return b
	}
	if a == "#" {
		return b
	}
	if b == "#" {
		return a
	}
	for i := 0; i < la; i++ {
		if a[i]-'0' > b[i]-'0' {
			return a
		} else if a[i]-'0' < b[i]-'0' {
			return b
		}
	}
	return a
}
func largestNumber(cost []int, target int) string {

	dp := make([]string, target+1)
	for i := 1; i < target; i++ {
		dp[i] = "#"
	}
	dp[target] = "0"

	for k, v := range cost {
		for i := v; i <= target; i++ {
			if dp[i-v] != "#" {
				dp[i] = maxStr(dp[i], maxStr(dp[i-v]+string(k+1+'0'), string(k+1+'0')+dp[i-v]))
			}
		}
	}
	return dp[target]
}

// @lc code=end

/*
//贪心优化
func largestNumber(cost []int, target int) string {
    // 1. 长度DP
    dp := make([]int, target+1)
    for i := 1; i <= target; i++ {
        dp[i] = -1
    }
    dp[0] = 0

    for d := 1; d <= 9; d++ {
        c := cost[d-1]
        for i := c; i <= target; i++ {
            if dp[i-c] != -1 && dp[i-c]+1 > dp[i] {
                dp[i] = dp[i-c] + 1
            }
        }
    }

    if dp[target] < 0 {
        return "0"
    }

    // 2. 贪心构造结果
    res := make([]byte, 0, dp[target])
    remain := target
    for remain > 0 {
        for d := 9; d >= 1; d-- {
            c := cost[d-1]
            if c <= remain && dp[remain] == dp[remain-c]+1 {
                res = append(res, byte('0'+d))
                remain -= c
                break
            }
        }
    }
    return string(res)
}
*/