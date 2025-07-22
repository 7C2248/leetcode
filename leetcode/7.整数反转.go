/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
import (
	"fmt"
	"strconv"
)

func reverse(x int) int {

	var reverse0 bool
	if x < 0 {
		x = -x
		reverse0 = true
	} else {
		reverse0 = false
	}
	result := ""

	xx := fmt.Sprintf("%d", x)

	for i := len(xx) - 1; i >= 0; i-- {
		result += string(xx[i])
	}
	aa, _ := strconv.Atoi(result)
	if (aa & 0xf80000000) > 0 {
		return 0
	} else {
		if reverse0 {
			return -aa
		} else {
			return aa
		}
	}

}

// @lc code=end

