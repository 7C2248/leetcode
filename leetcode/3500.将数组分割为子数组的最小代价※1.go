/*
 * @lc app=leetcode.cn id=3500 lang=golang
 *
 * [3500] 将数组分割为子数组的最小代价
 */

// @lc code=start
func minimumCost(nums []int, cost []int, k int) int64 {

	n := len(nums)
	sumN := make([]int, n+1)
	sumC := make([]int, n+1)
	suffixC := make([]int,n+1)
	for i := 1; i <= n; i++ {
		sumN[i] += sumN[i-1] + nums[i-1]
		sumC[i] += sumC[i-1] + cost[i-1]
	}
	for i:=n-1;i>-1;i--{
		suffixC[i] += suffixC[i+1]+cost[i]
	}
	// 不再是表达nums[:i]的最小划分代价，转变为计算dp[n]的中间变量。
	// dp的转移由题目要求进行决定
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
		for j := 0; j < i; j++{
			segCost := sumN[i] * (sumC[i] - sumC[j])
			extra := k * suffixC[j]
			total := dp[j] + segCost + extra
			dp[i] = min(dp[i], total)
		}
	}
	return int64(dp[n])
}

// @lc code=end

totalCost
=sum((sumN[r]+ k*i) * (sumC[r]-sumC[l]))
=sum(sn *cost) + sum(k*i*cost)

sum(k*i*cost) 
= k* sum(i*cost)
= k * Σ( i * cost^r_l)
	  = 1*cost^r0_l0 + 2*cost^r1_l1 + ... + i*cost^ri_li
	  = (cost^r0_l0+cost^r1_l1+...+cost^ri_li)+(cost^r1_l1+...+cost^ri_li)+...+cost^ri_li
	  = cost^ri_l0 + cost^ri_l1 +... +cost^ri_li