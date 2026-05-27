/*
 * @lc app=leetcode.cn id=2321 lang=golang
 *
 * [2321] 拼接数组的最大分数
 */

// @lc code=start

func maximumsSplicedArray(nums1 []int, nums2 []int) int {

	n := len(nums1)
	curMax, curMin := nums1[0]-nums2[0], nums1[0]-nums2[0]
	sum1, sum2 := nums1[0], nums2[0]
	maximum, minimum := curMax, curMin
	for i := 1; i < n; i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
		cur := nums1[i] - nums2[i]
		curMax = max(curMax, 0) + cur
		curMin = min(curMin, 0) + cur
		if curMax > maximum {
			maximum = curMax
		}
		if curMin < minimum {
			minimum = curMin
		}
	}

	return max(maximum+sum2, sum1-minimum)
}

// @lc code=end

/*
[20,40,20,70,30]
[20,60,60,90,100]
[20,60,80,130,120]
[20,60,80,150,160]
[20,60,80,150,180]

[50,20,50,40,20] [30,-20,30,-30,-10]
[50,70,70,90,60] [30,10,10,0,-40]
[50,70,120,110,110] [30,10,40,-20,-10]
[50,70,120,160,130] [30,10,40,10,-30]
[50,70,120,160,180] [30,10,40,10,0]

交换后的最大值要求我们找到差值绝对值最大的连续子数组和，
即max(sub_nums1-sub_nums2)
注意到两个子数组的元素数量相等，从而可以按照元素进行拆分，
所以上式又可以表示为

max(sum(nums1[l]-nums2[l],...,nums1[r]-nums2[r]))

记nums1-nums2为新数组，上述计算过程正式求该数组的连续子元素最大和，从而

dp[i-1] + nums[i] > nums[i]
max(dp[i-1],0) +num[i]

可以解决
而对于绝对值只需要同时计算最大和以及最小和即可
*/