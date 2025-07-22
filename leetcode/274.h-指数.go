/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
func sort(aa []int, left int, right int) {

	if right-left <= 1 {
		return
	}

	i, j := left, right-1
	for i < j {
		for aa[j] > aa[left] && i < j {
			j--
		}
		for aa[i] <= aa[left] && i < j {
			i++
		}

		aa[i], aa[j] = aa[j], aa[i]
	}
	aa[left], aa[i] = aa[i], aa[left]
	sort(aa, left, i)
	sort(aa, i+1, right)
	return
}
func hIndex(citations []int) int {
	n := len(citations)
	sort(citations, 0, n)
	h := 0
	for i := 0; i < n; i++ {
		temp := []int{n - i, citations[i], n}
		sort(temp, 0, 3)
		if temp[0] > h {
			h = temp[0]
		}
	}
	return h
}

// @lc code=end

