/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
type nn struct {
	Gene  string
	count int
}

func minMutation(startGene string, endGene string, bank []string) int {

	n := len(startGene)
	queue := []nn{{Gene: startGene, count: 0}}

	visited := make(map[string]bool)

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		currentGene, currentCount := current.Gene, current.count

		if currentGene == endGene {
			return currentCount
		}

		for i := range bank {
			count := 0
			for j := 0; j < n; j++ {

				if bank[i][j] != currentGene[j] {
					count++
				}
				if count > 1 {
					break
				}

			}
			if count == 1 && !visited[bank[i]] {
				visited[bank[i]] = true
				queue = append(queue, nn{Gene: bank[i], count: currentCount + 1})
			}
		}

	}

	return -1
}

// @lc code=end

