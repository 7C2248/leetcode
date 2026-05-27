/*
 * @lc app=leetcode.cn id=2266 lang=golang
 *
 * [2266] 统计打字方案数
 */

// @lc code=start
func countTexts(pressedKeys string) int {

	n := len(pressedKeys)
	dp := make([]int, n)
	sum := 1
	mod := int(1e9 + 7)
	for i := 0; i < n; {
		var maxlen int
		if pressedKeys[i] == '9' || pressedKeys[i] == '7' {
			maxlen = 4
		} else {
			maxlen = 3
		}
		var j int
		for j = i; j < n && pressedKeys[j] == pressedKeys[i]; j++ {

			for temp := j - 1; j-temp < maxlen && temp >= i; temp-- {
				dp[j] += dp[temp]
				dp[j] %= mod
			}
			if j-maxlen >= i {
				dp[j] += dp[j-maxlen]
				dp[j] %= mod
			} else {
				dp[j] += 1
			}
		}

		sum = sum * dp[j-1] % mod

		i = j
	}

	return sum
}

// @lc code=end

