/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
import (
	"fmt"
)

func isPalindrome(x int) bool {
	xx := fmt.Sprintf("%d", x)
	var i, j int = 0, len(xx) - 1
	for i <= j {
		if xx[i] == xx[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

