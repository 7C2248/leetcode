/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	sum := 0
	kk := map[byte][]int{'I': []int{0, 1}, 'V': []int{1, 5}, 'X': []int{2, 10}, 'L': []int{3, 50}, 'C': []int{4, 100}, 'D': []int{5, 500}, 'M': []int{6, 1000}}
	for k := 0; k < len(s); {
		if k == len(s)-1 {
			sum += kk[s[k]][1]
			return sum
		}
		if kk[s[k]][0]%2 == 0 && kk[s[k]][0] < kk[s[k+1]][0] && kk[s[k]][0] >= kk[s[k+1]][0]-2 {
			sum += kk[s[k+1]][1] - kk[s[k]][1]
			k += 2
			continue
		}
		sum += kk[s[k]][1]
		k += 1
	}
	return sum
}

// @lc code=end

