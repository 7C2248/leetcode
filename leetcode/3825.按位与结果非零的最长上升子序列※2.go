




func longestSubsequence(nums []int) (ans int) {
	// 统计最大的二进制位
	w := bits.Len(uint(slices.Max(nums)))

	// 枚举每个二进制位
	for i := range w {
		f := []int{}
		for _, x := range nums {
			// 只有当当前二进制位非0，才进行递增子序列搜索
			if x>>i&1 == 0 { // x 二进制的第 i 位是 0
				continue
			}
			// 标准二分+贪心的递增子序列搜索
			j := sort.SearchInts(f, x)
			if j < len(f) {
				f[j] = x
			} else {
				f = append(f, x)
			}
		}
		ans = max(ans, len(f))
	}
	return
}



