/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
type nn struct {
	geng  string
	count int
}

func minMutation(startGene string, endGene string, bank []string) int {

	n := len(startGene)

	sign := false
	for _, w := range bank {
		if w == endGene {
			sign = true
			break
		}
	}
	if !sign {
		return -1
	}

	queueFront := []nn{{geng: startGene, count: 0}}
	visitedFront := map[string]int{startGene: 0}

	queueBack := []nn{{geng: endGene, count: 1}}
	visitedBack := map[string]int{endGene: 1}

	for len(queueFront) > 0 && len(queueBack) > 0 {

		if len(queueFront) > len(queueBack) {
			queueFront, queueBack = queueBack, queueFront
			visitedFront, visitedBack = visitedBack, visitedFront
		}
		//按层次进行处理
		nextQueue := []nn{}
		for _, cur := range queueFront {
			curGeng, curStep := cur.geng, cur.count
			for _, w := range bank {
				diff := 0
				for i := 0; i < n; i++ {
					if curGeng[i] != w[i] {
						diff++
						if diff > 1 {
							break
						}
					}
				}
				if diff == 1 {
					if step, ok := visitedBack[w]; ok {
						return curStep + step
					}
					if _, ok := visitedFront[w]; !ok {
						visitedFront[w] = curStep + 1
						nextQueue = append(nextQueue, nn{geng: w, count: curStep + 1})
					}
				}
			}
		}
		queueFront = nextQueue
	}
	return -1
}

// @lc code=end

