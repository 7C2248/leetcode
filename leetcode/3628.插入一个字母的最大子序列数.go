/*
 * @lc app=leetcode.cn id=3628 lang=golang
 *
 * [3628] 插入一个字母的最大子序列数
 */

// @lc code=start
func numOfSubsequences(s string) int64 {

	l := len(s)
	if l < 2 {
		return 0
	}
	C := 0
	dc := make([]int, l)
	L := 0
	dl := 0
	dp := [3]int{0, 0, 0}
	for i := 0; i < l; i++ {
		if s[i] == 'L' {
			dp[0] += 1
			dc[i] += 1
		}
		if s[i] == 'C' {
			dl += 1
			dp[1] += dp[0]
		}
		if s[i] == 'T' {
			L += dl
			dp[2] += dp[1]
		}
		if i > 0 {
			dc[i] += dc[i-1]
		}
	}
	T := 0
	dt := 0
	rt := 0
	for i := l - 1; i > -1; i-- {
		if s[i] == 'C' {
			dt += 1
		}
		if s[i] == 'L' {
			T += dt
		}
		if s[i] == 'T' {
			rt += 1
			C = max(C, dc[i]*rt)
		}
	}
	return int64(max(max(L, C), T) + dp[2])
}

// @lc code=end

/*
// 空间优化 时间复杂度和我的实现相同都为O(n)
func numOfSubsequences(s string) int64 {

	t := strings.Count(s, "T" )
	var l, lc, lct, c, ct, lt int
	for _, b := range s {
		switch b {
		case 'L':
			l = l + 1
		case 'C':
			c, lc = c+1, lc+l
		case 'T':
			t, ct, lct = t-1, ct+c, lct+lc
		}
        lt = max(lt, l*t)
	}
    return int64(lct + max(lc,ct,lt))
}
*/