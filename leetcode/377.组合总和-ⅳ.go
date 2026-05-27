/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
import (
	"sort"
)

func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	for i := 0; i < target; i++ {
		dp[i] = -1
	}
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	dp[0] = 0
	for i := 1; i <= target; i++ {
		for j := 0; j < n; j++ {
			if i-nums[j] >= 0 {
				if dp[i-nums[j]] != -1 {
					dp[i] += 1 + dp[i-nums[j]]
				}
			} else {
				break
			}
		}
	}
	return dp[target]
}

/*
//标准解法 dp[x] 表示选取的元素之和等于 x 的方案数
func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1
    for i := 1; i <= target; i++ {
        for _, num := range nums {
            if num <= i {
                dp[i] += dp[i-num]
            }
        }
    }
    return dp[target]
}
*/

// @lc code=end

