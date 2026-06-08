/*
 * @lc app=leetcode.cn id=3287 lang=golang
 *
 * [3287] 求出数组中最大序列值
 */

// @lc code=start

// ---------- 128 位掩码工具 ----------
type bitset struct{ lo, hi uint64 }

func (b *bitset) set(v int) {
	if v < 64 {
		b.lo |= 1 << v
	} else {
		b.hi |= 1 << (v - 64)
	}
}
func (b *bitset) has(v int) bool {
	if v < 64 {
		return (b.lo>>v)&1 == 1
	}
	return (b.hi>>(v-64))&1 == 1
}

// 返回一个新的 bitset，表示原集合中每个值都 OR num 的结果
func (b *bitset) applyOr(num int) bitset {
	var res bitset
	// 只遍历可能的值，0~127
	for v := 0; v < 128; v++ {
		if b.has(v) {
			res.set(v | num)
		}
	}
	return res
}

// 将 b2 合并到 b1（按位或）
func (b *bitset) merge(b2 bitset) {
	b.lo |= b2.lo
	b.hi |= b2.hi
}

// 位运算优化
func maxValue(nums []int, k int) int {
	n := len(nums)

	// ---------- 前缀 DP，记录每个位置的 preK ----------
	preK := make([]bitset, n+1)
	dp := make([]bitset, k+1)
	dp[0].set(0) // 只有 0 可达

	for i := 1; i <= n; i++ {
		num := nums[i-1]
		// 倒序更新，保证每个元素只用一次
		for j := k; j >= 1; j-- {
			// 从 dp[j-1] 的集合应用 num 再合并
			newSet := dp[j-1].applyOr(num)
			dp[j].merge(newSet)
		}
		if i >= k {
			preK[i] = dp[k]
		}
	}

	// ---------- 后缀 DP，记录每个位置的 sufK ----------
	sufK := make([]bitset, n+1)
	dp2 := make([]bitset, k+1)
	dp2[0].set(0)

	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		for j := k; j >= 1; j-- {
			newSet := dp2[j-1].applyOr(num)
			dp2[j].merge(newSet)
		}
		if n-i >= k {
			sufK[i] = dp2[k]
		}
	}

	// ---------- 枚举分割点，计算最大 XOR ----------
	ans := 0
	for i := k; i <= n-k; i++ {
		for v1 := 0; v1 < 128; v1++ {
			if !preK[i].has(v1) {
				continue
			}
			for v2 := 0; v2 < 128; v2++ {
				if !sufK[i].has(v2) {
					continue
				}
				if xor := v1 ^ v2; xor > ans {
					ans = xor
				}
			}
		}
	}
	return ans
}

// @lc code=end

/*

// 设定分割点，对分割点遍历，搜索前后长度为k的所有 或可能。
// 对每个分割点，遍历或可能计算异或最大值。

func maxValue(nums []int, k int) int {

	// 遍历以所有位置为起始，长度为k的子序列的可能异或值
    findORs := func(nums []int, k int) []map[int]bool {
		dp := make([]map[int]bool, 0)
		prev := make([]map[int]bool, k + 1)
		for i := 0; i <= k; i++ {
			prev[i] = make(map[int]bool)
		}
		prev[0][0] = true

		for i := 0; i < len(nums); i++ {
			for j := min(k - 1, i + 1); j >= 0; j-- {
				for x := range prev[j] {
					prev[j + 1][x | nums[i]] = true
				}
			}
			current := make(map[int]bool)
			for key := range prev[k] {
				current[key] = true
			}
			dp = append(dp, current)
		}
		return dp
	}

	A := findORs(nums, k)
	reverse(nums)
	B := findORs(nums, k)
	mx := 0

	for i := k - 1; i < len(nums) - k; i++ {
		for a := range A[i] {
			for b := range B[len(nums) - i - 2] {
				if a ^ b > mx {
					mx = a ^ b
				}
			}
		}
	}
	return mx
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

*/

/*

//我的实现，需要优化，对于每个起始通过dp保存之前起始的可行值。
func maxValue(nums []int, k int) int {

	n := len(nums)
	orSum := 0
	for _, v := range nums {
		orSum |= v
	}
	result := -1
	// 分割点
	for i := k; i <= n-k; i++ {
		//在前 i 个元素中，选出长度为 j 的子序列，能达到的所有 OR 值的集合。
		pre := make([][]bool, k+1)
		//在后 i 个元素中，选出长度为 j 的子序列，能达到的所有 OR 值的集合。
		suf := make([][]bool, k+1)
		for j := 0; j <= k; j++ {
			pre[j] = make([]bool, orSum+1)
			suf[j] = make([]bool, orSum+1)
		}
		pre[0][0] = true
		suf[0][0] = true
		v1 := []int{}
		// 遍历前i个元素
		for j := 0; j < i; j++ {
			// 倒序更新
			for kk := k - 1; kk > -1; kk-- {
				for oS, v0 := range pre[kk] {
					if v0 {
						pre[kk+1][oS|nums[j]] = true
					}
				}
			}
			for kk, v := range pre[k] {
				if v {
					v1 = append(v1, kk)
				}
			}
		}

		v2 := []int{}
		for j := n - 1; j >= i; j-- {
			for kk := k - 1; kk > -1; kk-- {
				for oS, v0 := range suf[kk] {
					if v0 {
						suf[kk+1][oS|nums[j]] = true
					}
				}
			}
			for kk, v := range suf[k] {
				if v {
					v2 = append(v2, kk)
				}
			}
		}

		for _, v11 := range v1 {
			for _, v22 := range v2 {
				result = max(result, v11^v22)
			}
		}

	}
	return result
}
*/
