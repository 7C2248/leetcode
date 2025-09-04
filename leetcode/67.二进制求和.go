/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}

	na, nb := len(a), len(b)
	result := make([]byte, nb+1)
	index := nb

	carry := byte(0)

	for i := 0; i < nb; i++ {
		var bitA, bitB byte

		if i < na {
			bitA = a[na-1-i] - '0'
		} else {
			bitA = 0
		}

		bitB = b[nb-1-i] - '0'

		sum := bitA + bitB + carry
		result[index] = sum%2 + '0'
		carry = sum / 2
		index--
	}

	if carry > 0 {
		result[index] = '1'
		index--
	}

	return string(result[index+1:])
}

// @lc code=end

