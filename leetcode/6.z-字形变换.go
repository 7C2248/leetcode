/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start

func convert(s string, numRows int) string {

	if numRows == 1 || len(s) <= numRows {
		return s
	}
	result := ""
	k := 0
	for i := 0; i < numRows; i++ {
		k = (numRows - 1 - i) % (numRows - 1)
		for j := 0; i+j < len(s); j += numRows*2 - 2 {
			result += string(s[i+j])
			if k != 0 && i+j+k*2 < len(s) {
				result += string(s[i+j+k*2])
			}
		}
	}
	return result
}

// @lc code=end

