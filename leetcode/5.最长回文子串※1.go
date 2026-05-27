/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func normal_found1(s string, pos int, baselen int) int {
	n := len(s)
	d1 := baselen
	for i := 1; pos+baselen-1+i < n && pos-baselen+1-i > -1; i++ {
		if s[pos+baselen-1+i] == s[pos-baselen+1-i] {
			d1 += 1
		} else {
			break
		}
	}
	return d1
}
func longestPalindrome(s string) string {
	/*manacher*/
	l, r := 0, -1
	n := len(s)
	temp := "#"
	for i := 0; i < n; i++ {
		temp += string(s[i]) + "#"
	}
	dmax, pos := 0, 0
	d := make([]int, n*2+1)

	for i := 0; i < n*2+1; i++ {
		if i > r {
			d[i] = normal_found1(temp, i, 1)
		} else {
			j := l + r - i
			if j-d[j]+1 <= l {
				d[i] = normal_found1(temp, i, r-i+1)
			} else {
				d[i] = d[j]
			}
		}
		if d[i] > dmax {
			dmax = d[i]
			pos = i
		}

		if i+d[i]-1 > r {
			r = i + d[i] - 1
			l = i - d[i] + 1
		}
	}

	start, end := 0, 0
	dmax = (dmax - 1) / 2
	if temp[pos] == '#' {
		start, end = (pos-1)/2-dmax+1, (pos-1)/2+dmax
	} else {
		start, end = (pos-1)/2-dmax, (pos-1)/2+dmax
	}
	return s[start : end+1]
}
func longestPalindrome0(s string) string {

	n := len(s)
	if n <= 1 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	max, start := 1, 0
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			if s[i] == s[i+length-1] {
				if i+1 <= i+length-2 {
					dp[i][i+length-1] = dp[i+1][i+length-2]
				} else {
					dp[i][i+length-1] = true
				}
			} else {
				dp[i][i+length-1] = false
			}

			if dp[i][i+length-1] == true {
				if length > max {
					max = length
					start = i
				}
			}
		}
	}

	return s[start : start+max]
}

// @lc code=end

