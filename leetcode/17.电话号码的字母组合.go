/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var mapping = map[string][]string{"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"}}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return mapping[digits]
	}
	result := []string{}
	comb := letterCombinations(digits[1:])
	cn := mapping[digits[0:1]]

	for _, i := range cn {
		for _, j := range comb {
			result = append(result, i+j)
		}
	}
	return result
}

// @lc code=end

