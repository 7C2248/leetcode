/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
type Trie struct {
	neib   []*Trie
	letter byte
}

type LetterIndex struct {
	alphabit [26][]*Trie
}

func BuildTrieFromMatrix(matrix [][]byte) *LetterIndex {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return &LetterIndex{}
	}

	n, m := len(matrix), len(matrix[0])
	index := &LetterIndex{}
	for i := 0; i < 26; i++ {
		index.alphabit[i] = make([]*Trie, 0)
	}

	nodes := make([][]*Trie, n)
	for i := 0; i < n; i++ {
		nodes[i] = make([]*Trie, m)
		for j := 0; j < m; j++ {
			node := &Trie{
				neib:   make([]*Trie, 0),
				letter: matrix[i][j],
			}
			nodes[i][j] = node

			idx := matrix[i][j] - 'a'
			index.alphabit[idx] = append(index.alphabit[idx], node)
		}
	}

	// 构建邻接关系（上下左右四个方向）
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			currentNode := nodes[i][j]

			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < n && nj >= 0 && nj < m {
					neighbor := nodes[ni][nj]
					currentNode.neib = append(currentNode.neib, neighbor)
				}
			}
		}
	}

	return index
}

func findWords(board [][]byte, words []string) []string {

	index := BuildTrieFromMatrix(board)
	result := []string{}

	for i := range words {
		word, nw := words[i], len(words[i])
		start := index.alphabit[word[0]-'a']
		found := false

		for j := range start {
			visited := make(map[*Trie]bool)
			var dfs func(now *Trie, index int) bool

			dfs = func(now *Trie, index int) bool {
				if now.letter != word[index] {
					return false
				}

				if index == nw-1 {
					return true
				}
				visited[now] = true
				defer delete(visited, now)

				for i := range now.neib {
					if !visited[now.neib[i]] {
						if dfs(now.neib[i], index+1) {
							return true
						}
					}
				}
				return false
			}

			if dfs(start[j], 0) {
				found = true
				break
			}
		}
		if found {
			result = append(result, words[i])
		}
	}

	return result
}

// @lc code=end

