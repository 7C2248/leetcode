/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	sum := [16]int{}
	for i := range nums {
		for j := 0; j < 16; j++ {
			h, l := (nums[i]>>(j*2+1))&1, (nums[i]>>(j*2))&1
			sum[j] += (h << 16) + l
		}
	}
	result := 0
	for i := range sum {
		h, l := (sum[i]>>16)%3, (sum[i]&0xffff)%3
		result |= ((h << 1) | l) << (i * 2)
	}
	return int(int32(result))
}

// @lc code=end

