/*
 * @lc app=leetcode.cn id=2518 lang=golang
 *
 * [2518] 好分区的数目
 */

// @lc code=start

func countPartitions(nums []int, k int) int {
	const mod = 1_000_000_007
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < 2*k {
		return 0
	}

	// 正向计算空间需求过大，通过计算不满足的方案减少计算量。
	dp := make([]int, k)
	dp[0] = 1
	for _, num := range nums {
		for s := k - 1; s >= num; s-- {
			dp[s] = (dp[s] + dp[s-num]) % mod
		}
	}

	// 总方案 2^n
	pow2 := 1
	for i := 0; i < len(nums); i++ {
		pow2 = (pow2 * 2) % mod
	}

	// 不合法方案 = 2 * Σ_{s=0}^{k-1} dp[s]
	invalid := 0
	for _, v := range dp {
		invalid = (invalid + v) % mod
	}
	invalid = (invalid * 2) % mod

	ans := (pow2 - invalid) % mod
	if ans < 0 {
		ans += mod
	}
	return ans
}

// @lc code=end

