/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {

	l := len(intervals)
	qs(intervals, 0, l)
	k := 0
	var j int
	for i := 0; i < l; {
		for j = i + 1; j < l && intervals[i][1] >= intervals[j][0]; j++ {

			if intervals[j][1] > intervals[i][1] {
				intervals[i][1] = intervals[j][1]
			}

		}
		intervals[k] = intervals[i]
		k++
		i = j
	}

	return intervals[:k]
}

func qs(aa [][]int, left int, right int) {

	if right-left <= 1 {
		return
	}

	i, j := left, right-1
	for i < j {
		for ((aa[j][0] > aa[left][0]) || (aa[j][0] == aa[left][0] && aa[j][1] > aa[left][1])) && i < j {
			j--
		}
		for (aa[i][0] < aa[left][0] || (aa[i][0] == aa[left][0] && aa[i][1] <= aa[left][1])) && i < j {
			i++
		}

		aa[i], aa[j] = aa[j], aa[i]
	}
	aa[left], aa[i] = aa[i], aa[left]
	qs(aa, left, i)
	qs(aa, i+1, right)
	return

}

// @lc code=end

