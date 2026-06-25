/*
 * @lc app=leetcode.cn id=3578 lang=golang
 *
 * [3578] 统计极差最大为 K 的分割方式数
 */

// @lc code=start
func countPartitions(nums []int, k int) int {
	const mod = 1_000_000_000 + 7
	n := len(nums)
	pre := make([]int,n+2) 
	pre[1] = 1
	// 单调队列，存储对应极值的index。
    maxQ := make([]int, 0)
    minQ := make([]int, 0)
	L := 0
	dp := 0
	for i := 0; i < n ; i++ {
		x := nums[i]
		
		// 每次入队新元素时，将之前比它小的元素出队，
		// 因为这些元素不可能成为下一个极大值
		for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= x {
            maxQ = maxQ[:len(maxQ)-1]
        }
        maxQ = append(maxQ, i)

        for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= x {
            minQ = minQ[:len(minQ)-1]
        }
        minQ = append(minQ, i)

		// 窗口收缩时，如果被移出去的下标 L 恰好是 maxQ 或 minQ 的队首，
		// 说明当前最大/最小值已经离开窗口，直接弹出队首。
		// 弹出后新的队首自然就是剩余元素中的极值。
		for len(maxQ) > 0 && len(minQ) > 0 &&
            nums[maxQ[0]]-nums[minQ[0]] > k {
            if maxQ[0] == L {
                maxQ = maxQ[1:]
            }
            if minQ[0] == L {
                minQ = minQ[1:]
            }
            L++
        }

		dp = pre[i+1]-pre[L]+ mod
		dp %= mod
		pre[i+2] = pre[i+1]+dp
		pre[i+2] %= mod
	}
	
	return dp
}
// @lc code=end

  9,4,1,3,7
1,1,2,4,8,16
1,1,1,2,4,6
