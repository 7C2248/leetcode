/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
type nn struct {
	word  string
	count int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	n := len(beginWord)

	sign := false
	for _, w := range wordList {
		if w == endWord {
			sign = true
			break
		}
	}
	if !sign {
		return 0
	}

	queueFront := []nn{{word: beginWord, count: 1}}
	visitedFront := map[string]int{beginWord: 1}

	queueBack := []nn{{word: endWord, count: 1}}
	visitedBack := map[string]int{endWord: 1}

	for len(queueFront) > 0 && len(queueBack) > 0 {

		if len(queueFront) > len(queueBack) {
			queueFront, queueBack = queueBack, queueFront
			visitedFront, visitedBack = visitedBack, visitedFront
		}
		//按层次进行处理
		nextQueue := []nn{}
		for _, cur := range queueFront {
			curWord, curStep := cur.word, cur.count
			for _, w := range wordList {
				diff := 0
				for i := 0; i < n; i++ {
					if curWord[i] != w[i] {
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
						nextQueue = append(nextQueue, nn{word: w, count: curStep + 1})
					}
				}
			}
		}
		queueFront = nextQueue
	}
	return 0
}

// @lc code=end

