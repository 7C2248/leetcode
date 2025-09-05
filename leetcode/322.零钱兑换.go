/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
import (
	"sort"
)

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	sort.Ints(coins)
	mk := make([]int, amount+1)
	mk[amount] = -1
	for _, v := range coins {
		if v <= amount {
			mk[v] = 1
		}
	}

	for i := coins[0]; i < amount; i++ {
		if mk[i] <= 0 {
			continue
		}
		for _, v := range coins {
			if i+v <= amount {
				if mk[i+v] > 0 {
					mk[i+v] = min(mk[i+v], mk[i]+1)
				} else {
					mk[i+v] = mk[i] + 1
				}
			}
		}
	}
	return mk[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

