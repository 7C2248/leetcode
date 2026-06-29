/*
 * @lc app=leetcode.cn id=1745 lang=golang
 *
 * [1745] 分割回文串 IV
 */

// @lc code=start
func checkPartitioning(s string) bool {
	k := 3
	n := len(s)
	pre := make([]bool, n+1)
	now := make([]bool, n+1)
	pre[0] = true

	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			// 奇数
			for m := 0; j+m < n && j-m > -1; m++ {
				if s[j+m] == s[j-m] {
					if pre[j-m] {
						now[j+m+1] = pre[j-m]
					}
				} else {
					break
				}
			}
			//偶数
			for m := 0; j+m+1 < n && j-m > -1; m++ {
				if s[j+m+1] == s[j-m] {
					if pre[j-m] {
						now[j+m+2] = pre[j-m]
					}
				} else {
					break
				}
			}
		}
		pre, now = now, pre
		for j := 0; j <= n; j++ {
			now[j] = false
		}
	}
	return pre[n]
}

// @lc code=end
