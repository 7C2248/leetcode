func minRemovals(nums []int, target int) int {
	n := len(nums)
	maximum := getMax(nums)
	if target > maximum {
		return -1
	}
	pre := make([]int, maximum+1)
	now := make([]int, maximum+1)
	for i := 0; i <= maximum; i++ {
		pre[i] = -1
	}
	pre[0] = 0

	for _, num := range nums {
		copy(now, pre)
		for x, v := range pre {
			if v != -1 {
				nx := x ^ num
				if v+1 > now[nx] {
					now[nx] = v + 1
				}
			}
		}
		pre, now = now, pre
	}

	if pre[target] == -1 {
		return -1
	}
	return n - pre[target]
}

func getMax(a []int) int {
	maximum := -1
	for _, v := range a {
		if v > maximum {
			maximum = v
		}
	}
	n := 0
	for maximum > 0 {
		maximum >>= 1
		n += 1
	}

	return (1 << n) - 1
}