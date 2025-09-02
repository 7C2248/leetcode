/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start
type node struct {
	sum  int
	pair [2]int
}
type heap struct {
	data []*node
}

func build(data []*node) *heap {
	h := heap{}
	for i := range data {
		h.push(data[i])
	}
	return &h
}
func (this *heap) Len() int {
	return len(this.data)
}
func (this *heap) down(index int) {
	n := this.Len()
	for {
		l := (index << 1) + 1
		r := (index << 1) + 2
		if l >= n {
			break
		}
		smalles := l
		if r < n && this.data[r].sum < this.data[l].sum {
			smalles = r
		}
		if this.data[smalles].sum < this.data[index].sum {
			this.data[index], this.data[smalles] = this.data[smalles], this.data[index]
			index = smalles
		} else {
			break
		}
	}
}
func (this *heap) up(index int) {
	for index > 0 {
		if this.data[(index-1)>>1].sum <= this.data[index].sum {
			break
		}
		this.data[(index-1)>>1], this.data[index] = this.data[index], this.data[(index-1)>>1]
		index = (index - 1) >> 1
	}
}
func (this *heap) pop() *node {
	result, n := this.data[0], this.Len()-1
	this.data[0] = this.data[n]
	this.data = this.data[0:n]
	if this.Len() > 0 {
		this.down(0)
	}
	return result
}
func (this *heap) push(v *node) {
	this.data = append(this.data, v)
	this.up(this.Len() - 1)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := build([]*node{})
	n1, n2 := len(nums1), len(nums2)
	count := 0
	result := [][]int{}

	for i := 0; i < k && i < n1; i++ {
		h.push(&node{nums1[i] + nums2[0], [2]int{i, 0}})
	}

	for count < k && h.Len() > 0 {
		r := h.pop()
		i, j := r.pair[0], r.pair[1]
		result = append(result, []int{nums1[i], nums2[j]})
		if j+1 < n2 {
			h.push(&node{nums1[i] + nums2[j+1], [2]int{i, j + 1}})
		}
		count++
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

