/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start

type WordDictionary struct {
	children [26]*WordDictionary
	isend    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &WordDictionary{}
		}
		node = node.children[ch]
	}
	node.isend = true
}

func (this *WordDictionary) Search(word string) bool {
	node := this

	i := 0
	for i < len(word) && word[i] != '.' {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
		i++
	}
	if i == len(word) && node.isend {
		return true
	}
	if i == len(word) && !node.isend {
		return false
	}

	//现在还没有结束，只有可能是因为 word[i] == '.'
	for j := 0; j < 26; j++ {
		if node.children[j] != nil {
			if node.children[j].Search(word[i+1:]) {
				return true
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

