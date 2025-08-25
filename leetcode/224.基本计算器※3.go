/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start

func calculate(s string) int {

	signstack := make([]byte, 0)
	numstack := make([]int, 0)
	l := len(s)
	isNumber := false

	for i := 0; i < l; {

		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] < '0' || s[i] > '9' {

			switch s[i] {
			case '(':
				signstack = append(signstack, s[i])
				isNumber = false
			case ')':
				for len(signstack) > 0 && signstack[len(signstack)-1] != '(' {
					compute(&signstack, &numstack)
				}
				signstack = signstack[:len(signstack)-1]
				isNumber = true
			case '+', '*', '/':
				for len(signstack) > 0 && signstack[len(signstack)-1] != '(' && getPriority(signstack[len(signstack)-1]) >= getPriority(s[i]) {
					compute(&signstack, &numstack)
				}
				isNumber = false
				signstack = append(signstack, s[i])
			case '-':
				if !isNumber {
					signstack = append(signstack, 'n')
				} else {
					for len(signstack) > 0 && signstack[len(signstack)-1] != '(' && getPriority(signstack[len(signstack)-1]) >= getPriority(s[i]) {
						compute(&signstack, &numstack)
					}
					signstack = append(signstack, s[i])
				}
				isNumber = false
			}
			i++
			continue
		}
		num := 0
		for ; i < l && s[i] >= '0' && s[i] <= '9'; i++ {
			num = num*10 + int(s[i]-'0')
		}
		numstack = append(numstack, num)
		isNumber = true

	}
	for len(signstack) > 0 {
		compute(&signstack, &numstack)
	}

	return numstack[0]
}

func getPriority(sign byte) (result int) {

	switch sign {

	case 'n':
		result = 3
	case '*', '/':
		result = 2
	case '+', '-':
		result = 1
	default:
		result = 0
	}

	return
}

func compute(signstack *[]byte, numstack *[]int) {
	sign := (*signstack)[len(*signstack)-1]
	*signstack = (*signstack)[:len(*signstack)-1]
	result := 0
	if sign == 'n' {

		result = -(*numstack)[len(*numstack)-1]
		*numstack = (*numstack)[:len(*numstack)-1]

	} else {
		n0, n1 := (*numstack)[len(*numstack)-1], (*numstack)[len(*numstack)-2]
		*numstack = (*numstack)[:len(*numstack)-2]
		switch sign {
		case '+':
			result = n1 + n0
		case '-':
			result = n1 - n0
		case '*':
			result = n1 * n0
		case '/':
			result = n1 / n0
		default:
			panic("unsupported operator")
		}
	}
	*numstack = append(*numstack, result)
	return
}

// @lc code=end

