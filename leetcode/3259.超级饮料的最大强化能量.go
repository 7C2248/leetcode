/*
 * @lc app=leetcode.cn id=3259 lang=golang
 *
 * [3259] 超级饮料的最大强化能量
 */

// @lc code=start
func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	dpA, dpB := 0, 0
	n := len(energyDrinkA)
	for i := 0; i < n; i++ {
		dpA, dpB = max(dpA+energyDrinkA[i], dpB), max(dpB+energyDrinkB[i], dpA)
	}
	return int64(max(dpA, dpB))
}

// @lc code=end

