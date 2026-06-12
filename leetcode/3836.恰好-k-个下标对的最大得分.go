func maxScore(nums1 []int, nums2 []int, k int) int64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 < l2 {
		l1, l2 = l2, l1
		nums1, nums2 = nums2, nums1
	}
	pre := make([][]int, l2+1)
	now := make([][]int, l2+1)
	for i := 0; i <= l2; i++ {
		pre[i] = make([]int, k+1)
		now[i] = make([]int, k+1)
	}
	nowk := make([]bool, k+1)
	prek := make([]bool, k+1)
	nowk[0] = true
	prek[0] = true
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			for l := 1; l <= min(min(i, j), k); l++ {
				best := nums1[i-1] * nums2[j-1]
				if l > 1 {
					best += pre[j-1][l-1]
				}
				if i > 1 && prek[l] {
					best = max(best, pre[j][l])
				}
				if j > 1 && nowk[l] {
					best = max(best, now[j-1][l])
				}
				now[j][l] = best
				nowk[l] = true
			}
		}
		prek, nowk = nowk, prek
		for j := 1; j <= k; j++ {
			nowk[j] = false
		}
		pre, now = now, pre
	}
	return int64(pre[l2][k])
}