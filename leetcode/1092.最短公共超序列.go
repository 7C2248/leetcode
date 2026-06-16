/*
 * @lc app=leetcode.cn id=1092 lang=golang
 *
 * [1092] 最短公共超序列
 */

// @lc code=start
func shortestCommonSupersequence(str1 string, str2 string) string {

	l1, l2 := len(str1), len(str2)
	if l1 < l2 {
		l1, l2 = l2, l1
		str1, str2 = str2, str1
	}
	// 字符串存储时间开销大
	dp := make([]string, l2+1)
	for i := 1; i <= l1; i++ {
		prev := dp[0]
		for j := 1; j <= l2; j++ {
			temp := dp[j]
			if str2[j-1] == str1[i-1] {
				dp[j] = prev + string(str2[j-1])
			} else {
				if len(dp[j-1]) > len(dp[j]) {
					dp[j] = dp[j-1]
				}
			}
			prev = temp
		}
	}

	LCS := dp[l2]
	lcs := len(dp[l2])

	result := ""
	i, j, k := 0, 0, 0
	for (i < l1 || j < l2) && k < lcs {
		for i < l1 && str1[i] != LCS[k] {
			result += string(str1[i])
			i += 1
		}
		for j < l2 && str2[j] != LCS[k] {
			result += string(str2[j])
			j += 1
		}
		result += string(LCS[k])
		k += 1
		i += 1
		j += 1
	}
	if i < l1 {
		result += str1[i:l1]
	}
	if j < l2 {
		result += str2[j:l2]
	}
	return result
}

// @lc code=end

/*
func shortestCommonSupersequence(str1 string, str2 string) string {
    l1, l2 := len(str1), len(str2)

    // 二维 int dp，只存长度
    dp := make([][]int, l1+1)
    for i := range dp {
        dp[i] = make([]int, l2+1)
    }

    for i := 1; i <= l1; i++ {
        for j := 1; j <= l2; j++ {
            if str1[i-1] == str2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                if dp[i-1][j] > dp[i][j-1] {
                    dp[i][j] = dp[i-1][j]
                } else {
                    dp[i][j] = dp[i][j-1]
                }
            }
        }
    }

    // 反向遍历dp构造 SCS
    var sb strings.Builder
    i, j := l1, l2
    for i > 0 || j > 0 {
        if i == 0 {
            sb.WriteByte(str2[j-1])
            j--
        } else if j == 0 {
            sb.WriteByte(str1[i-1])
            i--
        } else if str1[i-1] == str2[j-1] {
            sb.WriteByte(str1[i-1])
            i--
            j--
        } else if dp[i-1][j] > dp[i][j-1] {
            sb.WriteByte(str1[i-1])
            i--
        } else {
            sb.WriteByte(str2[j-1])
            j--
        }
    }

    // 反转结果字符串
    res := []byte(sb.String())
    for left, right := 0, len(res)-1; left < right; left, right = left+1, right-1 {
        res[left], res[right] = res[right], res[left]
    }
    return string(res)
}
*/