/*
 * @lc app=leetcode.cn id=2111 lang=golang
 *
 * [2111] 使数组 K 递增的最少操作次数
 */

// @lc code=start
func kIncreasing(arr []int, k int) int {
	n := len(arr)
	result := 0
	// 最大组大小，预分配 dp 的容量
	maxSize := (n + k - 1) / k
	// 长度为0，容量为maxSize
	dp := make([]int, 0, maxSize)

	for i := 0; i < k; i++ {
		dp = dp[:0] // 重置长度，但保留容量
		cnt := 0
		for j := i; j < n; j += k {
			x := arr[j]
			// 二分查找，找第一个 > x 的位置
			pos := sort.Search(len(dp), func(r int) bool { return dp[r] > x })
			if pos == len(dp) {
				dp = append(dp, x)
			} else {
				dp[pos] = x
			}
			cnt++
		}
		result += cnt - len(dp)
	}
	return result
}

// @lc code=end
