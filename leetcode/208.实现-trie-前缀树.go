/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	exist  bool
	suffix []*Trie
	letter byte
}

func Constructor() Trie {
	return Trie{exist: false, suffix: []*Trie{}, letter: 0}
}

func (this *Trie) Insert(word string) {
	n, now := len(word), this
	for i := 0; i < n; i++ {
		sign := false
		for j := 0; j < len(now.suffix); j++ {
			if word[i] == now.suffix[j].letter {
				now = now.suffix[j]
				sign = true
				break
			}
		}
		if !sign {
			r := &Trie{exist: false, suffix: []*Trie{}, letter: word[i]}
			now.suffix = append(now.suffix, r)
			now = r
		}
	}

	now.exist = true
}

func (this *Trie) Search(word string) bool {
	n, now := len(word), this
	for i := 0; i < n; i++ {
		sign := false
		for j := 0; j < len(now.suffix); j++ {
			if word[i] == now.suffix[j].letter {
				now = now.suffix[j]
				sign = true
				break
			}
		}
		if !sign {
			return false
		}
	}
	if now.exist {
		return true
	} else {
		return false
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	n, now := len(prefix), this
	for i := 0; i < n; i++ {
		sign := false
		for j := 0; j < len(now.suffix); j++ {
			if prefix[i] == now.suffix[j].letter {
				now = now.suffix[j]
				sign = true
				break
			}
		}
		if !sign {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

