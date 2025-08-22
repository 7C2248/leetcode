/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
type Trie struct {
	children map[byte]*Trie
	word     string
}

func (this *Trie) Insert(word string) {
	node := this
	for i := range word {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = &Trie{children: map[byte]*Trie{}}
		}
		node = node.children[ch]
	}
	node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) (ans []string) {
	root := &Trie{children: map[byte]*Trie{}}
	for _, word := range words {
		root.Insert(word)
	}

	m, n := len(board), len(board[0])

	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		nxt := node.children[ch]

		//检查前缀
		if nxt == nil {
			return
		}

		if nxt.word != "" {
			ans = append(ans, nxt.word)
			nxt.word = ""
		}

		if len(nxt.children) > 0 {
			board[x][y] = '#'
			for _, d := range dirs {
				nx, ny := x+d.x, y+d.y
				if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
					dfs(nxt, nx, ny)
				}
			}
			board[x][y] = ch
		}

		if len(nxt.children) == 0 {
			delete(node.children, ch)
		}
	}
	for i, row := range board {
		for j := range row {
			dfs(root, i, j)
		}
	}

	return
}

// @lc code=end

